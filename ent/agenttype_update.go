// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/agenttype"
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgenttypeUpdate is the builder for updating Agenttype entities.
type AgenttypeUpdate struct {
	config
	hooks    []Hook
	mutation *AgenttypeMutation
}

// Where adds a new predicate for the AgenttypeUpdate builder.
func (au *AgenttypeUpdate) Where(ps ...predicate.Agenttype) *AgenttypeUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetAgentid sets the "agentid" field.
func (au *AgenttypeUpdate) SetAgentid(s string) *AgenttypeUpdate {
	au.mutation.SetAgentid(s)
	return au
}

// SetNillableAgentid sets the "agentid" field if the given value is not nil.
func (au *AgenttypeUpdate) SetNillableAgentid(s *string) *AgenttypeUpdate {
	if s != nil {
		au.SetAgentid(*s)
	}
	return au
}

// ClearAgentid clears the value of the "agentid" field.
func (au *AgenttypeUpdate) ClearAgentid() *AgenttypeUpdate {
	au.mutation.ClearAgentid()
	return au
}

// SetAgentname sets the "agentname" field.
func (au *AgenttypeUpdate) SetAgentname(s string) *AgenttypeUpdate {
	au.mutation.SetAgentname(s)
	return au
}

// SetNillableAgentname sets the "agentname" field if the given value is not nil.
func (au *AgenttypeUpdate) SetNillableAgentname(s *string) *AgenttypeUpdate {
	if s != nil {
		au.SetAgentname(*s)
	}
	return au
}

// ClearAgentname clears the value of the "agentname" field.
func (au *AgenttypeUpdate) ClearAgentname() *AgenttypeUpdate {
	au.mutation.ClearAgentname()
	return au
}

// SetAgenttype sets the "agenttype" field.
func (au *AgenttypeUpdate) SetAgenttype(s string) *AgenttypeUpdate {
	au.mutation.SetAgenttype(s)
	return au
}

// SetNillableAgenttype sets the "agenttype" field if the given value is not nil.
func (au *AgenttypeUpdate) SetNillableAgenttype(s *string) *AgenttypeUpdate {
	if s != nil {
		au.SetAgenttype(*s)
	}
	return au
}

// ClearAgenttype clears the value of the "agenttype" field.
func (au *AgenttypeUpdate) ClearAgenttype() *AgenttypeUpdate {
	au.mutation.ClearAgenttype()
	return au
}

// Mutation returns the AgenttypeMutation object of the builder.
func (au *AgenttypeUpdate) Mutation() *AgenttypeMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AgenttypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		if err = au.check(); err != nil {
			return 0, err
		}
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgenttypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = au.check(); err != nil {
				return 0, err
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AgenttypeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AgenttypeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AgenttypeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AgenttypeUpdate) check() error {
	if v, ok := au.mutation.Agentid(); ok {
		if err := agenttype.AgentidValidator(v); err != nil {
			return &ValidationError{Name: "agentid", err: fmt.Errorf("ent: validator failed for field \"agentid\": %w", err)}
		}
	}
	if v, ok := au.mutation.Agentname(); ok {
		if err := agenttype.AgentnameValidator(v); err != nil {
			return &ValidationError{Name: "agentname", err: fmt.Errorf("ent: validator failed for field \"agentname\": %w", err)}
		}
	}
	if v, ok := au.mutation.Agenttype(); ok {
		if err := agenttype.AgenttypeValidator(v); err != nil {
			return &ValidationError{Name: "agenttype", err: fmt.Errorf("ent: validator failed for field \"agenttype\": %w", err)}
		}
	}
	return nil
}

func (au *AgenttypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agenttype.Table,
			Columns: agenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agenttype.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Agentid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agenttype.FieldAgentid,
		})
	}
	if au.mutation.AgentidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agenttype.FieldAgentid,
		})
	}
	if value, ok := au.mutation.Agentname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agenttype.FieldAgentname,
		})
	}
	if au.mutation.AgentnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agenttype.FieldAgentname,
		})
	}
	if value, ok := au.mutation.Agenttype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agenttype.FieldAgenttype,
		})
	}
	if au.mutation.AgenttypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agenttype.FieldAgenttype,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agenttype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AgenttypeUpdateOne is the builder for updating a single Agenttype entity.
