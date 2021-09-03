// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/fileimport"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileimportCreate is the builder for creating a Fileimport entity.
type FileimportCreate struct {
	config
	mutation *FileimportMutation
	hooks    []Hook
}

// SetFilename sets the "filename" field.
func (fc *FileimportCreate) SetFilename(s string) *FileimportCreate {
	fc.mutation.SetFilename(s)
	return fc
}

// SetNillableFilename sets the "filename" field if the given value is not nil.
func (fc *FileimportCreate) SetNillableFilename(s *string) *FileimportCreate {
	if s != nil {
		fc.SetFilename(*s)
	}
	return fc
}

// SetFiletype sets the "filetype" field.
func (fc *FileimportCreate) SetFiletype(s string) *FileimportCreate {
	fc.mutation.SetFiletype(s)
	return fc
}

// SetNillableFiletype sets the "filetype" field if the given value is not nil.
func (fc *FileimportCreate) SetNillableFiletype(s *string) *FileimportCreate {
	if s != nil {
		fc.SetFiletype(*s)
	}
	return fc
}

// SetImportdate sets the "importdate" field.
func (fc *FileimportCreate) SetImportdate(t time.Time) *FileimportCreate {
	fc.mutation.SetImportdate(t)
	return fc
}

// SetNillableImportdate sets the "importdate" field if the given value is not nil.
func (fc *FileimportCreate) SetNillableImportdate(t *time.Time) *FileimportCreate {
	if t != nil {
		fc.SetImportdate(*t)
	}
	return fc
}

// SetStatus sets the "status" field.
func (fc *FileimportCreate) SetStatus(s string) *FileimportCreate {
	fc.mutation.SetStatus(s)
	return fc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fc *FileimportCreate) SetNillableStatus(s *string) *FileimportCreate {
	if s != nil {
		fc.SetStatus(*s)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileimportCreate) SetID(i int) *FileimportCreate {
	fc.mutation.SetID(i)
	return fc
}

// Mutation returns the FileimportMutation object of the builder.
func (fc *FileimportCreate) Mutation() *FileimportMutation {
	return fc.mutation
}

// Save creates the Fileimport in the database.
func (fc *FileimportCreate) Save(ctx context.Context) (*Fileimport, error) {
	var (
		err  error
		node *Fileimport
	)
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileimportMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileimportCreate) SaveX(ctx context.Context) *Fileimport {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileimportCreate) check() error {
	if v, ok := fc.mutation.Filename(); ok {
		if err := fileimport.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf("ent: validator failed for field \"filename\": %w", err)}
		}
	}
	if v, ok := fc.mutation.Filetype(); ok {
		if err := fileimport.FiletypeValidator(v); err != nil {
			return &ValidationError{Name: "filetype", err: fmt.Errorf("ent: validator failed for field \"filetype\": %w", err)}
		}
	}
	if v, ok := fc.mutation.Status(); ok {
		if err := fileimport.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	return nil
}

func (fc *FileimportCreate) sqlSave(ctx context.Context) (*Fileimport, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
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

func (fc *FileimportCreate) createSpec() (*Fileimport, *sqlgraph.CreateSpec) {
	var (
		_node = &Fileimport{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fileimport.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileimport.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Filename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileimport.FieldFilename,
		})
		_node.Filename = &value
	}
	if value, ok := fc.mutation.Filetype(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileimport.FieldFiletype,
		})
		_node.Filetype = &value
	}
	if value, ok := fc.mutation.Importdate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fileimport.FieldImportdate,
		})
		_node.Importdate = &value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileimport.FieldStatus,
		})
		_node.Status = &value
	}
	return _node, _spec
}

// FileimportCreateBulk is the builder for creating many Fileimport entities in bulk.
type FileimportCreateBulk struct {
	config
	builders []*FileimportCreate
}

// Save creates the Fileimport entities in the database.
func (fcb *FileimportCreateBulk) Save(ctx context.Context) ([]*Fileimport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fileimport, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileimportMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileimportCreateBulk) SaveX(ctx context.Context) []*Fileimport {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}