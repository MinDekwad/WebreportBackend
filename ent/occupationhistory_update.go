// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/occupationhistory"
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// OccupationhistoryUpdate is the builder for updating Occupationhistory entities.
type OccupationhistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OccupationhistoryMutation
}

// Where adds a new predicate for the OccupationhistoryUpdate builder.
func (ou *OccupationhistoryUpdate) Where(ps ...predicate.Occupationhistory) *OccupationhistoryUpdate {
	ou.mutation.predicates = append(ou.mutation.predicates, ps...)
	return ou
}

// SetWalletID sets the "WalletID" field.
func (ou *OccupationhistoryUpdate) SetWalletID(s string) *OccupationhistoryUpdate {
	ou.mutation.SetWalletID(s)
	return ou
}

// SetOccupationName sets the "OccupationName" field.
func (ou *OccupationhistoryUpdate) SetOccupationName(s string) *OccupationhistoryUpdate {
	ou.mutation.SetOccupationName(s)
	return ou
}

// SetNillableOccupationName sets the "OccupationName" field if the given value is not nil.
func (ou *OccupationhistoryUpdate) SetNillableOccupationName(s *string) *OccupationhistoryUpdate {
	if s != nil {
		ou.SetOccupationName(*s)
	}
	return ou
}

// ClearOccupationName clears the value of the "OccupationName" field.
func (ou *OccupationhistoryUpdate) ClearOccupationName() *OccupationhistoryUpdate {
	ou.mutation.ClearOccupationName()
	return ou
}

// SetRankOccupation sets the "RankOccupation" field.
func (ou *OccupationhistoryUpdate) SetRankOccupation(i int) *OccupationhistoryUpdate {
	ou.mutation.ResetRankOccupation()
	ou.mutation.SetRankOccupation(i)
	return ou
}

// SetNillableRankOccupation sets the "RankOccupation" field if the given value is not nil.
func (ou *OccupationhistoryUpdate) SetNillableRankOccupation(i *int) *OccupationhistoryUpdate {
	if i != nil {
		ou.SetRankOccupation(*i)
	}
	return ou
}

// AddRankOccupation adds i to the "RankOccupation" field.
func (ou *OccupationhistoryUpdate) AddRankOccupation(i int) *OccupationhistoryUpdate {
	ou.mutation.AddRankOccupation(i)
	return ou
}

// ClearRankOccupation clears the value of the "RankOccupation" field.
func (ou *OccupationhistoryUpdate) ClearRankOccupation() *OccupationhistoryUpdate {
	ou.mutation.ClearRankOccupation()
	return ou
}

// SetDateCalRank sets the "DateCalRank" field.
func (ou *OccupationhistoryUpdate) SetDateCalRank(t time.Time) *OccupationhistoryUpdate {
	ou.mutation.SetDateCalRank(t)
	return ou
}

// SetNillableDateCalRank sets the "DateCalRank" field if the given value is not nil.
func (ou *OccupationhistoryUpdate) SetNillableDateCalRank(t *time.Time) *OccupationhistoryUpdate {
	if t != nil {
		ou.SetDateCalRank(*t)
	}
	return ou
}

// ClearDateCalRank clears the value of the "DateCalRank" field.
func (ou *OccupationhistoryUpdate) ClearDateCalRank() *OccupationhistoryUpdate {
	ou.mutation.ClearDateCalRank()
	return ou
}

// Mutation returns the OccupationhistoryMutation object of the builder.
func (ou *OccupationhistoryUpdate) Mutation() *OccupationhistoryMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OccupationhistoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		if err = ou.check(); err != nil {
			return 0, err
		}
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OccupationhistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ou.check(); err != nil {
				return 0, err
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OccupationhistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OccupationhistoryUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OccupationhistoryUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OccupationhistoryUpdate) check() error {
	if v, ok := ou.mutation.WalletID(); ok {
		if err := occupationhistory.WalletIDValidator(v); err != nil {
			return &ValidationError{Name: "WalletID", err: fmt.Errorf("ent: validator failed for field \"WalletID\": %w", err)}
		}
	}
	if v, ok := ou.mutation.OccupationName(); ok {
		if err := occupationhistory.OccupationNameValidator(v); err != nil {
			return &ValidationError{Name: "OccupationName", err: fmt.Errorf("ent: validator failed for field \"OccupationName\": %w", err)}
		}
	}
	return nil
}

