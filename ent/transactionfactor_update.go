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

// TransactionfactorUpdate is the builder for updating Transactionfactor entities.
type TransactionfactorUpdate struct {
	config
	hooks    []Hook
	mutation *TransactionfactorMutation
}

// Where adds a new predicate for the TransactionfactorUpdate builder.
func (tu *TransactionfactorUpdate) Where(ps ...predicate.Transactionfactor) *TransactionfactorUpdate {
	tu.mutation.predicates = append(tu.mutation.predicates, ps...)
	return tu
}

// SetTransactionFactorName sets the "TransactionFactorName" field.
func (tu *TransactionfactorUpdate) SetTransactionFactorName(s string) *TransactionfactorUpdate {
	tu.mutation.SetTransactionFactorName(s)
	return tu
}

// SetNillableTransactionFactorName sets the "TransactionFactorName" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillableTransactionFactorName(s *string) *TransactionfactorUpdate {
	if s != nil {
		tu.SetTransactionFactorName(*s)
	}
	return tu
}

// ClearTransactionFactorName clears the value of the "TransactionFactorName" field.
func (tu *TransactionfactorUpdate) ClearTransactionFactorName() *TransactionfactorUpdate {
	tu.mutation.ClearTransactionFactorName()
	return tu
}

// SetTransactionType sets the "TransactionType" field.
func (tu *TransactionfactorUpdate) SetTransactionType(s string) *TransactionfactorUpdate {
	tu.mutation.SetTransactionType(s)
	return tu
}

// SetNillableTransactionType sets the "TransactionType" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillableTransactionType(s *string) *TransactionfactorUpdate {
	if s != nil {
		tu.SetTransactionType(*s)
	}
	return tu
}

// ClearTransactionType clears the value of the "TransactionType" field.
func (tu *TransactionfactorUpdate) ClearTransactionType() *TransactionfactorUpdate {
	tu.mutation.ClearTransactionType()
	return tu
}

// SetPaymentChannel sets the "PaymentChannel" field.
func (tu *TransactionfactorUpdate) SetPaymentChannel(s string) *TransactionfactorUpdate {
	tu.mutation.SetPaymentChannel(s)
	return tu
}

// SetNillablePaymentChannel sets the "PaymentChannel" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillablePaymentChannel(s *string) *TransactionfactorUpdate {
	if s != nil {
		tu.SetPaymentChannel(*s)
	}
	return tu
}

// ClearPaymentChannel clears the value of the "PaymentChannel" field.
func (tu *TransactionfactorUpdate) ClearPaymentChannel() *TransactionfactorUpdate {
	tu.mutation.ClearPaymentChannel()
	return tu
}

// SetPaymentType sets the "PaymentType" field.
func (tu *TransactionfactorUpdate) SetPaymentType(s string) *TransactionfactorUpdate {
	tu.mutation.SetPaymentType(s)
	return tu
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillablePaymentType(s *string) *TransactionfactorUpdate {
	if s != nil {
		tu.SetPaymentType(*s)
	}
	return tu
}

// ClearPaymentType clears the value of the "PaymentType" field.
func (tu *TransactionfactorUpdate) ClearPaymentType() *TransactionfactorUpdate {
	tu.mutation.ClearPaymentType()
	return tu
}

// SetNumDay sets the "NumDay" field.
func (tu *TransactionfactorUpdate) SetNumDay(i int) *TransactionfactorUpdate {
	tu.mutation.ResetNumDay()
	tu.mutation.SetNumDay(i)
	return tu
}

// SetNillableNumDay sets the "NumDay" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillableNumDay(i *int) *TransactionfactorUpdate {
	if i != nil {
		tu.SetNumDay(*i)
	}
	return tu
}

// AddNumDay adds i to the "NumDay" field.
func (tu *TransactionfactorUpdate) AddNumDay(i int) *TransactionfactorUpdate {
	tu.mutation.AddNumDay(i)
	return tu
}

