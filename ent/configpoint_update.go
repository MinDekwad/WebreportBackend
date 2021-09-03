// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/configpoint"
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConfigpointUpdate is the builder for updating Configpoint entities.
type ConfigpointUpdate struct {
	config
	hooks    []Hook
	mutation *ConfigpointMutation
}

// Where adds a new predicate for the ConfigpointUpdate builder.
func (cu *ConfigpointUpdate) Where(ps ...predicate.Configpoint) *ConfigpointUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetTransactionName sets the "TransactionName" field.
func (cu *ConfigpointUpdate) SetTransactionName(s string) *ConfigpointUpdate {
	cu.mutation.SetTransactionName(s)
	return cu
}

// SetNillableTransactionName sets the "TransactionName" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableTransactionName(s *string) *ConfigpointUpdate {
	if s != nil {
		cu.SetTransactionName(*s)
	}
	return cu
}

// ClearTransactionName clears the value of the "TransactionName" field.
func (cu *ConfigpointUpdate) ClearTransactionName() *ConfigpointUpdate {
	cu.mutation.ClearTransactionName()
	return cu
}

// SetTransactionType sets the "TransactionType" field.
func (cu *ConfigpointUpdate) SetTransactionType(s string) *ConfigpointUpdate {
	cu.mutation.SetTransactionType(s)
	return cu
}

// SetNillableTransactionType sets the "TransactionType" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableTransactionType(s *string) *ConfigpointUpdate {
	if s != nil {
		cu.SetTransactionType(*s)
	}
	return cu
}

// ClearTransactionType clears the value of the "TransactionType" field.
func (cu *ConfigpointUpdate) ClearTransactionType() *ConfigpointUpdate {
	cu.mutation.ClearTransactionType()
	return cu
}

// SetPaymentChannel sets the "PaymentChannel" field.
func (cu *ConfigpointUpdate) SetPaymentChannel(s string) *ConfigpointUpdate {
	cu.mutation.SetPaymentChannel(s)
	return cu
}

// SetNillablePaymentChannel sets the "PaymentChannel" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillablePaymentChannel(s *string) *ConfigpointUpdate {
	if s != nil {
		cu.SetPaymentChannel(*s)
	}
	return cu
}

// ClearPaymentChannel clears the value of the "PaymentChannel" field.
func (cu *ConfigpointUpdate) ClearPaymentChannel() *ConfigpointUpdate {
	cu.mutation.ClearPaymentChannel()
	return cu
}

// SetPaymentType sets the "PaymentType" field.
func (cu *ConfigpointUpdate) SetPaymentType(s string) *ConfigpointUpdate {
	cu.mutation.SetPaymentType(s)
	return cu
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillablePaymentType(s *string) *ConfigpointUpdate {
	if s != nil {
		cu.SetPaymentType(*s)
	}
	return cu
}

// ClearPaymentType clears the value of the "PaymentType" field.
func (cu *ConfigpointUpdate) ClearPaymentType() *ConfigpointUpdate {
	cu.mutation.ClearPaymentType()
	return cu
}

// SetDummyWallet sets the "DummyWallet" field.
func (cu *ConfigpointUpdate) SetDummyWallet(s string) *ConfigpointUpdate {
	cu.mutation.SetDummyWallet(s)
	return cu
}

// SetNillableDummyWallet sets the "DummyWallet" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableDummyWallet(s *string) *ConfigpointUpdate {
	if s != nil {
		cu.SetDummyWallet(*s)
	}
	return cu
}

// ClearDummyWallet clears the value of the "DummyWallet" field.
func (cu *ConfigpointUpdate) ClearDummyWallet() *ConfigpointUpdate {
	cu.mutation.ClearDummyWallet()
	return cu
}

