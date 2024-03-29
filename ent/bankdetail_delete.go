// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/bankdetail"
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BankdetailDelete is the builder for deleting a Bankdetail entity.
type BankdetailDelete struct {
	config
	hooks    []Hook
	mutation *BankdetailMutation
}

// Where adds a new predicate to the BankdetailDelete builder.
func (bd *BankdetailDelete) Where(ps ...predicate.Bankdetail) *BankdetailDelete {
	bd.mutation.predicates = append(bd.mutation.predicates, ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BankdetailDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(bd.hooks) == 0 {
		affected, err = bd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BankdetailMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bd.mutation = mutation
			affected, err = bd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bd.hooks) - 1; i >= 0; i-- {
			mut = bd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BankdetailDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BankdetailDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: bankdetail.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bankdetail.FieldID,
			},
		},
	}
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
}

// BankdetailDeleteOne is the builder for deleting a single Bankdetail entity.
type BankdetailDeleteOne struct {
	bd *BankdetailDelete
}

// Exec executes the deletion query.
func (bdo *BankdetailDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bankdetail.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BankdetailDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