// ClearNumDay clears the value of the "NumDay" field.
func (tu *TransactionfactorUpdate) ClearNumDay() *TransactionfactorUpdate {
	tu.mutation.ClearNumDay()
	return tu
}

// SetDate sets the "Date" field.
func (tu *TransactionfactorUpdate) SetDate(s string) *TransactionfactorUpdate {
	tu.mutation.SetDate(s)
	return tu
}

// SetNillableDate sets the "Date" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillableDate(s *string) *TransactionfactorUpdate {
	if s != nil {
		tu.SetDate(*s)
	}
	return tu
}

// ClearDate clears the value of the "Date" field.
func (tu *TransactionfactorUpdate) ClearDate() *TransactionfactorUpdate {
	tu.mutation.ClearDate()
	return tu
}

// SetUpdateDate sets the "UpdateDate" field.
func (tu *TransactionfactorUpdate) SetUpdateDate(t time.Time) *TransactionfactorUpdate {
	tu.mutation.SetUpdateDate(t)
	return tu
}

// SetNillableUpdateDate sets the "UpdateDate" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillableUpdateDate(t *time.Time) *TransactionfactorUpdate {
	if t != nil {
		tu.SetUpdateDate(*t)
	}
	return tu
}

// ClearUpdateDate clears the value of the "UpdateDate" field.
func (tu *TransactionfactorUpdate) ClearUpdateDate() *TransactionfactorUpdate {
	tu.mutation.ClearUpdateDate()
	return tu
}

// SetStatusApprove sets the "StatusApprove" field.
func (tu *TransactionfactorUpdate) SetStatusApprove(s string) *TransactionfactorUpdate {
	tu.mutation.SetStatusApprove(s)
	return tu
}

// SetNillableStatusApprove sets the "StatusApprove" field if the given value is not nil.
func (tu *TransactionfactorUpdate) SetNillableStatusApprove(s *string) *TransactionfactorUpdate {
	if s != nil {
		tu.SetStatusApprove(*s)
	}
	return tu
}

// ClearStatusApprove clears the value of the "StatusApprove" field.
func (tu *TransactionfactorUpdate) ClearStatusApprove() *TransactionfactorUpdate {
	tu.mutation.ClearStatusApprove()
	return tu
}

// AddTransactionhistoryIDs adds the "transactionhistory" edge to the Transactionfactorhistory entity by IDs.
func (tu *TransactionfactorUpdate) AddTransactionhistoryIDs(ids ...int) *TransactionfactorUpdate {
	tu.mutation.AddTransactionhistoryIDs(ids...)
	return tu
}

// AddTransactionhistory adds the "transactionhistory" edges to the Transactionfactorhistory entity.
func (tu *TransactionfactorUpdate) AddTransactionhistory(t ...*Transactionfactorhistory) *TransactionfactorUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTransactionhistoryIDs(ids...)
}

// Mutation returns the TransactionfactorMutation object of the builder.
func (tu *TransactionfactorUpdate) Mutation() *TransactionfactorMutation {
	return tu.mutation
}

// ClearTransactionhistory clears all "transactionhistory" edges to the Transactionfactorhistory entity.
func (tu *TransactionfactorUpdate) ClearTransactionhistory() *TransactionfactorUpdate {
	tu.mutation.ClearTransactionhistory()
	return tu
}

// RemoveTransactionhistoryIDs removes the "transactionhistory" edge to Transactionfactorhistory entities by IDs.
func (tu *TransactionfactorUpdate) RemoveTransactionhistoryIDs(ids ...int) *TransactionfactorUpdate {
	tu.mutation.RemoveTransactionhistoryIDs(ids...)
	return tu
}