// SetAmount sets the "Amount" field.
func (cu *ConfigpointUpdate) SetAmount(i int) *ConfigpointUpdate {
	cu.mutation.ResetAmount()
	cu.mutation.SetAmount(i)
	return cu
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableAmount(i *int) *ConfigpointUpdate {
	if i != nil {
		cu.SetAmount(*i)
	}
	return cu
}

// AddAmount adds i to the "Amount" field.
func (cu *ConfigpointUpdate) AddAmount(i int) *ConfigpointUpdate {
	cu.mutation.AddAmount(i)
	return cu
}

// ClearAmount clears the value of the "Amount" field.
func (cu *ConfigpointUpdate) ClearAmount() *ConfigpointUpdate {
	cu.mutation.ClearAmount()
	return cu
}

// SetPoint sets the "Point" field.
func (cu *ConfigpointUpdate) SetPoint(i int) *ConfigpointUpdate {
	cu.mutation.ResetPoint()
	cu.mutation.SetPoint(i)
	return cu
}

// SetNillablePoint sets the "Point" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillablePoint(i *int) *ConfigpointUpdate {
	if i != nil {
		cu.SetPoint(*i)
	}
	return cu
}

// AddPoint adds i to the "Point" field.
func (cu *ConfigpointUpdate) AddPoint(i int) *ConfigpointUpdate {
	cu.mutation.AddPoint(i)
	return cu
}

// ClearPoint clears the value of the "Point" field.
func (cu *ConfigpointUpdate) ClearPoint() *ConfigpointUpdate {
	cu.mutation.ClearPoint()
	return cu
}

// SetExpire sets the "Expire" field.
func (cu *ConfigpointUpdate) SetExpire(i int) *ConfigpointUpdate {
	cu.mutation.ResetExpire()
	cu.mutation.SetExpire(i)
	return cu
}

// SetNillableExpire sets the "Expire" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableExpire(i *int) *ConfigpointUpdate {
	if i != nil {
		cu.SetExpire(*i)
	}
	return cu
}

// AddExpire adds i to the "Expire" field.
func (cu *ConfigpointUpdate) AddExpire(i int) *ConfigpointUpdate {
	cu.mutation.AddExpire(i)
	return cu
}

// ClearExpire clears the value of the "Expire" field.
func (cu *ConfigpointUpdate) ClearExpire() *ConfigpointUpdate {
	cu.mutation.ClearExpire()
	return cu
}

// SetUpdateDate sets the "UpdateDate" field.
func (cu *ConfigpointUpdate) SetUpdateDate(t time.Time) *ConfigpointUpdate {
	cu.mutation.SetUpdateDate(t)
	return cu
}

// SetNillableUpdateDate sets the "UpdateDate" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableUpdateDate(t *time.Time) *ConfigpointUpdate {
	if t != nil {
		cu.SetUpdateDate(*t)
	}
	return cu
}

// ClearUpdateDate clears the value of the "UpdateDate" field.
func (cu *ConfigpointUpdate) ClearUpdateDate() *ConfigpointUpdate {
	cu.mutation.ClearUpdateDate()
	return cu
}

// SetExpireDate sets the "ExpireDate" field.
func (cu *ConfigpointUpdate) SetExpireDate(t time.Time) *ConfigpointUpdate {
	cu.mutation.SetExpireDate(t)
	return cu
}

// SetNillableExpireDate sets the "ExpireDate" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableExpireDate(t *time.Time) *ConfigpointUpdate {
	if t != nil {
		cu.SetExpireDate(*t)
	}
	return cu
}

// ClearExpireDate clears the value of the "ExpireDate" field.
func (cu *ConfigpointUpdate) ClearExpireDate() *ConfigpointUpdate {
	cu.mutation.ClearExpireDate()
	return cu
}

// SetStatusTransaction sets the "StatusTransaction" field.
func (cu *ConfigpointUpdate) SetStatusTransaction(s string) *ConfigpointUpdate {
	cu.mutation.SetStatusTransaction(s)
	return cu
}

// SetNillableStatusTransaction sets the "StatusTransaction" field if the given value is not nil.
func (cu *ConfigpointUpdate) SetNillableStatusTransaction(s *string) *ConfigpointUpdate {
	if s != nil {
		cu.SetStatusTransaction(*s)
	}
	return cu
}

// ClearStatusTransaction clears the value of the "StatusTransaction" field.
func (cu *ConfigpointUpdate) ClearStatusTransaction() *ConfigpointUpdate {
	cu.mutation.ClearStatusTransaction()
	return cu
}