func (ou *OccupationhistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   occupationhistory.Table,
			Columns: occupationhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: occupationhistory.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.WalletID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occupationhistory.FieldWalletID,
		})
	}
	if value, ok := ou.mutation.OccupationName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occupationhistory.FieldOccupationName,
		})
	}
	if ou.mutation.OccupationNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: occupationhistory.FieldOccupationName,
		})
	}
	if value, ok := ou.mutation.RankOccupation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: occupationhistory.FieldRankOccupation,
		})
	}
	if value, ok := ou.mutation.AddedRankOccupation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: occupationhistory.FieldRankOccupation,
		})
	}
	if ou.mutation.RankOccupationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: occupationhistory.FieldRankOccupation,
		})
	}
	if value, ok := ou.mutation.DateCalRank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: occupationhistory.FieldDateCalRank,
		})
	}
	if ou.mutation.DateCalRankCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: occupationhistory.FieldDateCalRank,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{occupationhistory.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// OccupationhistoryUpdateOne is the builder for updating a single Occupationhistory entity.
type OccupationhistoryUpdateOne struct {
	config
	hooks    []Hook
	mutation *OccupationhistoryMutation
}

// SetWalletID sets the "WalletID" field.
func (ouo *OccupationhistoryUpdateOne) SetWalletID(s string) *OccupationhistoryUpdateOne {
	ouo.mutation.SetWalletID(s)
	return ouo
}

// SetOccupationName sets the "OccupationName" field.
func (ouo *OccupationhistoryUpdateOne) SetOccupationName(s string) *OccupationhistoryUpdateOne {
	ouo.mutation.SetOccupationName(s)
	return ouo
}

// SetNillableOccupationName sets the "OccupationName" field if the given value is not nil.
func (ouo *OccupationhistoryUpdateOne) SetNillableOccupationName(s *string) *OccupationhistoryUpdateOne {
	if s != nil {
		ouo.SetOccupationName(*s)
	}
	return ouo
}

// ClearOccupationName clears the value of the "OccupationName" field.
func (ouo *OccupationhistoryUpdateOne) ClearOccupationName() *OccupationhistoryUpdateOne {
	ouo.mutation.ClearOccupationName()
	return ouo
}

// SetRankOccupation sets the "RankOccupation" field.
func (ouo *OccupationhistoryUpdateOne) SetRankOccupation(i int) *OccupationhistoryUpdateOne {
	ouo.mutation.ResetRankOccupation()
	ouo.mutation.SetRankOccupation(i)
	return ouo
}

// SetNillableRankOccupation sets the "RankOccupation" field if the given value is not nil.
func (ouo *OccupationhistoryUpdateOne) SetNillableRankOccupation(i *int) *OccupationhistoryUpdateOne {
	if i != nil {
		ouo.SetRankOccupation(*i)
	}
	return ouo
}

// AddRankOccupation adds i to the "RankOccupation" field.
func (ouo *OccupationhistoryUpdateOne) AddRankOccupation(i int) *OccupationhistoryUpdateOne {
	ouo.mutation.AddRankOccupation(i)
	return ouo
}

// ClearRankOccupation clears the value of the "RankOccupation" field.
func (ouo *OccupationhistoryUpdateOne) ClearRankOccupation() *OccupationhistoryUpdateOne {
	ouo.mutation.ClearRankOccupation()
	return ouo
}

// SetDateCalRank sets the "DateCalRank" field.
func (ouo *OccupationhistoryUpdateOne) SetDateCalRank(t time.Time) *OccupationhistoryUpdateOne {
	ouo.mutation.SetDateCalRank(t)
	return ouo
}

// SetNillableDateCalRank sets the "DateCalRank" field if the given value is not nil.
func (ouo *OccupationhistoryUpdateOne) SetNillableDateCalRank(t *time.Time) *OccupationhistoryUpdateOne {
	if t != nil {
		ouo.SetDateCalRank(*t)
	}
	return ouo
}

// ClearDateCalRank clears the value of the "DateCalRank" field.
func (ouo *OccupationhistoryUpdateOne) ClearDateCalRank() *OccupationhistoryUpdateOne {
	ouo.mutation.ClearDateCalRank()
	return ouo
}

// Mutation returns the OccupationhistoryMutation object of the builder.
func (ouo *OccupationhistoryUpdateOne) Mutation() *OccupationhistoryMutation {
	return ouo.mutation
}

// Save executes the query and returns the updated Occupationhistory entity.
func (ouo *OccupationhistoryUpdateOne) Save(ctx context.Context) (*Occupationhistory, error) {
	var (
		err  error
		node *Occupationhistory
	)
	if len(ouo.hooks) == 0 {
		if err = ouo.check(); err != nil {
			return nil, err
		}
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OccupationhistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ouo.check(); err != nil {
				return nil, err
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			mut = ouo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ouo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OccupationhistoryUpdateOne) SaveX(ctx context.Context) *Occupationhistory {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OccupationhistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OccupationhistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OccupationhistoryUpdateOne) check() error {
	if v, ok := ouo.mutation.WalletID(); ok {
		if err := occupationhistory.WalletIDValidator(v); err != nil {
			return &ValidationError{Name: "WalletID", err: fmt.Errorf("ent: validator failed for field \"WalletID\": %w", err)}
		}
	}
	if v, ok := ouo.mutation.OccupationName(); ok {
		if err := occupationhistory.OccupationNameValidator(v); err != nil {
			return &ValidationError{Name: "OccupationName", err: fmt.Errorf("ent: validator failed for field \"OccupationName\": %w", err)}
		}
	}
	return nil
}

func (ouo *OccupationhistoryUpdateOne) sqlSave(ctx context.Context) (_node *Occupationhistory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   occupationhistory.Table,
			Columns: occupationhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: occupationhistory.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Occupationhistory.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.WalletID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occupationhistory.FieldWalletID,
		})
	}
	if value, ok := ouo.mutation.OccupationName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: occupationhistory.FieldOccupationName,
		})
	}
	if ouo.mutation.OccupationNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: occupationhistory.FieldOccupationName,
		})
	}
	if value, ok := ouo.mutation.RankOccupation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: occupationhistory.FieldRankOccupation,
		})
	}
	if value, ok := ouo.mutation.AddedRankOccupation(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: occupationhistory.FieldRankOccupation,
		})
	}
	if ouo.mutation.RankOccupationCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: occupationhistory.FieldRankOccupation,
		})
	}
	if value, ok := ouo.mutation.DateCalRank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: occupationhistory.FieldDateCalRank,
		})
	}
	if ouo.mutation.DateCalRankCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: occupationhistory.FieldDateCalRank,
		})
	}
	_node = &Occupationhistory{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{occupationhistory.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
