// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/pendingkyc"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PendingkycCreate is the builder for creating a Pendingkyc entity.
type PendingkycCreate struct {
	config
	mutation *PendingkycMutation
	hooks    []Hook
}

// SetWalletID sets the "WalletID" field.
func (pc *PendingkycCreate) SetWalletID(s string) *PendingkycCreate {
	pc.mutation.SetWalletID(s)
	return pc
}

// SetNillableWalletID sets the "WalletID" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableWalletID(s *string) *PendingkycCreate {
	if s != nil {
		pc.SetWalletID(*s)
	}
	return pc
}

// SetName sets the "Name" field.
func (pc *PendingkycCreate) SetName(s string) *PendingkycCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "Name" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableName(s *string) *PendingkycCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetAgentID sets the "AgentID" field.
func (pc *PendingkycCreate) SetAgentID(s string) *PendingkycCreate {
	pc.mutation.SetAgentID(s)
	return pc
}

// SetNillableAgentID sets the "AgentID" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableAgentID(s *string) *PendingkycCreate {
	if s != nil {
		pc.SetAgentID(*s)
	}
	return pc
}

// SetAgentNameLastname sets the "AgentNameLastname" field.
func (pc *PendingkycCreate) SetAgentNameLastname(s string) *PendingkycCreate {
	pc.mutation.SetAgentNameLastname(s)
	return pc
}

// SetNillableAgentNameLastname sets the "AgentNameLastname" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableAgentNameLastname(s *string) *PendingkycCreate {
	if s != nil {
		pc.SetAgentNameLastname(*s)
	}
	return pc
}

// SetKYCDate sets the "KYCDate" field.
func (pc *PendingkycCreate) SetKYCDate(s string) *PendingkycCreate {
	pc.mutation.SetKYCDate(s)
	return pc
}

// SetNillableKYCDate sets the "KYCDate" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableKYCDate(s *string) *PendingkycCreate {
	if s != nil {
		pc.SetKYCDate(*s)
	}
	return pc
}

// SetDateGen sets the "DateGen" field.
func (pc *PendingkycCreate) SetDateGen(t time.Time) *PendingkycCreate {
	pc.mutation.SetDateGen(t)
	return pc
}

// SetNillableDateGen sets the "DateGen" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableDateGen(t *time.Time) *PendingkycCreate {
	if t != nil {
		pc.SetDateGen(*t)
	}
	return pc
}

// SetStatusGen sets the "StatusGen" field.
func (pc *PendingkycCreate) SetStatusGen(b bool) *PendingkycCreate {
	pc.mutation.SetStatusGen(b)
	return pc
}

// SetNillableStatusGen sets the "StatusGen" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableStatusGen(b *bool) *PendingkycCreate {
	if b != nil {
		pc.SetStatusGen(*b)
	}
	return pc
}

// SetPoint sets the "Point" field.
func (pc *PendingkycCreate) SetPoint(i int) *PendingkycCreate {
	pc.mutation.SetPoint(i)
	return pc
}

// SetNillablePoint sets the "Point" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillablePoint(i *int) *PendingkycCreate {
	if i != nil {
		pc.SetPoint(*i)
	}
	return pc
}

// SetFileimportID sets the "FileimportID" field.
func (pc *PendingkycCreate) SetFileimportID(i int) *PendingkycCreate {
	pc.mutation.SetFileimportID(i)
	return pc
}