// RemoveTransactionhistory removes "transactionhistory" edges to Transactionfactorhistory entities.
func (tu *TransactionfactorUpdate) RemoveTransactionhistory(t ...*Transactionfactorhistory) *TransactionfactorUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTransactionhistoryIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TransactionfactorUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*TransactionfactorMutation)
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
func (tu *TransactionfactorUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TransactionfactorUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TransactionfactorUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TransactionfactorUpdate) check() error {
	if v, ok := tu.mutation.TransactionFactorName(); ok {
		if err := transactionfactor.TransactionFactorNameValidator(v); err != nil {
			return &ValidationError{Name: "TransactionFactorName", err: fmt.Errorf("ent: validator failed for field \"TransactionFactorName\": %w", err)}
		}
	}
	if v, ok := tu.mutation.TransactionType(); ok {
		if err := transactionfactor.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "TransactionType", err: fmt.Errorf("ent: validator failed for field \"TransactionType\": %w", err)}
		}
	}
	if v, ok := tu.mutation.PaymentChannel(); ok {
		if err := transactionfactor.PaymentChannelValidator(v); err != nil {
			return &ValidationError{Name: "PaymentChannel", err: fmt.Errorf("ent: validator failed for field \"PaymentChannel\": %w", err)}
		}
	}
	if v, ok := tu.mutation.PaymentType(); ok {
		if err := transactionfactor.PaymentTypeValidator(v); err != nil {
			return &ValidationError{Name: "PaymentType", err: fmt.Errorf("ent: validator failed for field \"PaymentType\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Date(); ok {
		if err := transactionfactor.DateValidator(v); err != nil {
			return &ValidationError{Name: "Date", err: fmt.Errorf("ent: validator failed for field \"Date\": %w", err)}
		}
	}
	if v, ok := tu.mutation.StatusApprove(); ok {
		if err := transactionfactor.StatusApproveValidator(v); err != nil {
			return &ValidationError{Name: "StatusApprove", err: fmt.Errorf("ent: validator failed for field \"StatusApprove\": %w", err)}
		}
	}
	return nil
}

func (tu *TransactionfactorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactionfactor.Table,
			Columns: transactionfactor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactionfactor.FieldID,
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
	if value, ok := tu.mutation.TransactionFactorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldTransactionFactorName,
		})
	}
	if tu.mutation.TransactionFactorNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldTransactionFactorName,
		})
	}
	if value, ok := tu.mutation.TransactionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldTransactionType,
		})
	}
	if tu.mutation.TransactionTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldTransactionType,
		})
	}
	if value, ok := tu.mutation.PaymentChannel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldPaymentChannel,
		})
	}
	if tu.mutation.PaymentChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldPaymentChannel,
		})
	}
	if value, ok := tu.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldPaymentType,
		})
	}
	if tu.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldPaymentType,
		})
	}
	if value, ok := tu.mutation.NumDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactor.FieldNumDay,
		})
	}
	if value, ok := tu.mutation.AddedNumDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactor.FieldNumDay,
		})
	}
	if tu.mutation.NumDayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: transactionfactor.FieldNumDay,
		})
	}
	if value, ok := tu.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldDate,
		})
	}
	if tu.mutation.DateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldDate,
		})
	}
	if value, ok := tu.mutation.UpdateDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactionfactor.FieldUpdateDate,
		})
	}
	if tu.mutation.UpdateDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: transactionfactor.FieldUpdateDate,
		})
	}
	if value, ok := tu.mutation.StatusApprove(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldStatusApprove,
		})
	}
	if tu.mutation.StatusApproveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldStatusApprove,
		})
	}
	if tu.mutation.TransactionhistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactionfactor.TransactionhistoryTable,
			Columns: []string{transactionfactor.TransactionhistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactorhistory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTransactionhistoryIDs(); len(nodes) > 0 && !tu.mutation.TransactionhistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactionfactor.TransactionhistoryTable,
			Columns: []string{transactionfactor.TransactionhistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactorhistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TransactionhistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactionfactor.TransactionhistoryTable,
			Columns: []string{transactionfactor.TransactionhistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactorhistory.FieldID,
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
			err = &NotFoundError{transactionfactor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TransactionfactorUpdateOne is the builder for updating a single Transactionfactor entity.
type TransactionfactorUpdateOne struct {
	config
	hooks    []Hook
	mutation *TransactionfactorMutation
}

// SetTransactionFactorName sets the "TransactionFactorName" field.
func (tuo *TransactionfactorUpdateOne) SetTransactionFactorName(s string) *TransactionfactorUpdateOne {
	tuo.mutation.SetTransactionFactorName(s)
	return tuo
}

// SetNillableTransactionFactorName sets the "TransactionFactorName" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillableTransactionFactorName(s *string) *TransactionfactorUpdateOne {
	if s != nil {
		tuo.SetTransactionFactorName(*s)
	}
	return tuo
}

// ClearTransactionFactorName clears the value of the "TransactionFactorName" field.
func (tuo *TransactionfactorUpdateOne) ClearTransactionFactorName() *TransactionfactorUpdateOne {
	tuo.mutation.ClearTransactionFactorName()
	return tuo
}

// SetTransactionType sets the "TransactionType" field.
func (tuo *TransactionfactorUpdateOne) SetTransactionType(s string) *TransactionfactorUpdateOne {
	tuo.mutation.SetTransactionType(s)
	return tuo
}

// SetNillableTransactionType sets the "TransactionType" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillableTransactionType(s *string) *TransactionfactorUpdateOne {
	if s != nil {
		tuo.SetTransactionType(*s)
	}
	return tuo
}

// ClearTransactionType clears the value of the "TransactionType" field.
func (tuo *TransactionfactorUpdateOne) ClearTransactionType() *TransactionfactorUpdateOne {
	tuo.mutation.ClearTransactionType()
	return tuo
}

// SetPaymentChannel sets the "PaymentChannel" field.
func (tuo *TransactionfactorUpdateOne) SetPaymentChannel(s string) *TransactionfactorUpdateOne {
	tuo.mutation.SetPaymentChannel(s)
	return tuo
}

// SetNillablePaymentChannel sets the "PaymentChannel" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillablePaymentChannel(s *string) *TransactionfactorUpdateOne {
	if s != nil {
		tuo.SetPaymentChannel(*s)
	}
	return tuo
}

// ClearPaymentChannel clears the value of the "PaymentChannel" field.
func (tuo *TransactionfactorUpdateOne) ClearPaymentChannel() *TransactionfactorUpdateOne {
	tuo.mutation.ClearPaymentChannel()
	return tuo
}

// SetPaymentType sets the "PaymentType" field.
func (tuo *TransactionfactorUpdateOne) SetPaymentType(s string) *TransactionfactorUpdateOne {
	tuo.mutation.SetPaymentType(s)
	return tuo
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillablePaymentType(s *string) *TransactionfactorUpdateOne {
	if s != nil {
		tuo.SetPaymentType(*s)
	}
	return tuo
}

// ClearPaymentType clears the value of the "PaymentType" field.
func (tuo *TransactionfactorUpdateOne) ClearPaymentType() *TransactionfactorUpdateOne {
	tuo.mutation.ClearPaymentType()
	return tuo
}

// SetNumDay sets the "NumDay" field.
func (tuo *TransactionfactorUpdateOne) SetNumDay(i int) *TransactionfactorUpdateOne {
	tuo.mutation.ResetNumDay()
	tuo.mutation.SetNumDay(i)
	return tuo
}

// SetNillableNumDay sets the "NumDay" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillableNumDay(i *int) *TransactionfactorUpdateOne {
	if i != nil {
		tuo.SetNumDay(*i)
	}
	return tuo
}

// AddNumDay adds i to the "NumDay" field.
func (tuo *TransactionfactorUpdateOne) AddNumDay(i int) *TransactionfactorUpdateOne {
	tuo.mutation.AddNumDay(i)
	return tuo
}

// ClearNumDay clears the value of the "NumDay" field.
func (tuo *TransactionfactorUpdateOne) ClearNumDay() *TransactionfactorUpdateOne {
	tuo.mutation.ClearNumDay()
	return tuo
}

// SetDate sets the "Date" field.
func (tuo *TransactionfactorUpdateOne) SetDate(s string) *TransactionfactorUpdateOne {
	tuo.mutation.SetDate(s)
	return tuo
}

// SetNillableDate sets the "Date" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillableDate(s *string) *TransactionfactorUpdateOne {
	if s != nil {
		tuo.SetDate(*s)
	}
	return tuo
}

