// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/transactionfactorhistory"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionfactorhistoryDelete is the builder for deleting a Transactionfactorhistory entity.
type TransactionfactorhistoryDelete struct {
	config
	hooks    []Hook
	mutation *TransactionfactorhistoryMutation
}

// Where adds a new predicate to the TransactionfactorhistoryDelete builder.
func (td *TransactionfactorhistoryDelete) Where(ps ...predicate.Transactionfactorhistory) *TransactionfactorhistoryDelete {
	td.mutation.predicates = append(td.mutation.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TransactionfactorhistoryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionfactorhistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TransactionfactorhistoryDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TransactionfactorhistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: transactionfactorhistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactionfactorhistory.FieldID,
			},
		},
	}
	if ps := td.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TransactionfactorhistoryDeleteOne is the builder for deleting a single Transactionfactorhistory entity.
type TransactionfactorhistoryDeleteOne struct {
	td *TransactionfactorhistoryDelete
}

// Exec executes the deletion query.
func (tdo *TransactionfactorhistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{transactionfactorhistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TransactionfactorhistoryDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}