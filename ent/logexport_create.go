// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/logexport"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LogexportCreate is the builder for creating a Logexport entity.
type LogexportCreate struct {
	config
	mutation *LogexportMutation
	hooks    []Hook
}

// SetUserName sets the "UserName" field.
func (lc *LogexportCreate) SetUserName(s string) *LogexportCreate {
	lc.mutation.SetUserName(s)
	return lc
}

// SetNillableUserName sets the "UserName" field if the given value is not nil.
func (lc *LogexportCreate) SetNillableUserName(s *string) *LogexportCreate {
	if s != nil {
		lc.SetUserName(*s)
	}
	return lc
}

// SetFileName sets the "FileName" field.
func (lc *LogexportCreate) SetFileName(s string) *LogexportCreate {
	lc.mutation.SetFileName(s)
	return lc
}

// SetNillableFileName sets the "FileName" field if the given value is not nil.
func (lc *LogexportCreate) SetNillableFileName(s *string) *LogexportCreate {
	if s != nil {
		lc.SetFileName(*s)
	}
	return lc
}

// SetExportDate sets the "ExportDate" field.
func (lc *LogexportCreate) SetExportDate(t time.Time) *LogexportCreate {
	lc.mutation.SetExportDate(t)
	return lc
}

// SetNillableExportDate sets the "ExportDate" field if the given value is not nil.
func (lc *LogexportCreate) SetNillableExportDate(t *time.Time) *LogexportCreate {
	if t != nil {
		lc.SetExportDate(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LogexportCreate) SetID(i int) *LogexportCreate {
	lc.mutation.SetID(i)
	return lc
}

// Mutation returns the LogexportMutation object of the builder.
func (lc *LogexportCreate) Mutation() *LogexportMutation {
	return lc.mutation
}

// Save creates the Logexport in the database.
func (lc *LogexportCreate) Save(ctx context.Context) (*Logexport, error) {
	var (
		err  error
		node *Logexport
	)
	if len(lc.hooks) == 0 {
		if err = lc.check(); err != nil {
			return nil, err
		}
		node, err = lc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LogexportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = lc.check(); err != nil {
				return nil, err
			}
			lc.mutation = mutation
			node, err = lc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lc.hooks) - 1; i >= 0; i-- {
			mut = lc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LogexportCreate) SaveX(ctx context.Context) *Logexport {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (lc *LogexportCreate) check() error {
	if v, ok := lc.mutation.UserName(); ok {
		if err := logexport.UserNameValidator(v); err != nil {
			return &ValidationError{Name: "UserName", err: fmt.Errorf("ent: validator failed for field \"UserName\": %w", err)}
		}
	}
	if v, ok := lc.mutation.FileName(); ok {
		if err := logexport.FileNameValidator(v); err != nil {
			return &ValidationError{Name: "FileName", err: fmt.Errorf("ent: validator failed for field \"FileName\": %w", err)}
		}
	}
	return nil
}

func (lc *LogexportCreate) sqlSave(ctx context.Context) (*Logexport, error) {
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
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

func (lc *LogexportCreate) createSpec() (*Logexport, *sqlgraph.CreateSpec) {
	var (
		_node = &Logexport{config: lc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: logexport.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logexport.FieldID,
			},
		}
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.UserName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logexport.FieldUserName,
		})
		_node.UserName = value
	}
	if value, ok := lc.mutation.FileName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: logexport.FieldFileName,
		})
		_node.FileName = value
	}
	if value, ok := lc.mutation.ExportDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: logexport.FieldExportDate,
		})
		_node.ExportDate = value
	}
	return _node, _spec
}

// LogexportCreateBulk is the builder for creating many Logexport entities in bulk.
type LogexportCreateBulk struct {
	config
	builders []*LogexportCreate
}

// Save creates the Logexport entities in the database.
func (lcb *LogexportCreateBulk) Save(ctx context.Context) ([]*Logexport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Logexport, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LogexportMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LogexportCreateBulk) SaveX(ctx context.Context) []*Logexport {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}