// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-api-report2/ent/occupationhistory"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OccupationhistoryCreate is the builder for creating a Occupationhistory entity.
type OccupationhistoryCreate struct {
	config
	mutation *OccupationhistoryMutation
	hooks    []Hook
}

// SetWalletID sets the "WalletID" field.
func (oc *OccupationhistoryCreate) SetWalletID(s string) *OccupationhistoryCreate {
	oc.mutation.SetWalletID(s)
	return oc
}

// SetOccupationName sets the "OccupationName" field.
func (oc *OccupationhistoryCreate) SetOccupationName(s string) *OccupationhistoryCreate {
	oc.mutation.SetOccupationName(s)
	return oc
}

// SetNillableOccupationName sets the "OccupationName" field if the given value is not nil.
func (oc *OccupationhistoryCreate) SetNillableOccupationName(s *string) *OccupationhistoryCreate {
	if s != nil {
		oc.SetOccupationName(*s)
	}
	return oc
}

// SetRankOccupation sets the "RankOccupation" field.
func (oc *OccupationhistoryCreate) SetRankOccupation(i int) *OccupationhistoryCreate {
	oc.mutation.SetRankOccupation(i)
	return oc
}

// SetNillableRankOccupation sets the "RankOccupation" field if the given value is not nil.
func (oc *OccupationhistoryCreate) SetNillableRankOccupation(i *int) *OccupationhistoryCreate {
	if i != nil {
		oc.SetRankOccupation(*i)
	}
	return oc
}

// SetDateCalRank sets the "DateCalRank" field.
func (oc *OccupationhistoryCreate) SetDateCalRank(t time.Time) *OccupationhistoryCreate {
	oc.mutation.SetDateCalRank(t)
	return oc
}

// SetNillableDateCalRank sets the "DateCalRank" field if the given value is not nil.
func (oc *OccupationhistoryCreate) SetNillableDateCalRank(t *time.Time) *OccupationhistoryCreate {
	if t != nil {
		oc.SetDateCalRank(*t)
	}
	return oc
}

// SetID sets the "id" field.
func (oc *OccupationhistoryCreate) SetID(i int) *OccupationhistoryCreate {
	oc.mutation.SetID(i)
	return oc
}

// Mutation returns the OccupationhistoryMutation object of the builder.
func (oc *OccupationhistoryCreate) Mutation() *OccupationhistoryMutation {
	return oc.mutation
}

// Save creates the Occupationhistory in the database.
func (oc *OccupationhistoryCreate) Save(ctx context.Context) (*Occupationhistory, error) {
	var (
		err  error
		node *Occupationhistory
	)
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OccupationhistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OccupationhistoryCreate) SaveX(ctx context.Context) *Occupationhistory {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (oc *OccupationhistoryCreate) check() error {
	if _, ok := oc.mutation.WalletID(); !ok {
		return &ValidationError{Name: "WalletID", err: errors.New("ent: missing required field \"WalletID\"")}
	}
	if v, ok := oc.mutation.WalletID(); ok {
		if err := occupationhistory.WalletIDValidator(v); err != nil {
			return &ValidationError{Name: "WalletID", err: fmt.Errorf("ent: validator failed for field \"WalletID\": %w", err)}
		}
	}
	if v, ok := oc.mutation.OccupationName(); ok {
		if err := occupationhistory.OccupationNameValidator(v); err != nil {
			return &ValidationError{Name: "OccupationName", err: fmt.Errorf("ent: validator failed for field \"OccupationName\": %w", err)}
		}
	}
	return nil
}

func (oc *OccupationhistoryCreate) sqlSave(ctx context.Context) (*Occupationhistory, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
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

func (oc *OccupationhistoryCreate) createSpec() (*Occupationhistory, *sqlgraph.CreateSpec) {
	var (
		_node = &Occupationhistory{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: occupationhistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: occupationhistory.FieldID,
			},
		}
	)
	if id, ok := oc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oc.mutation.WalletID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occupationhistory.FieldWalletID,
		})
		_node.WalletID = value
	}
	if value, ok := oc.mutation.OccupationName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occupationhistory.FieldOccupationName,
		})
		_node.OccupationName = value
	}
	if value, ok := oc.mutation.RankOccupation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: occupationhistory.FieldRankOccupation,
		})
		_node.RankOccupation = value
	}
	if value, ok := oc.mutation.DateCalRank(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: occupationhistory.FieldDateCalRank,
		})
		_node.DateCalRank = &value
	}
	return _node, _spec
}

// OccupationhistoryCreateBulk is the builder for creating many Occupationhistory entities in bulk.
type OccupationhistoryCreateBulk struct {
	config
	builders []*OccupationhistoryCreate
}

// Save creates the Occupationhistory entities in the database.
func (ocb *OccupationhistoryCreateBulk) Save(ctx context.Context) ([]*Occupationhistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Occupationhistory, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OccupationhistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OccupationhistoryCreateBulk) SaveX(ctx context.Context) []*Occupationhistory {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
