// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/merchanttransaction"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MerchantTransactionCreate is the builder for creating a MerchantTransaction entity.
type MerchantTransactionCreate struct {
	config
	mutation *MerchantTransactionMutation
	hooks    []Hook
}

// SetTransactionID sets the "transaction_id" field.
func (mtc *MerchantTransactionCreate) SetTransactionID(s string) *MerchantTransactionCreate {
	mtc.mutation.SetTransactionID(s)
	return mtc
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableTransactionID(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetTransactionID(*s)
	}
	return mtc
}

// SetDateTime sets the "dateTime" field.
func (mtc *MerchantTransactionCreate) SetDateTime(t time.Time) *MerchantTransactionCreate {
	mtc.mutation.SetDateTime(t)
	return mtc
}

// SetNillableDateTime sets the "dateTime" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableDateTime(t *time.Time) *MerchantTransactionCreate {
	if t != nil {
		mtc.SetDateTime(*t)
	}
	return mtc
}

// SetAmount sets the "amount" field.
func (mtc *MerchantTransactionCreate) SetAmount(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetAmount(f)
	return mtc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableAmount(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetAmount(*f)
	}
	return mtc
}

// SetPaymentType sets the "PaymentType" field.
func (mtc *MerchantTransactionCreate) SetPaymentType(s string) *MerchantTransactionCreate {
	mtc.mutation.SetPaymentType(s)
	return mtc
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillablePaymentType(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetPaymentType(*s)
	}
	return mtc
}

// SetPaymentChannel sets the "PaymentChannel" field.
func (mtc *MerchantTransactionCreate) SetPaymentChannel(s string) *MerchantTransactionCreate {
	mtc.mutation.SetPaymentChannel(s)
	return mtc
}

// SetNillablePaymentChannel sets the "PaymentChannel" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillablePaymentChannel(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetPaymentChannel(*s)
	}
	return mtc
}

// SetStatus sets the "Status" field.
func (mtc *MerchantTransactionCreate) SetStatus(s string) *MerchantTransactionCreate {
	mtc.mutation.SetStatus(s)
	return mtc
}

// SetNillableStatus sets the "Status" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableStatus(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetStatus(*s)
	}
	return mtc
}

// SetMerchantID sets the "MerchantID" field.
func (mtc *MerchantTransactionCreate) SetMerchantID(s string) *MerchantTransactionCreate {
	mtc.mutation.SetMerchantID(s)
	return mtc
}

// SetNillableMerchantID sets the "MerchantID" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableMerchantID(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetMerchantID(*s)
	}
	return mtc
}

// SetTerminalID sets the "TerminalID" field.
func (mtc *MerchantTransactionCreate) SetTerminalID(s string) *MerchantTransactionCreate {
	mtc.mutation.SetTerminalID(s)
	return mtc
}

// SetNillableTerminalID sets the "TerminalID" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableTerminalID(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetTerminalID(*s)
	}
	return mtc
}

// SetMerchantFullName sets the "MerchantFullName" field.
func (mtc *MerchantTransactionCreate) SetMerchantFullName(s string) *MerchantTransactionCreate {
	mtc.mutation.SetMerchantFullName(s)
	return mtc
}

// SetNillableMerchantFullName sets the "MerchantFullName" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableMerchantFullName(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetMerchantFullName(*s)
	}
	return mtc
}

// SetFromAccount sets the "FromAccount" field.
func (mtc *MerchantTransactionCreate) SetFromAccount(s string) *MerchantTransactionCreate {
	mtc.mutation.SetFromAccount(s)
	return mtc
}

// SetNillableFromAccount sets the "FromAccount" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableFromAccount(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetFromAccount(*s)
	}
	return mtc
}

// SetSettlementAccount sets the "SettlementAccount" field.
func (mtc *MerchantTransactionCreate) SetSettlementAccount(s string) *MerchantTransactionCreate {
	mtc.mutation.SetSettlementAccount(s)
	return mtc
}

