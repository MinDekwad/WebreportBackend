// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/configpoint"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ConfigpointCreate is the builder for creating a Configpoint entity.
type ConfigpointCreate struct {
	config
	mutation *ConfigpointMutation
	hooks    []Hook
}

// SetTransactionName sets the "TransactionName" field.
func (cc *ConfigpointCreate) SetTransactionName(s string) *ConfigpointCreate {
	cc.mutation.SetTransactionName(s)
	return cc
}

// SetNillableTransactionName sets the "TransactionName" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableTransactionName(s *string) *ConfigpointCreate {
	if s != nil {
		cc.SetTransactionName(*s)
	}
	return cc
}

// SetTransactionType sets the "TransactionType" field.
func (cc *ConfigpointCreate) SetTransactionType(s string) *ConfigpointCreate {
	cc.mutation.SetTransactionType(s)
	return cc
}

// SetNillableTransactionType sets the "TransactionType" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableTransactionType(s *string) *ConfigpointCreate {
	if s != nil {
		cc.SetTransactionType(*s)
	}
	return cc
}

// SetPaymentChannel sets the "PaymentChannel" field.
func (cc *ConfigpointCreate) SetPaymentChannel(s string) *ConfigpointCreate {
	cc.mutation.SetPaymentChannel(s)
	return cc
}

// SetNillablePaymentChannel sets the "PaymentChannel" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillablePaymentChannel(s *string) *ConfigpointCreate {
	if s != nil {
		cc.SetPaymentChannel(*s)
	}
	return cc
}

// SetPaymentType sets the "PaymentType" field.
func (cc *ConfigpointCreate) SetPaymentType(s string) *ConfigpointCreate {
	cc.mutation.SetPaymentType(s)
	return cc
}

// SetNillablePaymentType sets the "PaymentType" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillablePaymentType(s *string) *ConfigpointCreate {
	if s != nil {
		cc.SetPaymentType(*s)
	}
	return cc
}

// SetDummyWallet sets the "DummyWallet" field.
func (cc *ConfigpointCreate) SetDummyWallet(s string) *ConfigpointCreate {
	cc.mutation.SetDummyWallet(s)
	return cc
}

// SetNillableDummyWallet sets the "DummyWallet" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableDummyWallet(s *string) *ConfigpointCreate {
	if s != nil {
		cc.SetDummyWallet(*s)
	}
	return cc
}

// SetAmount sets the "Amount" field.
func (cc *ConfigpointCreate) SetAmount(i int) *ConfigpointCreate {
	cc.mutation.SetAmount(i)
	return cc
}

// SetNillableAmount sets the "Amount" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableAmount(i *int) *ConfigpointCreate {
	if i != nil {
		cc.SetAmount(*i)
	}
	return cc
}

// SetPoint sets the "Point" field.
func (cc *ConfigpointCreate) SetPoint(i int) *ConfigpointCreate {
	cc.mutation.SetPoint(i)
	return cc
}

// SetNillablePoint sets the "Point" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillablePoint(i *int) *ConfigpointCreate {
	if i != nil {
		cc.SetPoint(*i)
	}
	return cc
}

// SetExpire sets the "Expire" field.
func (cc *ConfigpointCreate) SetExpire(i int) *ConfigpointCreate {
	cc.mutation.SetExpire(i)
	return cc
}

// SetNillableExpire sets the "Expire" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableExpire(i *int) *ConfigpointCreate {
	if i != nil {
		cc.SetExpire(*i)
	}
	return cc
}

// SetUpdateDate sets the "UpdateDate" field.
func (cc *ConfigpointCreate) SetUpdateDate(t time.Time) *ConfigpointCreate {
	cc.mutation.SetUpdateDate(t)
	return cc
}

// SetNillableUpdateDate sets the "UpdateDate" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableUpdateDate(t *time.Time) *ConfigpointCreate {
	if t != nil {
		cc.SetUpdateDate(*t)
	}
	return cc
}

// SetExpireDate sets the "ExpireDate" field.
func (cc *ConfigpointCreate) SetExpireDate(t time.Time) *ConfigpointCreate {
	cc.mutation.SetExpireDate(t)
	return cc
}

// SetNillableExpireDate sets the "ExpireDate" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableExpireDate(t *time.Time) *ConfigpointCreate {
	if t != nil {
		cc.SetExpireDate(*t)
	}
	return cc
}

// SetStatusTransaction sets the "StatusTransaction" field.
func (cc *ConfigpointCreate) SetStatusTransaction(s string) *ConfigpointCreate {
	cc.mutation.SetStatusTransaction(s)
	return cc
}

