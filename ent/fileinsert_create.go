// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/fileinsert"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FileinsertCreate is the builder for creating a Fileinsert entity.
type FileinsertCreate struct {
	config
	mutation *FileinsertMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (fc *FileinsertCreate) SetName(s string) *FileinsertCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fc *FileinsertCreate) SetNillableName(s *string) *FileinsertCreate {
	if s != nil {
		fc.SetName(*s)
	}
	return fc
}

// SetImportdate sets the "importdate" field.
func (fc *FileinsertCreate) SetImportdate(t time.Time) *FileinsertCreate {
	fc.mutation.SetImportdate(t)
	return fc
}

// SetNillableImportdate sets the "importdate" field if the given value is not nil.
func (fc *FileinsertCreate) SetNillableImportdate(t *time.Time) *FileinsertCreate {
	if t != nil {
		fc.SetImportdate(*t)
	}
	return fc
}

// SetIsSuccess sets the "IsSuccess" field.
func (fc *FileinsertCreate) SetIsSuccess(b bool) *FileinsertCreate {
	fc.mutation.SetIsSuccess(b)
	return fc
}

// SetNillableIsSuccess sets the "IsSuccess" field if the given value is not nil.
func (fc *FileinsertCreate) SetNillableIsSuccess(b *bool) *FileinsertCreate {
	if b != nil {
		fc.SetIsSuccess(*b)
	}
	return fc
}

// SetUpdateDate sets the "UpdateDate" field.
func (fc *FileinsertCreate) SetUpdateDate(t time.Time) *FileinsertCreate {
	fc.mutation.SetUpdateDate(t)
	return fc
}

// SetNillableUpdateDate sets the "UpdateDate" field if the given value is not nil.
func (fc *FileinsertCreate) SetNillableUpdateDate(t *time.Time) *FileinsertCreate {
	if t != nil {
		fc.SetUpdateDate(*t)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileinsertCreate) SetID(i int) *FileinsertCreate {
	fc.mutation.SetID(i)
	return fc
}

// Mutation returns the FileinsertMutation object of the builder.
func (fc *FileinsertCreate) Mutation() *FileinsertMutation {
	return fc.mutation
}

// Save creates the Fileinsert in the database.
func (fc *FileinsertCreate) Save(ctx context.Context) (*Fileinsert, error) {
	var (
		err  error
		node *Fileinsert
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileinsertMutation)
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
func (fc *FileinsertCreate) SaveX(ctx context.Context) *Fileinsert {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fc *FileinsertCreate) defaults() {
	if _, ok := fc.mutation.IsSuccess(); !ok {
		v := fileinsert.DefaultIsSuccess
		fc.mutation.SetIsSuccess(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileinsertCreate) check() error {
	if v, ok := fc.mutation.Name(); ok {
		if err := fileinsert.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	return nil
}

func (fc *FileinsertCreate) sqlSave(ctx context.Context) (*Fileinsert, error) {
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

func (fc *FileinsertCreate) createSpec() (*Fileinsert, *sqlgraph.CreateSpec) {
	var (
		_node = &Fileinsert{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: fileinsert.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: fileinsert.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: fileinsert.FieldName,
		})
		_node.Name = &value
	}
	if value, ok := fc.mutation.Importdate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fileinsert.FieldImportdate,
		})
		_node.Importdate = &value
	}
	if value, ok := fc.mutation.IsSuccess(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: fileinsert.FieldIsSuccess,
		})
		_node.IsSuccess = value
	}
	if value, ok := fc.mutation.UpdateDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: fileinsert.FieldUpdateDate,
		})
		_node.UpdateDate = &value
	}
	return _node, _spec
}

// FileinsertCreateBulk is the builder for creating many Fileinsert entities in bulk.
type FileinsertCreateBulk struct {
	config
	builders []*FileinsertCreate
}

// Save creates the Fileinsert entities in the database.
func (fcb *FileinsertCreateBulk) Save(ctx context.Context) ([]*Fileinsert, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fileinsert, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileinsertMutation)
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
func (fcb *FileinsertCreateBulk) SaveX(ctx context.Context) []*Fileinsert {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
