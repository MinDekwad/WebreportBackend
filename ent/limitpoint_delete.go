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

// LimitpointDelete is the builder for deleting a Limitpoint entity.
type LimitpointDelete struct {
	config
	hooks    []Hook
	mutation *LimitpointMutation
}

// Where adds a new predicate to the LimitpointDelete builder.
func (ld *LimitpointDelete) Where(ps ...predicate.Limitpoint) *LimitpointDelete {
	ld.mutation.predicates = append(ld.mutation.predicates, ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LimitpointDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ld.hooks) == 0 {
		affected, err = ld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LimitpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ld.mutation = mutation
			affected, err = ld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ld.hooks) - 1; i >= 0; i-- {
			mut = ld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LimitpointDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LimitpointDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: limitpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: limitpoint.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
}

// LimitpointDeleteOne is the builder for deleting a single Limitpoint entity.
type LimitpointDeleteOne struct {
	ld *LimitpointDelete
}

// Exec executes the deletion query.
func (ldo *LimitpointDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{limitpoint.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LimitpointDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}
