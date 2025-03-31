package rules

import (
	"context"

	"github.com/dlukt/graphql-backend-starter/ent"
	"github.com/dlukt/graphql-backend-starter/ent/privacy"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
	"github.com/dlukt/graphql-backend-starter/rules/claims"
)

func ProfileDefaultMutationRule() privacy.MutationRule {
	return privacy.ProfileMutationRuleFunc(func(ctx context.Context, m *ent.ProfileMutation) error {
		claimSubject := claims.SubFromContext(ctx)
		if claimSubject == "" {
			return privacy.Denyf("no sub in context")
		}

		switch m.Op() {
		case ent.OpCreate:
			if claimSubject == "" {
				return privacy.Denyf("unauthenticated")
			}
			return privacy.Allow
		case ent.OpUpdate:
			sub, x := m.Sub()
			if !x {
				// should never happen
				return privacy.Denyf("sub of profile not present")
			}
			if sub != claimSubject {
				return privacy.Denyf("unauthorized to edit this profile")
			}
			return privacy.Allow
		case ent.OpDeleteOne:
			sub, x := m.Sub()
			if !x {
				// should never happen
				return privacy.Denyf("sub not present")
			}
			if sub != claimSubject {
				return privacy.Denyf("unauthorized to delete this profile")
			}
		default:
			return privacy.Skip
		}
		return privacy.Skip
	})
}

func ProfileCreateIfNotExists() privacy.QueryRule {
	return privacy.ProfileQueryRuleFunc(func(ctx context.Context, q *ent.ProfileQuery) error {
		c := claims.FromContext(ctx)
		if c == nil {
			return privacy.Skip
		}
		claimSubject := c.Sub
		if claimSubject == "" {
			return privacy.Skip
		}
		client := ent.FromContext(ctx)
		if client == nil {
			return privacy.Skip
		}

		allow := privacy.DecisionContext(ctx, privacy.Allow)
		if cnt := client.Profile.Query().Where(profile.Sub(claimSubject)).CountX(allow); cnt == 0 {
			client.Profile.Create().SetSub(claimSubject).SetName(c.PreferredUsername).ExecX(allow)
		}
		return privacy.Skip
	})
}
