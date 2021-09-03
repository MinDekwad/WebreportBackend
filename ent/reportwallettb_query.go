// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/reportwallettb"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReportwallettbQuery is the builder for querying Reportwallettb entities.
type ReportwallettbQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	fields     []string
	predicates []predicate.Reportwallettb
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReportwallettbQuery builder.
func (rq *ReportwallettbQuery) Where(ps ...predicate.Reportwallettb) *ReportwallettbQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit adds a limit step to the query.
func (rq *ReportwallettbQuery) Limit(limit int) *ReportwallettbQuery {
	rq.limit = &limit
	return rq
}

// Offset adds an offset step to the query.
func (rq *ReportwallettbQuery) Offset(offset int) *ReportwallettbQuery {
	rq.offset = &offset
	return rq
}

// Order adds an order step to the query.
func (rq *ReportwallettbQuery) Order(o ...OrderFunc) *ReportwallettbQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// First returns the first Reportwallettb entity from the query.
// Returns a *NotFoundError when no Reportwallettb was found.
func (rq *ReportwallettbQuery) First(ctx context.Context) (*Reportwallettb, error) {
	nodes, err := rq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{reportwallettb.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *ReportwallettbQuery) FirstX(ctx context.Context) *Reportwallettb {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Reportwallettb ID from the query.
// Returns a *NotFoundError when no Reportwallettb ID was found.
func (rq *ReportwallettbQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{reportwallettb.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *ReportwallettbQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Reportwallettb entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Reportwallettb entity is not found.
// Returns a *NotFoundError when no Reportwallettb entities are found.
func (rq *ReportwallettbQuery) Only(ctx context.Context) (*Reportwallettb, error) {
	nodes, err := rq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{reportwallettb.Label}
	default:
		return nil, &NotSingularError{reportwallettb.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *ReportwallettbQuery) OnlyX(ctx context.Context) *Reportwallettb {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Reportwallettb ID in the query.
// Returns a *NotSingularError when exactly one Reportwallettb ID is not found.
// Returns a *NotFoundError when no entities are found.
func (rq *ReportwallettbQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = &NotSingularError{reportwallettb.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *ReportwallettbQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Reportwallettbs.
func (rq *ReportwallettbQuery) All(ctx context.Context) ([]*Reportwallettb, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rq *ReportwallettbQuery) AllX(ctx context.Context) []*Reportwallettb {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Reportwallettb IDs.
func (rq *ReportwallettbQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := rq.Select(reportwallettb.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *ReportwallettbQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *ReportwallettbQuery) Count(ctx context.Context) (int, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rq *ReportwallettbQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *ReportwallettbQuery) Exist(ctx context.Context) (bool, error) {
	if err := rq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *ReportwallettbQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReportwallettbQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *ReportwallettbQuery) Clone() *ReportwallettbQuery {
	if rq == nil {
		return nil
	}
	return &ReportwallettbQuery{
		config:     rq.config,
		limit:      rq.limit,
		offset:     rq.offset,
		order:      append([]OrderFunc{}, rq.order...),
		predicates: append([]predicate.Reportwallettb{}, rq.predicates...),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Walletid string `json:"walletid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Reportwallettb.Query().
//		GroupBy(reportwallettb.FieldWalletid).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rq *ReportwallettbQuery) GroupBy(field string, fields ...string) *ReportwallettbGroupBy {
	group := &ReportwallettbGroupBy{config: rq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Walletid string `json:"walletid,omitempty"`
//	}
//
//	client.Reportwallettb.Query().
//		Select(reportwallettb.FieldWalletid).
//		Scan(ctx, &v)
//
func (rq *ReportwallettbQuery) Select(field string, fields ...string) *ReportwallettbSelect {
	rq.fields = append([]string{field}, fields...)
	return &ReportwallettbSelect{ReportwallettbQuery: rq}
}

func (rq *ReportwallettbQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rq.fields {
		if !reportwallettb.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *ReportwallettbQuery) sqlAll(ctx context.Context) ([]*Reportwallettb, error) {
	var (
		nodes = []*Reportwallettb{}
		_spec = rq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Reportwallettb{config: rq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rq *ReportwallettbQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *ReportwallettbQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (rq *ReportwallettbQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reportwallettb.Table,
			Columns: reportwallettb.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reportwallettb.FieldID,
			},
		},
		From:   rq.sql,
		Unique: true,
	}
	if fields := rq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reportwallettb.FieldID)
		for i := range fields {
			if fields[i] != reportwallettb.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector, reportwallettb.ValidColumn)
			}
		}
	}
	return _spec
}

func (rq *ReportwallettbQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(reportwallettb.Table)
	selector := builder.Select(t1.Columns(reportwallettb.Columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(reportwallettb.Columns...)...)
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector, reportwallettb.ValidColumn)
	}
	if offset := rq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ReportwallettbGroupBy is the group-by builder for Reportwallettb entities.
type ReportwallettbGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *ReportwallettbGroupBy) Aggregate(fns ...AggregateFunc) *ReportwallettbGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the group-by query and scans the result into the given value.
func (rgb *ReportwallettbGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rgb.path(ctx)
	if err != nil {
		return err
	}
	rgb.sql = query
	return rgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := rgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) StringsX(ctx context.Context) []string {
	v, err := rgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) StringX(ctx context.Context) string {
	v, err := rgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) IntsX(ctx context.Context) []int {
	v, err := rgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) IntX(ctx context.Context) int {
	v, err := rgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := rgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) Float64X(ctx context.Context) float64 {
	v, err := rgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(rgb.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := rgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := rgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (rgb *ReportwallettbGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rgb *ReportwallettbGroupBy) BoolX(ctx context.Context) bool {
	v, err := rgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rgb *ReportwallettbGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rgb.fields {
		if !reportwallettb.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rgb *ReportwallettbGroupBy) sqlQuery() *sql.Selector {
	selector := rgb.sql
	columns := make([]string, 0, len(rgb.fields)+len(rgb.fns))
	columns = append(columns, rgb.fields...)
	for _, fn := range rgb.fns {
		columns = append(columns, fn(selector, reportwallettb.ValidColumn))
	}
	return selector.Select(columns...).GroupBy(rgb.fields...)
}

// ReportwallettbSelect is the builder for selecting fields of Reportwallettb entities.
type ReportwallettbSelect struct {
	*ReportwallettbQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rs *ReportwallettbSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	rs.sql = rs.ReportwallettbQuery.sqlQuery(ctx)
	return rs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (rs *ReportwallettbSelect) ScanX(ctx context.Context, v interface{}) {
	if err := rs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Strings(ctx context.Context) ([]string, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (rs *ReportwallettbSelect) StringsX(ctx context.Context) []string {
	v, err := rs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = rs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (rs *ReportwallettbSelect) StringX(ctx context.Context) string {
	v, err := rs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Ints(ctx context.Context) ([]int, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (rs *ReportwallettbSelect) IntsX(ctx context.Context) []int {
	v, err := rs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = rs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (rs *ReportwallettbSelect) IntX(ctx context.Context) int {
	v, err := rs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (rs *ReportwallettbSelect) Float64sX(ctx context.Context) []float64 {
	v, err := rs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = rs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (rs *ReportwallettbSelect) Float64X(ctx context.Context) float64 {
	v, err := rs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(rs.fields) > 1 {
		return nil, errors.New("ent: ReportwallettbSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := rs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (rs *ReportwallettbSelect) BoolsX(ctx context.Context) []bool {
	v, err := rs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (rs *ReportwallettbSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = rs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{reportwallettb.Label}
	default:
		err = fmt.Errorf("ent: ReportwallettbSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (rs *ReportwallettbSelect) BoolX(ctx context.Context) bool {
	v, err := rs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (rs *ReportwallettbSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rs.sqlQuery().Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rs *ReportwallettbSelect) sqlQuery() sql.Querier {
	selector := rs.sql
	selector.Select(selector.Columns(rs.fields...)...)
	return selector
}