// SetNillableStatusTransaction sets the "StatusTransaction" field if the given value is not nil.
func (cc *ConfigpointCreate) SetNillableStatusTransaction(s *string) *ConfigpointCreate {
	if s != nil {
		cc.SetStatusTransaction(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ConfigpointCreate) SetID(i int) *ConfigpointCreate {
	cc.mutation.SetID(i)
	return cc
}

// Mutation returns the ConfigpointMutation object of the builder.
func (cc *ConfigpointCreate) Mutation() *ConfigpointMutation {
	return cc.mutation
}

// Save creates the Configpoint in the database.
func (cc *ConfigpointCreate) Save(ctx context.Context) (*Configpoint, error) {
	var (
		err  error
		node *Configpoint
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConfigpointMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			node, err = cc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ConfigpointCreate) SaveX(ctx context.Context) *Configpoint {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConfigpointCreate) check() error {
	if v, ok := cc.mutation.TransactionName(); ok {
		if err := configpoint.TransactionNameValidator(v); err != nil {
			return &ValidationError{Name: "TransactionName", err: fmt.Errorf("ent: validator failed for field \"TransactionName\": %w", err)}
		}
	}
	if v, ok := cc.mutation.TransactionType(); ok {
		if err := configpoint.TransactionTypeValidator(v); err != nil {
			return &ValidationError{Name: "TransactionType", err: fmt.Errorf("ent: validator failed for field \"TransactionType\": %w", err)}
		}
	}
	if v, ok := cc.mutation.PaymentChannel(); ok {
		if err := configpoint.PaymentChannelValidator(v); err != nil {
			return &ValidationError{Name: "PaymentChannel", err: fmt.Errorf("ent: validator failed for field \"PaymentChannel\": %w", err)}
		}
	}
	if v, ok := cc.mutation.PaymentType(); ok {
		if err := configpoint.PaymentTypeValidator(v); err != nil {
			return &ValidationError{Name: "PaymentType", err: fmt.Errorf("ent: validator failed for field \"PaymentType\": %w", err)}
		}
	}
	if v, ok := cc.mutation.DummyWallet(); ok {
		if err := configpoint.DummyWalletValidator(v); err != nil {
			return &ValidationError{Name: "DummyWallet", err: fmt.Errorf("ent: validator failed for field \"DummyWallet\": %w", err)}
		}
	}
	if v, ok := cc.mutation.StatusTransaction(); ok {
		if err := configpoint.StatusTransactionValidator(v); err != nil {
			return &ValidationError{Name: "StatusTransaction", err: fmt.Errorf("ent: validator failed for field \"StatusTransaction\": %w", err)}
		}
	}
	return nil
}

func (cc *ConfigpointCreate) sqlSave(ctx context.Context) (*Configpoint, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ConfigpointCreate) createSpec() (*Configpoint, *sqlgraph.CreateSpec) {
	var (
		_node = &Configpoint{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: configpoint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: configpoint.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.TransactionName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldTransactionName,
		})
		_node.TransactionName = value
	}
	if value, ok := cc.mutation.TransactionType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldTransactionType,
		})
		_node.TransactionType = &value
	}
	if value, ok := cc.mutation.PaymentChannel(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldPaymentChannel,
		})
		_node.PaymentChannel = &value
	}
	if value, ok := cc.mutation.PaymentType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldPaymentType,
		})
		_node.PaymentType = &value
	}
	if value, ok := cc.mutation.DummyWallet(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldDummyWallet,
		})
		_node.DummyWallet = &value
	}
	if value, ok := cc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldAmount,
		})
		_node.Amount = value
	}
	if value, ok := cc.mutation.Point(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldPoint,
		})
		_node.Point = value
	}
	if value, ok := cc.mutation.Expire(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: configpoint.FieldExpire,
		})
		_node.Expire = value
	}
	if value, ok := cc.mutation.UpdateDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configpoint.FieldUpdateDate,
		})
		_node.UpdateDate = value
	}
	if value, ok := cc.mutation.ExpireDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: configpoint.FieldExpireDate,
		})
		_node.ExpireDate = value
	}
	if value, ok := cc.mutation.StatusTransaction(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: configpoint.FieldStatusTransaction,
		})
		_node.StatusTransaction = &value
	}
	return _node, _spec
}

// ConfigpointCreateBulk is the builder for creating many Configpoint entities in bulk.
type ConfigpointCreateBulk struct {
	config
	builders []*ConfigpointCreate
}

// Save creates the Configpoint entities in the database.
func (ccb *ConfigpointCreateBulk) Save(ctx context.Context) ([]*Configpoint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Configpoint, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConfigpointMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConfigpointCreateBulk) SaveX(ctx context.Context) []*Configpoint {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}