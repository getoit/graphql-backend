// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dlukt/graphql-backend-starter/ent/predicate"
	"github.com/dlukt/graphql-backend-starter/ent/profile"
	"github.com/rs/xid"
)

// ProfileQuery is the builder for querying Profile entities.
type ProfileQuery struct {
	config
	ctx        *QueryContext
	order      []profile.OrderOption
	inters     []Interceptor
	predicates []predicate.Profile
	loadTotal  []func(context.Context, []*Profile) error
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfileQuery builder.
func (pq *ProfileQuery) Where(ps ...predicate.Profile) *ProfileQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProfileQuery) Limit(limit int) *ProfileQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProfileQuery) Offset(offset int) *ProfileQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProfileQuery) Unique(unique bool) *ProfileQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProfileQuery) Order(o ...profile.OrderOption) *ProfileQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// First returns the first Profile entity from the query.
// Returns a *NotFoundError when no Profile was found.
func (pq *ProfileQuery) First(ctx context.Context) (*Profile, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{profile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProfileQuery) FirstX(ctx context.Context) *Profile {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Profile ID from the query.
// Returns a *NotFoundError when no Profile ID was found.
func (pq *ProfileQuery) FirstID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{profile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProfileQuery) FirstIDX(ctx context.Context) xid.ID {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Profile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Profile entity is found.
// Returns a *NotFoundError when no Profile entities are found.
func (pq *ProfileQuery) Only(ctx context.Context) (*Profile, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{profile.Label}
	default:
		return nil, &NotSingularError{profile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProfileQuery) OnlyX(ctx context.Context) *Profile {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Profile ID in the query.
// Returns a *NotSingularError when more than one Profile ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProfileQuery) OnlyID(ctx context.Context) (id xid.ID, err error) {
	var ids []xid.ID
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{profile.Label}
	default:
		err = &NotSingularError{profile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProfileQuery) OnlyIDX(ctx context.Context) xid.ID {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Profiles.
func (pq *ProfileQuery) All(ctx context.Context) ([]*Profile, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryAll)
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Profile, *ProfileQuery]()
	return withInterceptors[[]*Profile](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProfileQuery) AllX(ctx context.Context) []*Profile {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Profile IDs.
func (pq *ProfileQuery) IDs(ctx context.Context) (ids []xid.ID, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryIDs)
	if err = pq.Select(profile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProfileQuery) IDsX(ctx context.Context) []xid.ID {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryCount)
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProfileQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProfileQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, ent.OpQueryExist)
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProfileQuery) Clone() *ProfileQuery {
	if pq == nil {
		return nil
	}
	return &ProfileQuery{
		config:     pq.config,
		ctx:        pq.ctx.Clone(),
		order:      append([]profile.OrderOption{}, pq.order...),
		inters:     append([]Interceptor{}, pq.inters...),
		predicates: append([]predicate.Profile{}, pq.predicates...),
		// clone intermediate query.
		sql:       pq.sql.Clone(),
		path:      pq.path,
		modifiers: append([]func(*sql.Selector){}, pq.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Profile.Query().
//		GroupBy(profile.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProfileQuery) GroupBy(field string, fields ...string) *ProfileGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProfileGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = profile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Profile.Query().
//		Select(profile.FieldCreateTime).
//		Scan(ctx, &v)
func (pq *ProfileQuery) Select(fields ...string) *ProfileSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProfileSelect{ProfileQuery: pq}
	sbuild.label = profile.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfileSelect configured with the given aggregations.
func (pq *ProfileQuery) Aggregate(fns ...AggregateFunc) *ProfileSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !profile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	if profile.Policy == nil {
		return errors.New("ent: uninitialized profile.Policy (forgotten import ent/runtime?)")
	}
	if err := profile.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *ProfileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Profile, error) {
	var (
		nodes = []*Profile{}
		_spec = pq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Profile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Profile{config: pq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProfileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProfileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(profile.Table, profile.Columns, sqlgraph.NewFieldSpec(profile.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for i := range fields {
			if fields[i] != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProfileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(profile.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = profile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *ProfileQuery) Modify(modifiers ...func(s *sql.Selector)) *ProfileSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// ProfileGroupBy is the group-by builder for Profile entities.
type ProfileGroupBy struct {
	selector
	build *ProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProfileGroupBy) Aggregate(fns ...AggregateFunc) *ProfileGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileQuery, *ProfileGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProfileGroupBy) sqlScan(ctx context.Context, root *ProfileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfileSelect is the builder for selecting fields of Profile entities.
type ProfileSelect struct {
	*ProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProfileSelect) Aggregate(fns ...AggregateFunc) *ProfileSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfileQuery, *ProfileSelect](ctx, ps.ProfileQuery, ps, ps.inters, v)
}

func (ps *ProfileSelect) sqlScan(ctx context.Context, root *ProfileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *ProfileSelect) Modify(modifiers ...func(s *sql.Selector)) *ProfileSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