// SetNillableFileimportID sets the "FileimportID" field if the given value is not nil.
func (pc *PendingkycCreate) SetNillableFileimportID(i *int) *PendingkycCreate {
	if i != nil {
		pc.SetFileimportID(*i)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PendingkycCreate) SetID(i int) *PendingkycCreate {
	pc.mutation.SetID(i)
	return pc
}

// Mutation returns the PendingkycMutation object of the builder.
func (pc *PendingkycCreate) Mutation() *PendingkycMutation {
	return pc.mutation
}

// Save creates the Pendingkyc in the database.
func (pc *PendingkycCreate) Save(ctx context.Context) (*Pendingkyc, error) {
	var (
		err  error
		node *Pendingkyc
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PendingkycMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PendingkycCreate) SaveX(ctx context.Context) *Pendingkyc {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (pc *PendingkycCreate) defaults() {
	if _, ok := pc.mutation.StatusGen(); !ok {
		v := pendingkyc.DefaultStatusGen
		pc.mutation.SetStatusGen(v)
	}
	if _, ok := pc.mutation.Point(); !ok {
		v := pendingkyc.DefaultPoint
		pc.mutation.SetPoint(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PendingkycCreate) check() error {
	if v, ok := pc.mutation.WalletID(); ok {
		if err := pendingkyc.WalletIDValidator(v); err != nil {
			return &ValidationError{Name: "WalletID", err: fmt.Errorf("ent: validator failed for field \"WalletID\": %w", err)}
		}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := pendingkyc.NameValidator(v); err != nil {
			return &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}
	if v, ok := pc.mutation.AgentID(); ok {
		if err := pendingkyc.AgentIDValidator(v); err != nil {
			return &ValidationError{Name: "AgentID", err: fmt.Errorf("ent: validator failed for field \"AgentID\": %w", err)}
		}
	}
	if v, ok := pc.mutation.AgentNameLastname(); ok {
		if err := pendingkyc.AgentNameLastnameValidator(v); err != nil {
			return &ValidationError{Name: "AgentNameLastname", err: fmt.Errorf("ent: validator failed for field \"AgentNameLastname\": %w", err)}
		}
	}
	if v, ok := pc.mutation.KYCDate(); ok {
		if err := pendingkyc.KYCDateValidator(v); err != nil {
			return &ValidationError{Name: "KYCDate", err: fmt.Errorf("ent: validator failed for field \"KYCDate\": %w", err)}
		}
	}
	return nil
}

func (pc *PendingkycCreate) sqlSave(ctx context.Context) (*Pendingkyc, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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

func (pc *PendingkycCreate) createSpec() (*Pendingkyc, *sqlgraph.CreateSpec) {
	var (
		_node = &Pendingkyc{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pendingkyc.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pendingkyc.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.WalletID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pendingkyc.FieldWalletID,
		})
		_node.WalletID = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pendingkyc.FieldName,
		})
		_node.Name = value
	}
	if value, ok := pc.mutation.AgentID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pendingkyc.FieldAgentID,
		})
		_node.AgentID = &value
	}
	if value, ok := pc.mutation.AgentNameLastname(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pendingkyc.FieldAgentNameLastname,
		})
		_node.AgentNameLastname = &value
	}
	if value, ok := pc.mutation.KYCDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pendingkyc.FieldKYCDate,
		})
		_node.KYCDate = &value
	}
	if value, ok := pc.mutation.DateGen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: pendingkyc.FieldDateGen,
		})
		_node.DateGen = &value
	}
	if value, ok := pc.mutation.StatusGen(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: pendingkyc.FieldStatusGen,
		})
		_node.StatusGen = value
	}
	if value, ok := pc.mutation.Point(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pendingkyc.FieldPoint,
		})
		_node.Point = value
	}
	if value, ok := pc.mutation.FileimportID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: pendingkyc.FieldFileimportID,
		})
		_node.FileimportID = &value
	}
	return _node, _spec
}

// PendingkycCreateBulk is the builder for creating many Pendingkyc entities in bulk.
type PendingkycCreateBulk struct {
	config
	builders []*PendingkycCreate
}

// Save creates the Pendingkyc entities in the database.
func (pcb *PendingkycCreateBulk) Save(ctx context.Context) ([]*Pendingkyc, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Pendingkyc, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PendingkycMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PendingkycCreateBulk) SaveX(ctx context.Context) []*Pendingkyc {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