// SetNillableSettlementAccount sets the "SettlementAccount" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableSettlementAccount(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetSettlementAccount(*s)
	}
	return mtc
}

// SetMDRFEETHB sets the "MDR_FEETHB" field.
func (mtc *MerchantTransactionCreate) SetMDRFEETHB(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetMDRFEETHB(f)
	return mtc
}

// SetNillableMDRFEETHB sets the "MDR_FEETHB" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableMDRFEETHB(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetMDRFEETHB(*f)
	}
	return mtc
}

// SetTransactionFEETHB sets the "TransactionFEETHB" field.
func (mtc *MerchantTransactionCreate) SetTransactionFEETHB(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetTransactionFEETHB(f)
	return mtc
}

// SetNillableTransactionFEETHB sets the "TransactionFEETHB" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableTransactionFEETHB(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetTransactionFEETHB(*f)
	}
	return mtc
}

// SetTotalFeeincVAT sets the "TotalFeeincVAT" field.
func (mtc *MerchantTransactionCreate) SetTotalFeeincVAT(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetTotalFeeincVAT(f)
	return mtc
}

// SetNillableTotalFeeincVAT sets the "TotalFeeincVAT" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableTotalFeeincVAT(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetTotalFeeincVAT(*f)
	}
	return mtc
}

// SetVATTHB sets the "VATTHB" field.
func (mtc *MerchantTransactionCreate) SetVATTHB(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetVATTHB(f)
	return mtc
}

// SetNillableVATTHB sets the "VATTHB" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableVATTHB(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetVATTHB(*f)
	}
	return mtc
}

// SetTotalFeeExcVAT sets the "TotalFeeExcVAT" field.
func (mtc *MerchantTransactionCreate) SetTotalFeeExcVAT(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetTotalFeeExcVAT(f)
	return mtc
}

// SetNillableTotalFeeExcVAT sets the "TotalFeeExcVAT" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableTotalFeeExcVAT(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetTotalFeeExcVAT(*f)
	}
	return mtc
}

// SetWHTTHB sets the "WHTTHB" field.
func (mtc *MerchantTransactionCreate) SetWHTTHB(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetWHTTHB(f)
	return mtc
}

// SetNillableWHTTHB sets the "WHTTHB" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableWHTTHB(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetWHTTHB(*f)
	}
	return mtc
}

// SetNetPayTHB sets the "NetPayTHB" field.
func (mtc *MerchantTransactionCreate) SetNetPayTHB(f float64) *MerchantTransactionCreate {
	mtc.mutation.SetNetPayTHB(f)
	return mtc
}

// SetNillableNetPayTHB sets the "NetPayTHB" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableNetPayTHB(f *float64) *MerchantTransactionCreate {
	if f != nil {
		mtc.SetNetPayTHB(*f)
	}
	return mtc
}

// SetTransactionType sets the "TransactionType" field.
func (mtc *MerchantTransactionCreate) SetTransactionType(s string) *MerchantTransactionCreate {
	mtc.mutation.SetTransactionType(s)
	return mtc
}

// SetNillableTransactionType sets the "TransactionType" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableTransactionType(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetTransactionType(*s)
	}
	return mtc
}

// SetBankCode sets the "BankCode" field.
func (mtc *MerchantTransactionCreate) SetBankCode(s string) *MerchantTransactionCreate {
	mtc.mutation.SetBankCode(s)
	return mtc
}

// SetNillableBankCode sets the "BankCode" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableBankCode(s *string) *MerchantTransactionCreate {
	if s != nil {
		mtc.SetBankCode(*s)
	}
	return mtc
}

// SetFileimportID sets the "FileimportID" field.
func (mtc *MerchantTransactionCreate) SetFileimportID(i int) *MerchantTransactionCreate {
	mtc.mutation.SetFileimportID(i)
	return mtc
}

// SetNillableFileimportID sets the "FileimportID" field if the given value is not nil.
func (mtc *MerchantTransactionCreate) SetNillableFileimportID(i *int) *MerchantTransactionCreate {
	if i != nil {
		mtc.SetFileimportID(*i)
	}
	return mtc
}