// Mutation returns the ConfigpointMutation object of the builder.
func (cu *ConfigpointUpdate) Mutation() *ConfigpointMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfigpointUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfigpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfigpointUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfigpointUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfigpointUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConfigpointUpdate) check() error {
	if v, ok := cu.mutation.TransactionName(); ok {
		if err := configpoint.TransactionNameValidator(v); err != nil {
			return &ValidationError{Name: "TransactionName", err: fmt.Errorf("ent: validator failed for field \"TransactionName\": %w", err)}
		}
	}
	if v, ok := cu.mutation.TransactionType(); ok {
		if err := configpoint.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "TransactionType", err: fmt.Errorf("ent: validator failed for field \"TransactionType\": %w", err)}
		}
	}
	if v, ok := cu.mutation.PaymentChannel(); ok {
		if err := configpoint.PaymentChannelValidator(v); err != nil {
			return &ValidationError{Name: "PaymentChannel", err: fmt.Errorf("ent: validator failed for field \"PaymentChannel\": %w", err)}
		}
	}
	if v, ok := cu.mutation.PaymentType(); ok {
		if err := configpoint.PaymentTypeValidator(v); err != nil {
			return &ValidationError{Name: "PaymentType", err: fmt.Errorf("ent: validator failed for field \"PaymentType\": %w", err)}
		}
	}
	if v, ok := cu.mutation.DummyWallet(); ok {
		if err := configpoint.DummyWalletValidator(v); err != nil {
			return &ValidationError{Name: "DummyWallet", err: fmt.Errorf("ent: validator failed for field \"DummyWallet\": %w", err)}
		}
	}
	if v, ok := cu.mutation.StatusTransaction(); ok {
		if err := configpoint.StatusTransactionValidator(v); err != nil {
			return &ValidationError{Name: "StatusTransaction", err: fmt.Errorf("ent: validator failed for field \"StatusTransaction\": %w", err)}
		}
	}
	return nil
}

func (cu *ConfigpointUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   configpoint.Table,
			Columns: configpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configpoint.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.TransactionName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldTransactionName,
		})
	}
	if cu.mutation.TransactionNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldTransactionName,
		})
	}
	if value, ok := cu.mutation.TransactionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldTransactionType,
		})
	}
	if cu.mutation.TransactionTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldTransactionType,
		})
	}
	if value, ok := cu.mutation.PaymentChannel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldPaymentChannel,
		})
	}
	if cu.mutation.PaymentChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldPaymentChannel,
		})
	}
	if value, ok := cu.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldPaymentType,
		})
	}
	if cu.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldPaymentType,
		})
	}
	if value, ok := cu.mutation.DummyWallet(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldDummyWallet,
		})
	}
	if cu.mutation.DummyWalletCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldDummyWallet,
		})
	}
	if value, ok := cu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldAmount,
		})
	}
	if value, ok := cu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldAmount,
		})
	}
	if cu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: configpoint.FieldAmount,
		})
	}
	if value, ok := cu.mutation.Point(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldPoint,
		})
	}
	if value, ok := cu.mutation.AddedPoint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldPoint,
		})
	}
	if cu.mutation.PointCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: configpoint.FieldPoint,
		})
	}
	if value, ok := cu.mutation.Expire(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldExpire,
		})
	}
	if value, ok := cu.mutation.AddedExpire(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldExpire,
		})
	}
	if cu.mutation.ExpireCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: configpoint.FieldExpire,
		})
	}
	if value, ok := cu.mutation.UpdateDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configpoint.FieldUpdateDate,
		})
	}
	if cu.mutation.UpdateDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: configpoint.FieldUpdateDate,
		})
	}
	if value, ok := cu.mutation.ExpireDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configpoint.FieldExpireDate,
		})
	}
	if cu.mutation.ExpireDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: configpoint.FieldExpireDate,
		})
	}
	if value, ok := cu.mutation.StatusTransaction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldStatusTransaction,
		})
	}
	if cu.mutation.StatusTransactionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldStatusTransaction,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ConfigpointUpdateOne is the builder for updating a single Configpoint entity.
type ConfigpointUpdateOne struct {
	config
	hooks    []Hook
	mutation *ConfigpointMutation
}

