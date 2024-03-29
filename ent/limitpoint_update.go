// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/limitpoint"
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LimitpointUpdate is the builder for updating Limitpoint entities.
type LimitpointUpdate struct {
	config
	hooks    []Hook
	mutation *LimitpointMutation
}

// Where adds a new predicate for the LimitpointUpdate builder.
func (lu *LimitpointUpdate) Where(ps ...predicate.Limitpoint) *LimitpointUpdate {
	lu.mutation.predicates = append(lu.mutation.predicates, ps...)
	return lu
}

// SetLimitPoint sets the "LimitPoint" field.
func (lu *LimitpointUpdate) SetLimitPoint(i int) *LimitpointUpdate {
	lu.mutation.ResetLimitPoint()
	lu.mutation.SetLimitPoint(i)
	return lu
}

// SetNillableLimitPoint sets the "LimitPoint" field if the given value is not nil.
func (lu *LimitpointUpdate) SetNillableLimitPoint(i *int) *LimitpointUpdate {
	if i != nil {
		lu.SetLimitPoint(*i)
	}
	return lu
}

// AddLimitPoint adds i to the "LimitPoint" field.
func (lu *LimitpointUpdate) AddLimitPoint(i int) *LimitpointUpdate {
	lu.mutation.AddLimitPoint(i)
	return lu
}

// ClearLimitPoint clears the value of the "LimitPoint" field.
func (lu *LimitpointUpdate) ClearLimitPoint() *LimitpointUpdate {
	lu.mutation.ClearLimitPoint()
	return lu
}

// Mutation returns the LimitpointMutation object of the builder.
func (lu *LimitpointUpdate) Mutation() *LimitpointMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LimitpointUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LimitpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LimitpointUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LimitpointUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LimitpointUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LimitpointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   limitpoint.Table,
			Columns: limitpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: limitpoint.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.LimitPoint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: limitpoint.FieldLimitPoint,
		})
	}
	if value, ok := lu.mutation.AddedLimitPoint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: limitpoint.FieldLimitPoint,
		})
	}
	if lu.mutation.LimitPointCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: limitpoint.FieldLimitPoint,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{limitpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LimitpointUpdateOne is the builder for updating a single Limitpoint entity.
type LimitpointUpdateOne struct {
	config
	hooks    []Hook
	mutation *LimitpointMutation
}

// SetLimitPoint sets the "LimitPoint" field.
func (luo *LimitpointUpdateOne) SetLimitPoint(i int) *LimitpointUpdateOne {
	luo.mutation.ResetLimitPoint()
	luo.mutation.SetLimitPoint(i)
	return luo
}

// SetNillableLimitPoint sets the "LimitPoint" field if the given value is not nil.
func (luo *LimitpointUpdateOne) SetNillableLimitPoint(i *int) *LimitpointUpdateOne {
	if i != nil {
		luo.SetLimitPoint(*i)
	}
	return luo
}

// AddLimitPoint adds i to the "LimitPoint" field.
func (luo *LimitpointUpdateOne) AddLimitPoint(i int) *LimitpointUpdateOne {
	luo.mutation.AddLimitPoint(i)
	return luo
}

// ClearLimitPoint clears the value of the "LimitPoint" field.
func (luo *LimitpointUpdateOne) ClearLimitPoint() *LimitpointUpdateOne {
	luo.mutation.ClearLimitPoint()
	return luo
}

// Mutation returns the LimitpointMutation object of the builder.
func (luo *LimitpointUpdateOne) Mutation() *LimitpointMutation {
	return luo.mutation
}

// Save executes the query and returns the updated Limitpoint entity.
func (luo *LimitpointUpdateOne) Save(ctx context.Context) (*Limitpoint, error) {
	var (
		err  error
		node *Limitpoint
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LimitpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			mut = luo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, luo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LimitpointUpdateOne) SaveX(ctx context.Context) *Limitpoint {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LimitpointUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LimitpointUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LimitpointUpdateOne) sqlSave(ctx context.Context) (_node *Limitpoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   limitpoint.Table,
			Columns: limitpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: limitpoint.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Limitpoint.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.LimitPoint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: limitpoint.FieldLimitPoint,
		})
	}
	if value, ok := luo.mutation.AddedLimitPoint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: limitpoint.FieldLimitPoint,
		})
	}
	if luo.mutation.LimitPointCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: limitpoint.FieldLimitPoint,
		})
	}
	_node = &Limitpoint{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{limitpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
