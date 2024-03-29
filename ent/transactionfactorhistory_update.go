// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/transactionfactor"
	"go-api-report2/ent/transactionfactorhistory"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TransactionfactorhistoryUpdate is the builder for updating Transactionfactorhistory entities.
type TransactionfactorhistoryUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionfactorhistoryMutation
}

// Where adds a new predicate for the TransactionfactorhistoryUpdate builder.
func (tu *TransactionfactorhistoryUpdate) Where(ps ...predicate.Transactionfactorhistory) *TransactionfactorhistoryUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetWalletID sets the "WalletID" field.
func (tu *TransactionfactorhistoryUpdate) SetWalletID(s string) *TransactionfactorhistoryUpdate {
	tu.mutation.SetWalletID(s)
	return tu
}

// SetRankTransactionFactor sets the "RankTransactionFactor" field.
func (tu *TransactionfactorhistoryUpdate) SetRankTransactionFactor(i int) *TransactionfactorhistoryUpdate {
	tu.mutation.ResetRankTransactionFactor()
	tu.mutation.SetRankTransactionFactor(i)
	return tu
}

// SetNillableRankTransactionFactor sets the "RankTransactionFactor" field if the given value is not nil.
func (tu *TransactionfactorhistoryUpdate) SetNillableRankTransactionFactor(i *int) *TransactionfactorhistoryUpdate {
	if i != nil {
		tu.SetRankTransactionFactor(*i)
	}
	return tu
}

// AddRankTransactionFactor adds i to the "RankTransactionFactor" field.
func (tu *TransactionfactorhistoryUpdate) AddRankTransactionFactor(i int) *TransactionfactorhistoryUpdate {
	tu.mutation.AddRankTransactionFactor(i)
	return tu
}

// ClearRankTransactionFactor clears the value of the "RankTransactionFactor" field.
func (tu *TransactionfactorhistoryUpdate) ClearRankTransactionFactor() *TransactionfactorhistoryUpdate {
	tu.mutation.ClearRankTransactionFactor()
	return tu
}

// SetDateCalRank sets the "DateCalRank" field.
func (tu *TransactionfactorhistoryUpdate) SetDateCalRank(t time.Time) *TransactionfactorhistoryUpdate {
	tu.mutation.SetDateCalRank(t)
	return tu
}

// SetNillableDateCalRank sets the "DateCalRank" field if the given value is not nil.
func (tu *TransactionfactorhistoryUpdate) SetNillableDateCalRank(t *time.Time) *TransactionfactorhistoryUpdate {
	if t != nil {
		tu.SetDateCalRank(*t)
	}
	return tu
}

// ClearDateCalRank clears the value of the "DateCalRank" field.
func (tu *TransactionfactorhistoryUpdate) ClearDateCalRank() *TransactionfactorhistoryUpdate {
	tu.mutation.ClearDateCalRank()
	return tu
}

// SetTransactionfactorID sets the "Transactionfactor" edge to the Transactionfactor entity by ID.
func (tu *TransactionfactorhistoryUpdate) SetTransactionfactorID(id int) *TransactionfactorhistoryUpdate {
	tu.mutation.SetTransactionfactorID(id)
	return tu
}

// SetNillableTransactionfactorID sets the "Transactionfactor" edge to the Transactionfactor entity by ID if the given value is not nil.
func (tu *TransactionfactorhistoryUpdate) SetNillableTransactionfactorID(id *int) *TransactionfactorhistoryUpdate {
	if id != nil {
		tu = tu.SetTransactionfactorID(*id)
	}
	return tu
}

// SetTransactionfactor sets the "Transactionfactor" edge to the Transactionfactor entity.
func (tu *TransactionfactorhistoryUpdate) SetTransactionfactor(t *Transactionfactor) *TransactionfactorhistoryUpdate {
	return tu.SetTransactionfactorID(t.ID)
}

// Mutation returns the TransactionfactorhistoryMutation object of the builder.
func (tu *TransactionfactorhistoryUpdate) Mutation() *TransactionfactorhistoryMutation {
	return tu.mutation
}