// SetTransactionName sets the "TransactionName" field.
func (cuo *ConfigpointUpdateOne) SetTransactionName(s string) *ConfigpointUpdateOne {
	cuo.mutation.SetTransactionName(s)
	return cuo
}

// SetNillableTransactionName sets the "TransactionName" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableTransactionName(s *string) *ConfigpointUpdateOne {
	if s != nil {
		cuo.SetTransactionName(*s)
	}
	return cuo
}

// ClearTransactionName clears the value of the "TransactionName" field.
func (cuo *ConfigpointUpdateOne) ClearTransactionName() *ConfigpointUpdateOne {
	cuo.mutation.ClearTransactionName()
	return cuo
}

// SetTransactionType sets the "TransactionType" field.
func (cuo *ConfigpointUpdateOne) SetTransactionType(s string) *ConfigpointUpdateOne {
	cuo.mutation.SetTransactionType(s)
	return cuo
}

// SetNillableTransactionType sets the "TransactionType" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableTransactionType(s *string) *ConfigpointUpdateOne {
	if s != nil {
		cuo.SetTransactionType(*s)
	}
	return cuo
}

// ClearTransactionType clears the value of the "TransactionType" field.
func (cuo *ConfigpointUpdateOne) ClearTransactionType() *ConfigpointUpdateOne {
	cuo.mutation.ClearTransactionType()
	return cuo
}

// SetPaymentChannel sets the "PaymentChannel" field.
func (cuo *ConfigpointUpdateOne) SetPaymentChannel(s string) *ConfigpointUpdateOne {
	cuo.mutation.SetPaymentChannel(s)
	return cuo
}

// SetNillablePaymentChannel sets the "PaymentChannel" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillablePaymentChannel(s *string) *ConfigpointUpdateOne {
	if s != nil {
		cuo.SetPaymentChannel(*s)
	}
	return cuo
}

// ClearPaymentChannel clears the value of the "PaymentChannel" field.
func (cuo *ConfigpointUpdateOne) ClearPaymentChannel() *ConfigpointUpdateOne {
	cuo.mutation.ClearPaymentChannel()
	return cuo
}

// SetPaymentType sets the "PaymentType" field.
func (cuo *ConfigpointUpdateOne) SetPaymentType(s string) *ConfigpointUpdateOne {
	cuo.mutation.SetPaymentType(s)
	return cuo
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillablePaymentType(s *string) *ConfigpointUpdateOne {
	if s != nil {
		cuo.SetPaymentType(*s)
	}
	return cuo
}

// ClearPaymentType clears the value of the "PaymentType" field.
func (cuo *ConfigpointUpdateOne) ClearPaymentType() *ConfigpointUpdateOne {
	cuo.mutation.ClearPaymentType()
	return cuo
}

// SetDummyWallet sets the "DummyWallet" field.
func (cuo *ConfigpointUpdateOne) SetDummyWallet(s string) *ConfigpointUpdateOne {
	cuo.mutation.SetDummyWallet(s)
	return cuo
}

// SetNillableDummyWallet sets the "DummyWallet" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableDummyWallet(s *string) *ConfigpointUpdateOne {
	if s != nil {
		cuo.SetDummyWallet(*s)
	}
	return cuo
}

// ClearDummyWallet clears the value of the "DummyWallet" field.
func (cuo *ConfigpointUpdateOne) ClearDummyWallet() *ConfigpointUpdateOne {
	cuo.mutation.ClearDummyWallet()
	return cuo
}

// SetAmount sets the "Amount" field.
func (cuo *ConfigpointUpdateOne) SetAmount(i int) *ConfigpointUpdateOne {
	cuo.mutation.ResetAmount()
	cuo.mutation.SetAmount(i)
	return cuo
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableAmount(i *int) *ConfigpointUpdateOne {
	if i != nil {
		cuo.SetAmount(*i)
	}
	return cuo
}

// AddAmount adds i to the "Amount" field.
func (cuo *ConfigpointUpdateOne) AddAmount(i int) *ConfigpointUpdateOne {
	cuo.mutation.AddAmount(i)
	return cuo
}

// ClearAmount clears the value of the "Amount" field.
func (cuo *ConfigpointUpdateOne) ClearAmount() *ConfigpointUpdateOne {
	cuo.mutation.ClearAmount()
	return cuo
}

// SetPoint sets the "Point" field.
func (cuo *ConfigpointUpdateOne) SetPoint(i int) *ConfigpointUpdateOne {
	cuo.mutation.ResetPoint()
	cuo.mutation.SetPoint(i)
	return cuo
}

// SetNillablePoint sets the "Point" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillablePoint(i *int) *ConfigpointUpdateOne {
	if i != nil {
		cuo.SetPoint(*i)
	}
	return cuo
}