// ClearDate clears the value of the "Date" field.
func (tuo *TransactionfactorUpdateOne) ClearDate() *TransactionfactorUpdateOne {
	tuo.mutation.ClearDate()
	return tuo
}

// SetUpdateDate sets the "UpdateDate" field.
func (tuo *TransactionfactorUpdateOne) SetUpdateDate(t time.Time) *TransactionfactorUpdateOne {
	tuo.mutation.SetUpdateDate(t)
	return tuo
}

// SetNillableUpdateDate sets the "UpdateDate" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillableUpdateDate(t *time.Time) *TransactionfactorUpdateOne {
	if t != nil {
		tuo.SetUpdateDate(*t)
	}
	return tuo
}

// ClearUpdateDate clears the value of the "UpdateDate" field.
func (tuo *TransactionfactorUpdateOne) ClearUpdateDate() *TransactionfactorUpdateOne {
	tuo.mutation.ClearUpdateDate()
	return tuo
}

// SetStatusApprove sets the "StatusApprove" field.
func (tuo *TransactionfactorUpdateOne) SetStatusApprove(s string) *TransactionfactorUpdateOne {
	tuo.mutation.SetStatusApprove(s)
	return tuo
}

// SetNillableStatusApprove sets the "StatusApprove" field if the given value is not nil.
func (tuo *TransactionfactorUpdateOne) SetNillableStatusApprove(s *string) *TransactionfactorUpdateOne {
	if s != nil {
		tuo.SetStatusApprove(*s)
	}
	return tuo
}

