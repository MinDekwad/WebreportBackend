// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/writelog"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WritelogUpdate is the builder for updating Writelog entities.
type WritelogUpdate struct {
	config
	hooks    []Hook
	mutation *WritelogMutation
}

// Where adds a new predicate for the WritelogUpdate builder.
func (wu *WritelogUpdate) Where(ps ...predicate.Writelog) *WritelogUpdate {
	wu.mutation.predicates = append(wu.mutation.predicates, ps...)
	return wu
}

// SetAdminID sets the "admin_id" field.
func (wu *WritelogUpdate) SetAdminID(i int) *WritelogUpdate {
	wu.mutation.ResetAdminID()
	wu.mutation.SetAdminID(i)
	return wu
}

// AddAdminID adds i to the "admin_id" field.
func (wu *WritelogUpdate) AddAdminID(i int) *WritelogUpdate {
	wu.mutation.AddAdminID(i)
	return wu
}

// SetResource sets the "resource" field.
func (wu *WritelogUpdate) SetResource(s string) *WritelogUpdate {
	wu.mutation.SetResource(s)
	return wu
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (wu *WritelogUpdate) SetNillableResource(s *string) *WritelogUpdate {
	if s != nil {
		wu.SetResource(*s)
	}
	return wu
}

// ClearResource clears the value of the "resource" field.
func (wu *WritelogUpdate) ClearResource() *WritelogUpdate {
	wu.mutation.ClearResource()
	return wu
}

// SetActions sets the "actions" field.
func (wu *WritelogUpdate) SetActions(s string) *WritelogUpdate {
	wu.mutation.SetActions(s)
	return wu
}

// SetNillableActions sets the "actions" field if the given value is not nil.
func (wu *WritelogUpdate) SetNillableActions(s *string) *WritelogUpdate {
	if s != nil {
		wu.SetActions(*s)
	}
	return wu
}

// ClearActions clears the value of the "actions" field.
func (wu *WritelogUpdate) ClearActions() *WritelogUpdate {
	wu.mutation.ClearActions()
	return wu
}

// SetCreatedAt sets the "created_at" field.
func (wu *WritelogUpdate) SetCreatedAt(t time.Time) *WritelogUpdate {
	wu.mutation.SetCreatedAt(t)
	return wu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wu *WritelogUpdate) SetNillableCreatedAt(t *time.Time) *WritelogUpdate {
	if t != nil {
		wu.SetCreatedAt(*t)
	}
	return wu
}

// ClearCreatedAt clears the value of the "created_at" field.
func (wu *WritelogUpdate) ClearCreatedAt() *WritelogUpdate {
	wu.mutation.ClearCreatedAt()
	return wu
}

// Mutation returns the WritelogMutation object of the builder.
func (wu *WritelogUpdate) Mutation() *WritelogMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WritelogUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WritelogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WritelogUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WritelogUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WritelogUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WritelogUpdate) check() error {
	if v, ok := wu.mutation.Resource(); ok {
		if err := writelog.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf("ent: validator failed for field \"resource\": %w", err)}
		}
	}
	if v, ok := wu.mutation.Actions(); ok {
		if err := writelog.ActionsValidator(v); err != nil {
			return &ValidationError{Name: "actions", err: fmt.Errorf("ent: validator failed for field \"actions\": %w", err)}
		}
	}
	return nil
}

