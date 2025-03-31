package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"code.icod.de/dalu/nethttpoidc"
	"code.icod.de/dalu/oidc/options"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/MadAppGang/httplog"
	"github.com/dlukt/graphql-backend-starter/config"
	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/graph"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"github.com/vektah/gqlparser/v2/ast"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/mattn/go-sqlite3"
)

var (
	graphqlDebug = true
)

// graphqlCmd represents the graphql command
var graphqlCmd = &cobra.Command{
	Use:   "graphql",
	Short: "run the graphql backend",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("graphql called")
		setDatabaseURI()
		// client := openDB(config.DatabaseURIDev)
		var client *ent.Client
		if useSQLite {
			fmt.Println("Running with SQLite")
			var e error
			client, e = ent.Open(dialect.SQLite, config.SqliteDSN)
			if e != nil {
				return e
			}
			ctx := context.Background()
			if e := client.Schema.Create(ctx); e != nil {
				return e
			}
		} else {
			fmt.Println("Running with PostgreSQL")
			client = openDB(config.DatabaseURI)
		}
		defer client.Close()
		if e := client.Schema.Create(
			context.Background(),
		); e != nil {
			log.Fatal("opening ent client", e)
		}

		srv := NewDefaultServer(graph.NewSchema(client))
		srv.Use(entgql.Transactioner{TxOpener: client})

		cfg := config.OidcConfigDev

		oidcHandler := nethttpoidc.New(srv,
			options.WithIssuer(cfg.Issuer),
			options.WithRequiredTokenType("JWT"),
			options.WithRequiredAudience(cfg.Audience),
			options.IsPermissive(),
		)

		corsHandler := cors.AllowAll()
		fmt.Println("debug:", graphqlDebug)
		if !graphqlDebug {
			http.Handle("/query", corsHandler.Handler(
				httplog.HandlerWithFormatter(
					httplog.DefaultLogFormatter,
					oidcHandler,
				)))
		} else {
			http.Handle("/", playground.Handler("bicki", "/query"))
			http.Handle("/query", corsHandler.Handler(
				httplog.HandlerWithFormatter(
					httplog.DefaultLogFormatterWithRequestHeader,
					oidcHandler,
				)))
		}

		fmt.Printf("listening on %s", config.ListenAddress)
		return http.ListenAndServe(config.ListenAddress, nil)
	},
}

func init() {
	rootCmd.AddCommand(graphqlCmd)

	graphqlCmd.Flags().BoolVar(
		&graphqlDebug,
		"debug",
		true,
		"debug enabled?",
	)
	graphqlCmd.Flags().StringVarP(
		&config.ListenAddress,
		"addr",
		"a",
		":8081",
		"listen address",
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// graphqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// graphqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setDatabaseURI() {
	if config.DatabaseURI == "" {
		config.DatabaseURI = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			config.DatabaseUser,
			url.PathEscape(config.DatabasePassword),
			config.DatabaseHost,
			config.DatabasePort,
			config.DatabaseName,
		)
	}
}

func openDB(databaseURL string) *ent.Client {
	db, e := sql.Open("pgx", databaseURL)
	if e != nil {
		log.Fatalln(e.Error())
	}
	driver := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(driver))
}

func NewDefaultServer(es graphql.ExecutableSchema) *handler.Server {
	srv := handler.New(es)

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return srv
}
