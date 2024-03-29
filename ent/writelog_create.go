// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"go-api-report2/ent/writelog"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WritelogCreate is the builder for creating a Writelog entity.
type WritelogCreate struct {
	config
	mutation *WritelogMutation
	hooks    []Hook
}

// SetAdminID sets the "admin_id" field.
func (wc *WritelogCreate) SetAdminID(i int) *WritelogCreate {
	wc.mutation.SetAdminID(i)
	return wc
}

// SetResource sets the "resource" field.
func (wc *WritelogCreate) SetResource(s string) *WritelogCreate {
	wc.mutation.SetResource(s)
	return wc
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (wc *WritelogCreate) SetNillableResource(s *string) *WritelogCreate {
	if s != nil {
		wc.SetResource(*s)
	}
	return wc
}

// SetActions sets the "actions" field.
func (wc *WritelogCreate) SetActions(s string) *WritelogCreate {
	wc.mutation.SetActions(s)
	return wc
}

// SetNillableActions sets the "actions" field if the given value is not nil.
func (wc *WritelogCreate) SetNillableActions(s *string) *WritelogCreate {
	if s != nil {
		wc.SetActions(*s)
	}
	return wc
}

// SetCreatedAt sets the "created_at" field.
func (wc *WritelogCreate) SetCreatedAt(t time.Time) *WritelogCreate {
	wc.mutation.SetCreatedAt(t)
	return wc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wc *WritelogCreate) SetNillableCreatedAt(t *time.Time) *WritelogCreate {
	if t != nil {
		wc.SetCreatedAt(*t)
	}
	return wc
}

// SetID sets the "id" field.
func (wc *WritelogCreate) SetID(i int) *WritelogCreate {
	wc.mutation.SetID(i)
	return wc
}

// Mutation returns the WritelogMutation object of the builder.
func (wc *WritelogCreate) Mutation() *WritelogMutation {
	return wc.mutation
}

// Save creates the Writelog in the database.
func (wc *WritelogCreate) Save(ctx context.Context) (*Writelog, error) {
	var (
		err  error
		node *Writelog
	)
	if len(wc.hooks) == 0 {
		if err = wc.check(); err != nil {
			return nil, err
		}
		node, err = wc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WritelogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wc.check(); err != nil {
				return nil, err
			}
			wc.mutation = mutation
			node, err = wc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wc.hooks) - 1; i >= 0; i-- {
			mut = wc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WritelogCreate) SaveX(ctx context.Context) *Writelog {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (wc *WritelogCreate) check() error {
	if _, ok := wc.mutation.AdminID(); !ok {
		return &ValidationError{Name: "admin_id", err: errors.New("ent: missing required field \"admin_id\"")}
	}
	if v, ok := wc.mutation.Resource(); ok {
		if err := writelog.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf("ent: validator failed for field \"resource\": %w", err)}
		}
	}
	if v, ok := wc.mutation.Actions(); ok {
		if err := writelog.ActionsValidator(v); err != nil {
			return &ValidationError{Name: "actions", err: fmt.Errorf("ent: validator failed for field \"actions\": %w", err)}
		}
	}
	return nil
}

func (wc *WritelogCreate) sqlSave(ctx context.Context) (*Writelog, error) {
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
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

func (wc *WritelogCreate) createSpec() (*Writelog, *sqlgraph.CreateSpec) {
	var (
		_node = &Writelog{config: wc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: writelog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: writelog.FieldID,
			},
		}
	)
	if id, ok := wc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wc.mutation.AdminID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: writelog.FieldAdminID,
		})
		_node.AdminID = value
	}
	if value, ok := wc.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: writelog.FieldResource,
		})
		_node.Resource = &value
	}
	if value, ok := wc.mutation.Actions(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: writelog.FieldActions,
		})
		_node.Actions = &value
	}
	if value, ok := wc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: writelog.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	return _node, _spec
}

// WritelogCreateBulk is the builder for creating many Writelog entities in bulk.
type WritelogCreateBulk struct {
	config
	builders []*WritelogCreate
}

// Save creates the Writelog entities in the database.
func (wcb *WritelogCreateBulk) Save(ctx context.Context) ([]*Writelog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Writelog, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WritelogMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WritelogCreateBulk) SaveX(ctx context.Context) []*Writelog {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
