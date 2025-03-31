package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/graph/generated"
	"github.com/rs/xid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id xid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []xid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Profiles is the resolver for the profiles field.
func (r *queryResolver) Profiles(ctx context.Context, after *entgql.Cursor[xid.ID], first *int, before *entgql.Cursor[xid.ID], last *int, orderBy *ent.ProfileOrder, where *ent.ProfileWhereInput) (*ent.ProfileConnection, error) {
	nc := ent.NewContext(ctx, r.client)
	return r.client.Profile.Query().Paginate(nc, after, first, before, last, ent.WithProfileOrder(orderBy), ent.WithProfileFilter(where.Filter))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