// ClearStatusApprove clears the value of the "StatusApprove" field.
func (tuo *TransactionfactorUpdateOne) ClearStatusApprove() *TransactionfactorUpdateOne {
	tuo.mutation.ClearStatusApprove()
	return tuo
}

// AddTransactionhistoryIDs adds the "transactionhistory" edge to the Transactionfactorhistory entity by IDs.
func (tuo *TransactionfactorUpdateOne) AddTransactionhistoryIDs(ids ...int) *TransactionfactorUpdateOne {
	tuo.mutation.AddTransactionhistoryIDs(ids...)
	return tuo
}

// AddTransactionhistory adds the "transactionhistory" edges to the Transactionfactorhistory entity.
func (tuo *TransactionfactorUpdateOne) AddTransactionhistory(t ...*Transactionfactorhistory) *TransactionfactorUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTransactionhistoryIDs(ids...)
}

// Mutation returns the TransactionfactorMutation object of the builder.
func (tuo *TransactionfactorUpdateOne) Mutation() *TransactionfactorMutation {
	return tuo.mutation
}

// ClearTransactionhistory clears all "transactionhistory" edges to the Transactionfactorhistory entity.
func (tuo *TransactionfactorUpdateOne) ClearTransactionhistory() *TransactionfactorUpdateOne {
	tuo.mutation.ClearTransactionhistory()
	return tuo
}

// RemoveTransactionhistoryIDs removes the "transactionhistory" edge to Transactionfactorhistory entities by IDs.
func (tuo *TransactionfactorUpdateOne) RemoveTransactionhistoryIDs(ids ...int) *TransactionfactorUpdateOne {
	tuo.mutation.RemoveTransactionhistoryIDs(ids...)
	return tuo
}

// RemoveTransactionhistory removes "transactionhistory" edges to Transactionfactorhistory entities.
func (tuo *TransactionfactorUpdateOne) RemoveTransactionhistory(t ...*Transactionfactorhistory) *TransactionfactorUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTransactionhistoryIDs(ids...)
}

