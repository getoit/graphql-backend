package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/dlukt/graphql-backend-starter/ent/hook"
	"github.com/dlukt/graphql-backend-starter/ent/privacy"
	"github.com/dlukt/graphql-backend-starter/hooks"
	"github.com/dlukt/graphql-backend-starter/rules"
	"github.com/rs/xid"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").GoType(xid.ID{}).DefaultFunc(xid.New).Unique(),
		field.Time("create_time").Default(time.Now).Immutable().Annotations(entgql.OrderField("CREATE_TIME")),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now).Optional().Annotations(entgql.OrderField("UPDATE_TIME")),
		field.String("sub").Unique().MaxLen(36),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.String("gender"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return nil
}

func (Profile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Policy of the Profile.
func (Profile) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// privacyrules.PrintQueryToken(),
			rules.ProfileDefaultMutationRule(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			// privacyrules.PrintQueryToken(),
			rules.ProfileCreateIfNotExists(),
			privacy.AlwaysAllowRule(),
		},
	}
}

// Hooks of the Profile.
func (Profile) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(hooks.ProfileCreateHook, ent.OpCreate),
	}
}
