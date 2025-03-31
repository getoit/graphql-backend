// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pr *ProfileQuery) CollectFields(ctx context.Context, satisfies ...string) (*ProfileQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pr, nil
	}
	if err := pr.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pr, nil
}

func (pr *ProfileQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(profile.Columns))
		selectedFields = []string{profile.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createTime":
			if _, ok := fieldSeen[profile.FieldCreateTime]; !ok {
				selectedFields = append(selectedFields, profile.FieldCreateTime)
				fieldSeen[profile.FieldCreateTime] = struct{}{}
			}
		case "updateTime":
			if _, ok := fieldSeen[profile.FieldUpdateTime]; !ok {
				selectedFields = append(selectedFields, profile.FieldUpdateTime)
				fieldSeen[profile.FieldUpdateTime] = struct{}{}
			}
		case "sub":
			if _, ok := fieldSeen[profile.FieldSub]; !ok {
				selectedFields = append(selectedFields, profile.FieldSub)
				fieldSeen[profile.FieldSub] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[profile.FieldName]; !ok {
				selectedFields = append(selectedFields, profile.FieldName)
				fieldSeen[profile.FieldName] = struct{}{}
			}
		case "gender":
			if _, ok := fieldSeen[profile.FieldGender]; !ok {
				selectedFields = append(selectedFields, profile.FieldGender)
				fieldSeen[profile.FieldGender] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pr.Select(selectedFields...)
	}
	return nil
}

type profilePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []ProfilePaginateOption
}

func newProfilePaginateArgs(rv map[string]any) *profilePaginateArgs {
	args := &profilePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &ProfileOrder{Field: &ProfileOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithProfileOrder(order))
			}
		case *ProfileOrder:
			if v != nil {
				args.opts = append(args.opts, WithProfileOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*ProfileWhereInput); ok {
		args.opts = append(args.opts, WithProfileFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