// Save executes the query and returns the updated Transactionfactor entity.
func (tuo *TransactionfactorUpdateOne) Save(ctx context.Context) (*Transactionfactor, error) {
	var (
		err  error
		node *Transactionfactor
	)
	if len(tuo.hooks) == 0 {
		if err = tuo.check(); err != nil {
			return nil, err
		}
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionfactorMutation)
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
func (tuo *TransactionfactorUpdateOne) SaveX(ctx context.Context) *Transactionfactor {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TransactionfactorUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TransactionfactorUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TransactionfactorUpdateOne) check() error {
	if v, ok := tuo.mutation.TransactionFactorName(); ok {
		if err := transactionfactor.TransactionFactorNameValidator(v); err != nil {
			return &ValidationError{Name: "TransactionFactorName", err: fmt.Errorf("ent: validator failed for field \"TransactionFactorName\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.TransactionType(); ok {
		if err := transactionfactor.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "TransactionType", err: fmt.Errorf("ent: validator failed for field \"TransactionType\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.PaymentChannel(); ok {
		if err := transactionfactor.PaymentChannelValidator(v); err != nil {
			return &ValidationError{Name: "PaymentChannel", err: fmt.Errorf("ent: validator failed for field \"PaymentChannel\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.PaymentType(); ok {
		if err := transactionfactor.PaymentTypeValidator(v); err != nil {
			return &ValidationError{Name: "PaymentType", err: fmt.Errorf("ent: validator failed for field \"PaymentType\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Date(); ok {
		if err := transactionfactor.DateValidator(v); err != nil {
			return &ValidationError{Name: "Date", err: fmt.Errorf("ent: validator failed for field \"Date\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.StatusApprove(); ok {
		if err := transactionfactor.StatusApproveValidator(v); err != nil {
			return &ValidationError{Name: "StatusApprove", err: fmt.Errorf("ent: validator failed for field \"StatusApprove\": %w", err)}
		}
	}
	return nil
}

func (tuo *TransactionfactorUpdateOne) sqlSave(ctx context.Context) (_node *Transactionfactor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   transactionfactor.Table,
			Columns: transactionfactor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: transactionfactor.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Transactionfactor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.TransactionFactorName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldTransactionFactorName,
		})
	}
	if tuo.mutation.TransactionFactorNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldTransactionFactorName,
		})
	}
	if value, ok := tuo.mutation.TransactionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldTransactionType,
		})
	}
	if tuo.mutation.TransactionTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldTransactionType,
		})
	}
	if value, ok := tuo.mutation.PaymentChannel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldPaymentChannel,
		})
	}
	if tuo.mutation.PaymentChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldPaymentChannel,
		})
	}
	if value, ok := tuo.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldPaymentType,
		})
	}
	if tuo.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldPaymentType,
		})
	}
	if value, ok := tuo.mutation.NumDay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactor.FieldNumDay,
		})
	}
	if value, ok := tuo.mutation.AddedNumDay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: transactionfactor.FieldNumDay,
		})
	}
	if tuo.mutation.NumDayCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: transactionfactor.FieldNumDay,
		})
	}
	if value, ok := tuo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldDate,
		})
	}
	if tuo.mutation.DateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldDate,
		})
	}
	if value, ok := tuo.mutation.UpdateDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: transactionfactor.FieldUpdateDate,
		})
	}
	if tuo.mutation.UpdateDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: transactionfactor.FieldUpdateDate,
		})
	}
	if value, ok := tuo.mutation.StatusApprove(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transactionfactor.FieldStatusApprove,
		})
	}
	if tuo.mutation.StatusApproveCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: transactionfactor.FieldStatusApprove,
		})
	}
	if tuo.mutation.TransactionhistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactionfactor.TransactionhistoryTable,
			Columns: []string{transactionfactor.TransactionhistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactorhistory.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTransactionhistoryIDs(); len(nodes) > 0 && !tuo.mutation.TransactionhistoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactionfactor.TransactionhistoryTable,
			Columns: []string{transactionfactor.TransactionhistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactorhistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TransactionhistoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transactionfactor.TransactionhistoryTable,
			Columns: []string{transactionfactor.TransactionhistoryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: transactionfactorhistory.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Transactionfactor{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{transactionfactor.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}