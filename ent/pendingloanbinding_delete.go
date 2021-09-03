// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/pendingloanbinding"
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PendingloanbindingDelete is the builder for deleting a Pendingloanbinding entity.
type PendingloanbindingDelete struct {
	config
	hooks    []Hook
	mutation *PendingloanbindingMutation
}

// Where adds a new predicate to the PendingloanbindingDelete builder.
func (pd *PendingloanbindingDelete) Where(ps ...predicate.Pendingloanbinding) *PendingloanbindingDelete {
	pd.mutation.predicates = append(pd.mutation.predicates, ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PendingloanbindingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PendingloanbindingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PendingloanbindingDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PendingloanbindingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: pendingloanbinding.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pendingloanbinding.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PendingloanbindingDeleteOne is the builder for deleting a single Pendingloanbinding entity.
type PendingloanbindingDeleteOne struct {
	pd *PendingloanbindingDelete
}

// Exec executes the deletion query.
func (pdo *PendingloanbindingDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pendingloanbinding.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PendingloanbindingDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