// AddPoint adds i to the "Point" field.
func (cuo *ConfigpointUpdateOne) AddPoint(i int) *ConfigpointUpdateOne {
	cuo.mutation.AddPoint(i)
	return cuo
}

// ClearPoint clears the value of the "Point" field.
func (cuo *ConfigpointUpdateOne) ClearPoint() *ConfigpointUpdateOne {
	cuo.mutation.ClearPoint()
	return cuo
}

// SetExpire sets the "Expire" field.
func (cuo *ConfigpointUpdateOne) SetExpire(i int) *ConfigpointUpdateOne {
	cuo.mutation.ResetExpire()
	cuo.mutation.SetExpire(i)
	return cuo
}

// SetNillableExpire sets the "Expire" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableExpire(i *int) *ConfigpointUpdateOne {
	if i != nil {
		cuo.SetExpire(*i)
	}
	return cuo
}

// AddExpire adds i to the "Expire" field.
func (cuo *ConfigpointUpdateOne) AddExpire(i int) *ConfigpointUpdateOne {
	cuo.mutation.AddExpire(i)
	return cuo
}

// ClearExpire clears the value of the "Expire" field.
func (cuo *ConfigpointUpdateOne) ClearExpire() *ConfigpointUpdateOne {
	cuo.mutation.ClearExpire()
	return cuo
}

// SetUpdateDate sets the "UpdateDate" field.
func (cuo *ConfigpointUpdateOne) SetUpdateDate(t time.Time) *ConfigpointUpdateOne {
	cuo.mutation.SetUpdateDate(t)
	return cuo
}

// SetNillableUpdateDate sets the "UpdateDate" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableUpdateDate(t *time.Time) *ConfigpointUpdateOne {
	if t != nil {
		cuo.SetUpdateDate(*t)
	}
	return cuo
}

// ClearUpdateDate clears the value of the "UpdateDate" field.
func (cuo *ConfigpointUpdateOne) ClearUpdateDate() *ConfigpointUpdateOne {
	cuo.mutation.ClearUpdateDate()
	return cuo
}

// SetExpireDate sets the "ExpireDate" field.
func (cuo *ConfigpointUpdateOne) SetExpireDate(t time.Time) *ConfigpointUpdateOne {
	cuo.mutation.SetExpireDate(t)
	return cuo
}

// SetNillableExpireDate sets the "ExpireDate" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableExpireDate(t *time.Time) *ConfigpointUpdateOne {
	if t != nil {
		cuo.SetExpireDate(*t)
	}
	return cuo
}

// ClearExpireDate clears the value of the "ExpireDate" field.
func (cuo *ConfigpointUpdateOne) ClearExpireDate() *ConfigpointUpdateOne {
	cuo.mutation.ClearExpireDate()
	return cuo
}

// SetStatusTransaction sets the "StatusTransaction" field.
func (cuo *ConfigpointUpdateOne) SetStatusTransaction(s string) *ConfigpointUpdateOne {
	cuo.mutation.SetStatusTransaction(s)
	return cuo
}

// SetNillableStatusTransaction sets the "StatusTransaction" field if the given value is not nil.
func (cuo *ConfigpointUpdateOne) SetNillableStatusTransaction(s *string) *ConfigpointUpdateOne {
	if s != nil {
		cuo.SetStatusTransaction(*s)
	}
	return cuo
}

// ClearStatusTransaction clears the value of the "StatusTransaction" field.
func (cuo *ConfigpointUpdateOne) ClearStatusTransaction() *ConfigpointUpdateOne {
	cuo.mutation.ClearStatusTransaction()
	return cuo
}

// Mutation returns the ConfigpointMutation object of the builder.
func (cuo *ConfigpointUpdateOne) Mutation() *ConfigpointMutation {
	return cuo.mutation
}