func (wu *WritelogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   writelog.Table,
			Columns: writelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: writelog.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.AdminID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: writelog.FieldAdminID,
		})
	}
	if value, ok := wu.mutation.AddedAdminID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: writelog.FieldAdminID,
		})
	}
	if value, ok := wu.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: writelog.FieldResource,
		})
	}
	if wu.mutation.ResourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: writelog.FieldResource,
		})
	}
	if value, ok := wu.mutation.Actions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: writelog.FieldActions,
		})
	}
	if wu.mutation.ActionsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: writelog.FieldActions,
		})
	}
	if value, ok := wu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: writelog.FieldCreatedAt,
		})
	}
	if wu.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: writelog.FieldCreatedAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{writelog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// WritelogUpdateOne is the builder for updating a single Writelog entity.
type WritelogUpdateOne struct {
	config
	hooks    []Hook
	mutation *WritelogMutation
}

// SetAdminID sets the "admin_id" field.
func (wuo *WritelogUpdateOne) SetAdminID(i int) *WritelogUpdateOne {
	wuo.mutation.ResetAdminID()
	wuo.mutation.SetAdminID(i)
	return wuo
}

// AddAdminID adds i to the "admin_id" field.
func (wuo *WritelogUpdateOne) AddAdminID(i int) *WritelogUpdateOne {
	wuo.mutation.AddAdminID(i)
	return wuo
}

// SetResource sets the "resource" field.
func (wuo *WritelogUpdateOne) SetResource(s string) *WritelogUpdateOne {
	wuo.mutation.SetResource(s)
	return wuo
}

// SetNillableResource sets the "resource" field if the given value is not nil.
func (wuo *WritelogUpdateOne) SetNillableResource(s *string) *WritelogUpdateOne {
	if s != nil {
		wuo.SetResource(*s)
	}
	return wuo
}

// ClearResource clears the value of the "resource" field.
func (wuo *WritelogUpdateOne) ClearResource() *WritelogUpdateOne {
	wuo.mutation.ClearResource()
	return wuo
}

// SetActions sets the "actions" field.
func (wuo *WritelogUpdateOne) SetActions(s string) *WritelogUpdateOne {
	wuo.mutation.SetActions(s)
	return wuo
}

// SetNillableActions sets the "actions" field if the given value is not nil.
func (wuo *WritelogUpdateOne) SetNillableActions(s *string) *WritelogUpdateOne {
	if s != nil {
		wuo.SetActions(*s)
	}
	return wuo
}

// ClearActions clears the value of the "actions" field.
func (wuo *WritelogUpdateOne) ClearActions() *WritelogUpdateOne {
	wuo.mutation.ClearActions()
	return wuo
}

// SetCreatedAt sets the "created_at" field.
func (wuo *WritelogUpdateOne) SetCreatedAt(t time.Time) *WritelogUpdateOne {
	wuo.mutation.SetCreatedAt(t)
	return wuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (wuo *WritelogUpdateOne) SetNillableCreatedAt(t *time.Time) *WritelogUpdateOne {
	if t != nil {
		wuo.SetCreatedAt(*t)
	}
	return wuo
}

// ClearCreatedAt clears the value of the "created_at" field.
func (wuo *WritelogUpdateOne) ClearCreatedAt() *WritelogUpdateOne {
	wuo.mutation.ClearCreatedAt()
	return wuo
}

// Mutation returns the WritelogMutation object of the builder.
func (wuo *WritelogUpdateOne) Mutation() *WritelogMutation {
	return wuo.mutation
}

// Save executes the query and returns the updated Writelog entity.
func (wuo *WritelogUpdateOne) Save(ctx context.Context) (*Writelog, error) {
	var (
		err  error
		node *Writelog
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WritelogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WritelogUpdateOne) SaveX(ctx context.Context) *Writelog {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WritelogUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WritelogUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WritelogUpdateOne) check() error {
	if v, ok := wuo.mutation.Resource(); ok {
		if err := writelog.ResourceValidator(v); err != nil {
			return &ValidationError{Name: "resource", err: fmt.Errorf("ent: validator failed for field \"resource\": %w", err)}
		}
	}
	if v, ok := wuo.mutation.Actions(); ok {
		if err := writelog.ActionsValidator(v); err != nil {
			return &ValidationError{Name: "actions", err: fmt.Errorf("ent: validator failed for field \"actions\": %w", err)}
		}
	}
	return nil
}

func (wuo *WritelogUpdateOne) sqlSave(ctx context.Context) (_node *Writelog, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   writelog.Table,
			Columns: writelog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: writelog.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Writelog.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.AdminID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: writelog.FieldAdminID,
		})
	}
	if value, ok := wuo.mutation.AddedAdminID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: writelog.FieldAdminID,
		})
	}
	if value, ok := wuo.mutation.Resource(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: writelog.FieldResource,
		})
	}
	if wuo.mutation.ResourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: writelog.FieldResource,
		})
	}
	if value, ok := wuo.mutation.Actions(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: writelog.FieldActions,
		})
	}
	if wuo.mutation.ActionsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: writelog.FieldActions,
		})
	}
	if value, ok := wuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: writelog.FieldCreatedAt,
		})
	}
	if wuo.mutation.CreatedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: writelog.FieldCreatedAt,
		})
	}
	_node = &Writelog{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{writelog.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