// SetID sets the "id" field.
func (mtc *MerchantTransactionCreate) SetID(i int) *MerchantTransactionCreate {
	mtc.mutation.SetID(i)
	return mtc
}

// Mutation returns the MerchantTransactionMutation object of the builder.
func (mtc *MerchantTransactionCreate) Mutation() *MerchantTransactionMutation {
	return mtc.mutation
}

// Save creates the MerchantTransaction in the database.
func (mtc *MerchantTransactionCreate) Save(ctx context.Context) (*MerchantTransaction, error) {
	var (
		err  error
		node *MerchantTransaction
	)
	if len(mtc.hooks) == 0 {
		if err = mtc.check(); err != nil {
			return nil, err
		}
		node, err = mtc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MerchantTransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mtc.check(); err != nil {
				return nil, err
			}
			mtc.mutation = mutation
			node, err = mtc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mtc.hooks) - 1; i >= 0; i-- {
			mut = mtc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mtc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mtc *MerchantTransactionCreate) SaveX(ctx context.Context) *MerchantTransaction {
	v, err := mtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (mtc *MerchantTransactionCreate) check() error {
	if v, ok := mtc.mutation.TransactionID(); ok {
		if err := merchanttransaction.TransactionIDValidator(v); err != nil {
			return &ValidationError{Name: "transaction_id", err: fmt.Errorf("ent: validator failed for field \"transaction_id\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.PaymentType(); ok {
		if err := merchanttransaction.PaymentTypeValidator(v); err != nil {
			return &ValidationError{Name: "PaymentType", err: fmt.Errorf("ent: validator failed for field \"PaymentType\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.PaymentChannel(); ok {
		if err := merchanttransaction.PaymentChannelValidator(v); err != nil {
			return &ValidationError{Name: "PaymentChannel", err: fmt.Errorf("ent: validator failed for field \"PaymentChannel\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.Status(); ok {
		if err := merchanttransaction.StatusValidator(v); err != nil {
			return &ValidationError{Name: "Status", err: fmt.Errorf("ent: validator failed for field \"Status\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.MerchantID(); ok {
		if err := merchanttransaction.MerchantIDValidator(v); err != nil {
			return &ValidationError{Name: "MerchantID", err: fmt.Errorf("ent: validator failed for field \"MerchantID\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.TerminalID(); ok {
		if err := merchanttransaction.TerminalIDValidator(v); err != nil {
			return &ValidationError{Name: "TerminalID", err: fmt.Errorf("ent: validator failed for field \"TerminalID\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.MerchantFullName(); ok {
		if err := merchanttransaction.MerchantFullNameValidator(v); err != nil {
			return &ValidationError{Name: "MerchantFullName", err: fmt.Errorf("ent: validator failed for field \"MerchantFullName\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.FromAccount(); ok {
		if err := merchanttransaction.FromAccountValidator(v); err != nil {
			return &ValidationError{Name: "FromAccount", err: fmt.Errorf("ent: validator failed for field \"FromAccount\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.SettlementAccount(); ok {
		if err := merchanttransaction.SettlementAccountValidator(v); err != nil {
			return &ValidationError{Name: "SettlementAccount", err: fmt.Errorf("ent: validator failed for field \"SettlementAccount\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.TransactionType(); ok {
		if err := merchanttransaction.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "TransactionType", err: fmt.Errorf("ent: validator failed for field \"TransactionType\": %w", err)}
		}
	}
	if v, ok := mtc.mutation.BankCode(); ok {
		if err := merchanttransaction.BankCodeValidator(v); err != nil {
			return &ValidationError{Name: "BankCode", err: fmt.Errorf("ent: validator failed for field \"BankCode\": %w", err)}
		}
	}
	return nil
}

func (mtc *MerchantTransactionCreate) sqlSave(ctx context.Context) (*MerchantTransaction, error) {
	_node, _spec := mtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mtc.driver, _spec); err != nil {
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

func (mtc *MerchantTransactionCreate) createSpec() (*MerchantTransaction, *sqlgraph.CreateSpec) {
	var (
		_node = &MerchantTransaction{config: mtc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: merchanttransaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: merchanttransaction.FieldID,
			},
		}
	)
	if id, ok := mtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mtc.mutation.TransactionID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldTransactionID,
		})
		_node.TransactionID = &value
	}
	if value, ok := mtc.mutation.DateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: merchanttransaction.FieldDateTime,
		})
		_node.DateTime = &value
	}
	if value, ok := mtc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldAmount,
		})
		_node.Amount = &value
	}
	if value, ok := mtc.mutation.PaymentType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldPaymentType,
		})
		_node.PaymentType = &value
	}
	if value, ok := mtc.mutation.PaymentChannel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldPaymentChannel,
		})
		_node.PaymentChannel = &value
	}
	if value, ok := mtc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldStatus,
		})
		_node.Status = &value
	}
	if value, ok := mtc.mutation.MerchantID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldMerchantID,
		})
		_node.MerchantID = &value
	}
	if value, ok := mtc.mutation.TerminalID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldTerminalID,
		})
		_node.TerminalID = &value
	}
	if value, ok := mtc.mutation.MerchantFullName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldMerchantFullName,
		})
		_node.MerchantFullName = &value
	}
	if value, ok := mtc.mutation.FromAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldFromAccount,
		})
		_node.FromAccount = &value
	}
	if value, ok := mtc.mutation.SettlementAccount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldSettlementAccount,
		})
		_node.SettlementAccount = &value
	}
	if value, ok := mtc.mutation.MDRFEETHB(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldMDRFEETHB,
		})
		_node.MDRFEETHB = &value
	}
	if value, ok := mtc.mutation.TransactionFEETHB(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldTransactionFEETHB,
		})
		_node.TransactionFEETHB = &value
	}
	if value, ok := mtc.mutation.TotalFeeincVAT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldTotalFeeincVAT,
		})
		_node.TotalFeeincVAT = &value
	}
	if value, ok := mtc.mutation.VATTHB(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldVATTHB,
		})
		_node.VATTHB = &value
	}
	if value, ok := mtc.mutation.TotalFeeExcVAT(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldTotalFeeExcVAT,
		})
		_node.TotalFeeExcVAT = &value
	}
	if value, ok := mtc.mutation.WHTTHB(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldWHTTHB,
		})
		_node.WHTTHB = &value
	}
	if value, ok := mtc.mutation.NetPayTHB(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: merchanttransaction.FieldNetPayTHB,
		})
		_node.NetPayTHB = &value
	}
	if value, ok := mtc.mutation.TransactionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldTransactionType,
		})
		_node.TransactionType = &value
	}
	if value, ok := mtc.mutation.BankCode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: merchanttransaction.FieldBankCode,
		})
		_node.BankCode = &value
	}
	if value, ok := mtc.mutation.FileimportID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: merchanttransaction.FieldFileimportID,
		})
		_node.FileimportID = &value
	}
	return _node, _spec
}

// MerchantTransactionCreateBulk is the builder for creating many MerchantTransaction entities in bulk.
type MerchantTransactionCreateBulk struct {
	config
	builders []*MerchantTransactionCreate
}

// Save creates the MerchantTransaction entities in the database.
func (mtcb *MerchantTransactionCreateBulk) Save(ctx context.Context) ([]*MerchantTransaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mtcb.builders))
	nodes := make([]*MerchantTransaction, len(mtcb.builders))
	mutators := make([]Mutator, len(mtcb.builders))
	for i := range mtcb.builders {
		func(i int, root context.Context) {
			builder := mtcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MerchantTransactionMutation)
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
					_, err = mutators[i+1].Mutate(root, mtcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mtcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mtcb *MerchantTransactionCreateBulk) SaveX(ctx context.Context) []*MerchantTransaction {
	v, err := mtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
