// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/userwallet"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserwalletDelete is the builder for deleting a Userwallet entity.
type UserwalletDelete struct {
	config
	hooks    []Hook
	mutation *UserwalletMutation
}

// Where adds a new predicate to the UserwalletDelete builder.
func (ud *UserwalletDelete) Where(ps ...predicate.Userwallet) *UserwalletDelete {
	ud.mutation.predicates = append(ud.mutation.predicates, ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UserwalletDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ud.hooks) == 0 {
		affected, err = ud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserwalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ud.mutation = mutation
			affected, err = ud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ud.hooks) - 1; i >= 0; i-- {
			mut = ud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UserwalletDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UserwalletDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userwallet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userwallet.FieldID,
			},
		},
	}
	if ps := ud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ud.driver, _spec)
}

// UserwalletDeleteOne is the builder for deleting a single Userwallet entity.
type UserwalletDeleteOne struct {
	ud *UserwalletDelete
}

// Exec executes the deletion query.
func (udo *UserwalletDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userwallet.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserwalletDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
