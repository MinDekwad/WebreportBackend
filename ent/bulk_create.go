// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/bulk"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BulkCreate is the builder for creating a Bulk entity.
type BulkCreate struct {
	config
	mutation *BulkMutation
	hooks    []Hook
}

// SetBulkCreditSameday sets the "bulkCreditSameday" field.
func (bc *BulkCreate) SetBulkCreditSameday(f float64) *BulkCreate {
	bc.mutation.SetBulkCreditSameday(f)
	return bc
}

// SetNillableBulkCreditSameday sets the "bulkCreditSameday" field if the given value is not nil.
func (bc *BulkCreate) SetNillableBulkCreditSameday(f *float64) *BulkCreate {
	if f != nil {
		bc.SetBulkCreditSameday(*f)
	}
	return bc
}

// SetBulkCreditSamedayFee sets the "bulkCreditSamedayFee" field.
func (bc *BulkCreate) SetBulkCreditSamedayFee(f float64) *BulkCreate {
	bc.mutation.SetBulkCreditSamedayFee(f)
	return bc
}

// SetNillableBulkCreditSamedayFee sets the "bulkCreditSamedayFee" field if the given value is not nil.
func (bc *BulkCreate) SetNillableBulkCreditSamedayFee(f *float64) *BulkCreate {
	if f != nil {
		bc.SetBulkCreditSamedayFee(*f)
	}
	return bc
}

// SetTransfertobankaccount sets the "transfertobankaccount" field.
func (bc *BulkCreate) SetTransfertobankaccount(f float64) *BulkCreate {
	bc.mutation.SetTransfertobankaccount(f)
	return bc
}

// SetNillableTransfertobankaccount sets the "transfertobankaccount" field if the given value is not nil.
func (bc *BulkCreate) SetNillableTransfertobankaccount(f *float64) *BulkCreate {
	if f != nil {
		bc.SetTransfertobankaccount(*f)
	}
	return bc
}

// SetDateTime sets the "dateTime" field.
func (bc *BulkCreate) SetDateTime(t time.Time) *BulkCreate {
	bc.mutation.SetDateTime(t)
	return bc
}

// SetNillableDateTime sets the "dateTime" field if the given value is not nil.
func (bc *BulkCreate) SetNillableDateTime(t *time.Time) *BulkCreate {
	if t != nil {
		bc.SetDateTime(*t)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BulkCreate) SetID(i int) *BulkCreate {
	bc.mutation.SetID(i)
	return bc
}

// Mutation returns the BulkMutation object of the builder.
func (bc *BulkCreate) Mutation() *BulkMutation {
	return bc.mutation
}

// Save creates the Bulk in the database.
func (bc *BulkCreate) Save(ctx context.Context) (*Bulk, error) {
	var (
		err  error
		node *Bulk
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BulkMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BulkCreate) SaveX(ctx context.Context) *Bulk {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (bc *BulkCreate) check() error {
	return nil
}

func (bc *BulkCreate) sqlSave(ctx context.Context) (*Bulk, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
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

func (bc *BulkCreate) createSpec() (*Bulk, *sqlgraph.CreateSpec) {
	var (
		_node = &Bulk{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bulk.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bulk.FieldID,
			},
		}
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.BulkCreditSameday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: bulk.FieldBulkCreditSameday,
		})
		_node.BulkCreditSameday = &value
	}
	if value, ok := bc.mutation.BulkCreditSamedayFee(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: bulk.FieldBulkCreditSamedayFee,
		})
		_node.BulkCreditSamedayFee = &value
	}
	if value, ok := bc.mutation.Transfertobankaccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: bulk.FieldTransfertobankaccount,
		})
		_node.Transfertobankaccount = &value
	}
	if value, ok := bc.mutation.DateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: bulk.FieldDateTime,
		})
		_node.DateTime = &value
	}
	return _node, _spec
}

// BulkCreateBulk is the builder for creating many Bulk entities in bulk.
type BulkCreateBulk struct {
	config
	builders []*BulkCreate
}

// Save creates the Bulk entities in the database.
func (bcb *BulkCreateBulk) Save(ctx context.Context) ([]*Bulk, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Bulk, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BulkMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BulkCreateBulk) SaveX(ctx context.Context) []*Bulk {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
