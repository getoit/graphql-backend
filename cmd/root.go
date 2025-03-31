/*
Copyright © 2025 Darko Luketic <info@icod.de>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/dlukt/graphql-backend-starter/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	useSQLite = false
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "starter",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.starter.yaml)")

	rootCmd.PersistentFlags().BoolVar(&useSQLite, "sqlite", false, "Use SQLite database")
	rootCmd.PersistentFlags().StringVar(
		&config.DatabaseURI,
		"db_uri",
		"",
		"postgres connection string (e.g. postgres://root:password@localhost:port/?sslmode=disable)",
	)
	rootCmd.PersistentFlags().StringVar(
		&config.DatabaseUser,
		"db_user",
		"dev",
		"postgres user",
	)
	rootCmd.PersistentFlags().StringVar(
		&config.DatabasePassword,
		"db_password",
		"dev",
		"postgres password",
	)
	rootCmd.PersistentFlags().StringVar(
		&config.DatabaseName,
		"db_name",
		"",
		"postgres database name",
	)
	rootCmd.PersistentFlags().StringVar(
		&config.DatabaseHost,
		"db_host",
		"127.0.0.1",
		"postgres host",
	)
	rootCmd.PersistentFlags().StringVar(
		&config.DatabasePort,
		"db_port",
		"5432",
		"postgres port",
	)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".blogs" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".blogs")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
