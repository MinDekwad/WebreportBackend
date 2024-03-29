// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/transactionfactoritemtmp"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionfactoritemtmpCreate is the builder for creating a Transactionfactoritemtmp entity.
type TransactionfactoritemtmpCreate struct {
	config
	mutation *TransactionfactoritemtmpMutation
	hooks    []Hook
}

// SetMin sets the "Min" field.
func (tc *TransactionfactoritemtmpCreate) SetMin(f float64) *TransactionfactoritemtmpCreate {
	tc.mutation.SetMin(f)
	return tc
}

// SetNillableMin sets the "Min" field if the given value is not nil.
func (tc *TransactionfactoritemtmpCreate) SetNillableMin(f *float64) *TransactionfactoritemtmpCreate {
	if f != nil {
		tc.SetMin(*f)
	}
	return tc
}

// SetMax sets the "Max" field.
func (tc *TransactionfactoritemtmpCreate) SetMax(f float64) *TransactionfactoritemtmpCreate {
	tc.mutation.SetMax(f)
	return tc
}

// SetNillableMax sets the "Max" field if the given value is not nil.
func (tc *TransactionfactoritemtmpCreate) SetNillableMax(f *float64) *TransactionfactoritemtmpCreate {
	if f != nil {
		tc.SetMax(*f)
	}
	return tc
}

// SetRank sets the "Rank" field.
func (tc *TransactionfactoritemtmpCreate) SetRank(i int) *TransactionfactoritemtmpCreate {
	tc.mutation.SetRank(i)
	return tc
}

// SetNillableRank sets the "Rank" field if the given value is not nil.
func (tc *TransactionfactoritemtmpCreate) SetNillableRank(i *int) *TransactionfactoritemtmpCreate {
	if i != nil {
		tc.SetRank(*i)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionfactoritemtmpCreate) SetID(i int) *TransactionfactoritemtmpCreate {
	tc.mutation.SetID(i)
	return tc
}

// Mutation returns the TransactionfactoritemtmpMutation object of the builder.
func (tc *TransactionfactoritemtmpCreate) Mutation() *TransactionfactoritemtmpMutation {
	return tc.mutation
}

// Save creates the Transactionfactoritemtmp in the database.
func (tc *TransactionfactoritemtmpCreate) Save(ctx context.Context) (*Transactionfactoritemtmp, error) {
	var (
		err  error
		node *Transactionfactoritemtmp
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionfactoritemtmpMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			node, err = tc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionfactoritemtmpCreate) SaveX(ctx context.Context) *Transactionfactoritemtmp {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionfactoritemtmpCreate) check() error {
	return nil
}

func (tc *TransactionfactoritemtmpCreate) sqlSave(ctx context.Context) (*Transactionfactoritemtmp, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	if _node.ID == 0 {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	return _node, nil
}

func (tc *TransactionfactoritemtmpCreate) createSpec() (*Transactionfactoritemtmp, *sqlgraph.CreateSpec) {
	var (
		_node = &Transactionfactoritemtmp{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transactionfactoritemtmp.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactionfactoritemtmp.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.Min(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transactionfactoritemtmp.FieldMin,
		})
		_node.Min = value
	}
	if value, ok := tc.mutation.Max(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: transactionfactoritemtmp.FieldMax,
		})
		_node.Max = value
	}
	if value, ok := tc.mutation.Rank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactoritemtmp.FieldRank,
		})
		_node.Rank = value
	}
	return _node, _spec
}

// TransactionfactoritemtmpCreateBulk is the builder for creating many Transactionfactoritemtmp entities in bulk.
type TransactionfactoritemtmpCreateBulk struct {
	config
	builders []*TransactionfactoritemtmpCreate
}

// Save creates the Transactionfactoritemtmp entities in the database.
func (tcb *TransactionfactoritemtmpCreateBulk) Save(ctx context.Context) ([]*Transactionfactoritemtmp, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transactionfactoritemtmp, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionfactoritemtmpMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				if nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionfactoritemtmpCreateBulk) SaveX(ctx context.Context) []*Transactionfactoritemtmp {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
