// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/userprofile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserprofileDelete is the builder for deleting a Userprofile entity.
type UserprofileDelete struct {
	config
	hooks    []Hook
	mutation *UserprofileMutation
}

// Where adds a new predicate to the UserprofileDelete builder.
func (ud *UserprofileDelete) Where(ps ...predicate.Userprofile) *UserprofileDelete {
	ud.mutation.predicates = append(ud.mutation.predicates, ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UserprofileDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ud.hooks) == 0 {
		affected, err = ud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserprofileMutation)
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
func (ud *UserprofileDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UserprofileDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: userprofile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userprofile.FieldID,
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

// UserprofileDeleteOne is the builder for deleting a single Userprofile entity.
type UserprofileDeleteOne struct {
	ud *UserprofileDelete
}

// Exec executes the deletion query.
func (udo *UserprofileDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userprofile.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserprofileDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