// Save executes the query and returns the updated Configpoint entity.
func (cuo *ConfigpointUpdateOne) Save(ctx context.Context) (*Configpoint, error) {
	var (
		err  error
		node *Configpoint
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfigpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfigpointUpdateOne) SaveX(ctx context.Context) *Configpoint {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfigpointUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfigpointUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConfigpointUpdateOne) check() error {
	if v, ok := cuo.mutation.TransactionName(); ok {
		if err := configpoint.TransactionNameValidator(v); err != nil {
			return &ValidationError{Name: "TransactionName", err: fmt.Errorf("ent: validator failed for field \"TransactionName\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.TransactionType(); ok {
		if err := configpoint.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "TransactionType", err: fmt.Errorf("ent: validator failed for field \"TransactionType\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.PaymentChannel(); ok {
		if err := configpoint.PaymentChannelValidator(v); err != nil {
			return &ValidationError{Name: "PaymentChannel", err: fmt.Errorf("ent: validator failed for field \"PaymentChannel\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.PaymentType(); ok {
		if err := configpoint.PaymentTypeValidator(v); err != nil {
			return &ValidationError{Name: "PaymentType", err: fmt.Errorf("ent: validator failed for field \"PaymentType\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.DummyWallet(); ok {
		if err := configpoint.DummyWalletValidator(v); err != nil {
			return &ValidationError{Name: "DummyWallet", err: fmt.Errorf("ent: validator failed for field \"DummyWallet\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.StatusTransaction(); ok {
		if err := configpoint.StatusTransactionValidator(v); err != nil {
			return &ValidationError{Name: "StatusTransaction", err: fmt.Errorf("ent: validator failed for field \"StatusTransaction\": %w", err)}
		}
	}
	return nil
}

func (cuo *ConfigpointUpdateOne) sqlSave(ctx context.Context) (_node *Configpoint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   configpoint.Table,
			Columns: configpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configpoint.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Configpoint.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.TransactionName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldTransactionName,
		})
	}
	if cuo.mutation.TransactionNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldTransactionName,
		})
	}
	if value, ok := cuo.mutation.TransactionType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldTransactionType,
		})
	}
	if cuo.mutation.TransactionTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldTransactionType,
		})
	}
	if value, ok := cuo.mutation.PaymentChannel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldPaymentChannel,
		})
	}
	if cuo.mutation.PaymentChannelCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldPaymentChannel,
		})
	}
	if value, ok := cuo.mutation.PaymentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldPaymentType,
		})
	}
	if cuo.mutation.PaymentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldPaymentType,
		})
	}
	if value, ok := cuo.mutation.DummyWallet(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldDummyWallet,
		})
	}
	if cuo.mutation.DummyWalletCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldDummyWallet,
		})
	}
	if value, ok := cuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldAmount,
		})
	}
	if value, ok := cuo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldAmount,
		})
	}
	if cuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: configpoint.FieldAmount,
		})
	}
	if value, ok := cuo.mutation.Point(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldPoint,
		})
	}
	if value, ok := cuo.mutation.AddedPoint(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldPoint,
		})
	}
	if cuo.mutation.PointCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: configpoint.FieldPoint,
		})
	}
	if value, ok := cuo.mutation.Expire(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldExpire,
		})
	}
	if value, ok := cuo.mutation.AddedExpire(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldExpire,
		})
	}
	if cuo.mutation.ExpireCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: configpoint.FieldExpire,
		})
	}
	if value, ok := cuo.mutation.UpdateDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configpoint.FieldUpdateDate,
		})
	}
	if cuo.mutation.UpdateDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: configpoint.FieldUpdateDate,
		})
	}
	if value, ok := cuo.mutation.ExpireDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configpoint.FieldExpireDate,
		})
	}
	if cuo.mutation.ExpireDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: configpoint.FieldExpireDate,
		})
	}
	if value, ok := cuo.mutation.StatusTransaction(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldStatusTransaction,
		})
	}
	if cuo.mutation.StatusTransactionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: configpoint.FieldStatusTransaction,
		})
	}
	_node = &Configpoint{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{configpoint.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