type AgenttypeUpdateOne struct {
	config
	hooks    []Hook
	mutation *AgenttypeMutation
}

// SetAgentid sets the "agentid" field.
func (auo *AgenttypeUpdateOne) SetAgentid(s string) *AgenttypeUpdateOne {
	auo.mutation.SetAgentid(s)
	return auo
}

// SetNillableAgentid sets the "agentid" field if the given value is not nil.
func (auo *AgenttypeUpdateOne) SetNillableAgentid(s *string) *AgenttypeUpdateOne {
	if s != nil {
		auo.SetAgentid(*s)
	}
	return auo
}

// ClearAgentid clears the value of the "agentid" field.
func (auo *AgenttypeUpdateOne) ClearAgentid() *AgenttypeUpdateOne {
	auo.mutation.ClearAgentid()
	return auo
}

// SetAgentname sets the "agentname" field.
func (auo *AgenttypeUpdateOne) SetAgentname(s string) *AgenttypeUpdateOne {
	auo.mutation.SetAgentname(s)
	return auo
}

// SetNillableAgentname sets the "agentname" field if the given value is not nil.
func (auo *AgenttypeUpdateOne) SetNillableAgentname(s *string) *AgenttypeUpdateOne {
	if s != nil {
		auo.SetAgentname(*s)
	}
	return auo
}

// ClearAgentname clears the value of the "agentname" field.
func (auo *AgenttypeUpdateOne) ClearAgentname() *AgenttypeUpdateOne {
	auo.mutation.ClearAgentname()
	return auo
}

// SetAgenttype sets the "agenttype" field.
func (auo *AgenttypeUpdateOne) SetAgenttype(s string) *AgenttypeUpdateOne {
	auo.mutation.SetAgenttype(s)
	return auo
}

// SetNillableAgenttype sets the "agenttype" field if the given value is not nil.
func (auo *AgenttypeUpdateOne) SetNillableAgenttype(s *string) *AgenttypeUpdateOne {
	if s != nil {
		auo.SetAgenttype(*s)
	}
	return auo
}

// ClearAgenttype clears the value of the "agenttype" field.
func (auo *AgenttypeUpdateOne) ClearAgenttype() *AgenttypeUpdateOne {
	auo.mutation.ClearAgenttype()
	return auo
}

// Mutation returns the AgenttypeMutation object of the builder.
func (auo *AgenttypeUpdateOne) Mutation() *AgenttypeMutation {
	return auo.mutation
}

// Save executes the query and returns the updated Agenttype entity.
func (auo *AgenttypeUpdateOne) Save(ctx context.Context) (*Agenttype, error) {
	var (
		err  error
		node *Agenttype
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgenttypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auo.check(); err != nil {
				return nil, err
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AgenttypeUpdateOne) SaveX(ctx context.Context) *Agenttype {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AgenttypeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AgenttypeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AgenttypeUpdateOne) check() error {
	if v, ok := auo.mutation.Agentid(); ok {
		if err := agenttype.AgentidValidator(v); err != nil {
			return &ValidationError{Name: "agentid", err: fmt.Errorf("ent: validator failed for field \"agentid\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Agentname(); ok {
		if err := agenttype.AgentnameValidator(v); err != nil {
			return &ValidationError{Name: "agentname", err: fmt.Errorf("ent: validator failed for field \"agentname\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Agenttype(); ok {
		if err := agenttype.AgenttypeValidator(v); err != nil {
			return &ValidationError{Name: "agenttype", err: fmt.Errorf("ent: validator failed for field \"agenttype\": %w", err)}
		}
	}
	return nil
}

func (auo *AgenttypeUpdateOne) sqlSave(ctx context.Context) (_node *Agenttype, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agenttype.Table,
			Columns: agenttype.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agenttype.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Agenttype.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Agentid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agenttype.FieldAgentid,
		})
	}
	if auo.mutation.AgentidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agenttype.FieldAgentid,
		})
	}
	if value, ok := auo.mutation.Agentname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agenttype.FieldAgentname,
		})
	}
	if auo.mutation.AgentnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agenttype.FieldAgentname,
		})
	}
	if value, ok := auo.mutation.Agenttype(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agenttype.FieldAgenttype,
		})
	}
	if auo.mutation.AgenttypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agenttype.FieldAgenttype,
		})
	}
	_node = &Agenttype{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agenttype.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