// ClearTransactionfactor clears the "Transactionfactor" edge to the Transactionfactor entity.
func (tu *TransactionfactorhistoryUpdate) ClearTransactionfactor() *TransactionfactorhistoryUpdate {
	tu.mutation.ClearTransactionfactor()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionfactorhistoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		if err = tu.check(); err != nil {
			return 0, err
		}
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionfactorhistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tu.check(); err != nil {
				return 0, err
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TransactionfactorhistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionfactorhistoryUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionfactorhistoryUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionfactorhistoryUpdate) check() error {
	if v, ok := tu.mutation.WalletID(); ok {
		if err := transactionfactorhistory.WalletIDValidator(v); err != nil {
			return &ValidationError{Name: "WalletID", err: fmt.Errorf("ent: validator failed for field \"WalletID\": %w", err)}
		}
	}
	return nil
}

func (tu *TransactionfactorhistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactionfactorhistory.Table,
			Columns: transactionfactorhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactionfactorhistory.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.WalletID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactorhistory.FieldWalletID,
		})
	}
	if value, ok := tu.mutation.RankTransactionFactor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactorhistory.FieldRankTransactionFactor,
		})
	}
	if value, ok := tu.mutation.AddedRankTransactionFactor(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactorhistory.FieldRankTransactionFactor,
		})
	}
	if tu.mutation.RankTransactionFactorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: transactionfactorhistory.FieldRankTransactionFactor,
		})
	}
	if value, ok := tu.mutation.DateCalRank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactionfactorhistory.FieldDateCalRank,
		})
	}
	if tu.mutation.DateCalRankCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: transactionfactorhistory.FieldDateCalRank,
		})
	}
	if tu.mutation.TransactionfactorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionfactorhistory.TransactionfactorTable,
			Columns: []string{transactionfactorhistory.TransactionfactorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TransactionfactorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionfactorhistory.TransactionfactorTable,
			Columns: []string{transactionfactorhistory.TransactionfactorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionfactorhistory.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TransactionfactorhistoryUpdateOne is the builder for updating a single Transactionfactorhistory entity.
type TransactionfactorhistoryUpdateOne struct {
	config
	hooks    []Hook
	mutation *TransactionfactorhistoryMutation
}

// SetWalletID sets the "WalletID" field.
func (tuo *TransactionfactorhistoryUpdateOne) SetWalletID(s string) *TransactionfactorhistoryUpdateOne {
	tuo.mutation.SetWalletID(s)
	return tuo
}

// SetRankTransactionFactor sets the "RankTransactionFactor" field.
func (tuo *TransactionfactorhistoryUpdateOne) SetRankTransactionFactor(i int) *TransactionfactorhistoryUpdateOne {
	tuo.mutation.ResetRankTransactionFactor()
	tuo.mutation.SetRankTransactionFactor(i)
	return tuo
}

// SetNillableRankTransactionFactor sets the "RankTransactionFactor" field if the given value is not nil.
func (tuo *TransactionfactorhistoryUpdateOne) SetNillableRankTransactionFactor(i *int) *TransactionfactorhistoryUpdateOne {
	if i != nil {
		tuo.SetRankTransactionFactor(*i)
	}
	return tuo
}

// AddRankTransactionFactor adds i to the "RankTransactionFactor" field.
func (tuo *TransactionfactorhistoryUpdateOne) AddRankTransactionFactor(i int) *TransactionfactorhistoryUpdateOne {
	tuo.mutation.AddRankTransactionFactor(i)
	return tuo
}

// ClearRankTransactionFactor clears the value of the "RankTransactionFactor" field.
func (tuo *TransactionfactorhistoryUpdateOne) ClearRankTransactionFactor() *TransactionfactorhistoryUpdateOne {
	tuo.mutation.ClearRankTransactionFactor()
	return tuo
}

// SetDateCalRank sets the "DateCalRank" field.
func (tuo *TransactionfactorhistoryUpdateOne) SetDateCalRank(t time.Time) *TransactionfactorhistoryUpdateOne {
	tuo.mutation.SetDateCalRank(t)
	return tuo
}

// SetNillableDateCalRank sets the "DateCalRank" field if the given value is not nil.
func (tuo *TransactionfactorhistoryUpdateOne) SetNillableDateCalRank(t *time.Time) *TransactionfactorhistoryUpdateOne {
	if t != nil {
		tuo.SetDateCalRank(*t)
	}
	return tuo
}

// ClearDateCalRank clears the value of the "DateCalRank" field.
func (tuo *TransactionfactorhistoryUpdateOne) ClearDateCalRank() *TransactionfactorhistoryUpdateOne {
	tuo.mutation.ClearDateCalRank()
	return tuo
}

// SetTransactionfactorID sets the "Transactionfactor" edge to the Transactionfactor entity by ID.
func (tuo *TransactionfactorhistoryUpdateOne) SetTransactionfactorID(id int) *TransactionfactorhistoryUpdateOne {
	tuo.mutation.SetTransactionfactorID(id)
	return tuo
}

// SetNillableTransactionfactorID sets the "Transactionfactor" edge to the Transactionfactor entity by ID if the given value is not nil.
func (tuo *TransactionfactorhistoryUpdateOne) SetNillableTransactionfactorID(id *int) *TransactionfactorhistoryUpdateOne {
	if id != nil {
		tuo = tuo.SetTransactionfactorID(*id)
	}
	return tuo
}

// SetTransactionfactor sets the "Transactionfactor" edge to the Transactionfactor entity.
func (tuo *TransactionfactorhistoryUpdateOne) SetTransactionfactor(t *Transactionfactor) *TransactionfactorhistoryUpdateOne {
	return tuo.SetTransactionfactorID(t.ID)
}

// Mutation returns the TransactionfactorhistoryMutation object of the builder.
func (tuo *TransactionfactorhistoryUpdateOne) Mutation() *TransactionfactorhistoryMutation {
	return tuo.mutation
}

// ClearTransactionfactor clears the "Transactionfactor" edge to the Transactionfactor entity.
func (tuo *TransactionfactorhistoryUpdateOne) ClearTransactionfactor() *TransactionfactorhistoryUpdateOne {
	tuo.mutation.ClearTransactionfactor()
	return tuo
}

// Save executes the query and returns the updated Transactionfactorhistory entity.
func (tuo *TransactionfactorhistoryUpdateOne) Save(ctx context.Context) (*Transactionfactorhistory, error) {
	var (
		err  error
		node *Transactionfactorhistory
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionfactorhistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tuo.check(); err != nil {
				return nil, err
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TransactionfactorhistoryUpdateOne) SaveX(ctx context.Context) *Transactionfactorhistory {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionfactorhistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionfactorhistoryUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionfactorhistoryUpdateOne) check() error {
	if v, ok := tuo.mutation.WalletID(); ok {
		if err := transactionfactorhistory.WalletIDValidator(v); err != nil {
			return &ValidationError{Name: "WalletID", err: fmt.Errorf("ent: validator failed for field \"WalletID\": %w", err)}
		}
	}
	return nil
}

func (tuo *TransactionfactorhistoryUpdateOne) sqlSave(ctx context.Context) (_node *Transactionfactorhistory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactionfactorhistory.Table,
			Columns: transactionfactorhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactionfactorhistory.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Transactionfactorhistory.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.WalletID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactorhistory.FieldWalletID,
		})
	}
	if value, ok := tuo.mutation.RankTransactionFactor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactorhistory.FieldRankTransactionFactor,
		})
	}
	if value, ok := tuo.mutation.AddedRankTransactionFactor(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactorhistory.FieldRankTransactionFactor,
		})
	}
	if tuo.mutation.RankTransactionFactorCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: transactionfactorhistory.FieldRankTransactionFactor,
		})
	}
	if value, ok := tuo.mutation.DateCalRank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactionfactorhistory.FieldDateCalRank,
		})
	}
	if tuo.mutation.DateCalRankCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: transactionfactorhistory.FieldDateCalRank,
		})
	}
	if tuo.mutation.TransactionfactorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionfactorhistory.TransactionfactorTable,
			Columns: []string{transactionfactorhistory.TransactionfactorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TransactionfactorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transactionfactorhistory.TransactionfactorTable,
			Columns: []string{transactionfactorhistory.TransactionfactorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Transactionfactorhistory{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionfactorhistory.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
