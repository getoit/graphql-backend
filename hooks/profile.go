package hooks

import (
	"context"
	"errors"

	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/ent/hook"
	"github.com/dlukt/graphql-backend-starter/rules/claims"
)

func ProfileCreateHook(next ent.Mutator) ent.Mutator {
	return hook.ProfileFunc(func(ctx context.Context, m *ent.ProfileMutation) (ent.Value, error) {
		sub := claims.SubFromContext(ctx)
		if sub == "" {
			return nil, errors.New("no subject in token")
		}
		m.SetSub(sub)
		return next.Mutate(ctx, m)
	})
}
