// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/agentkyc"
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AgentkycUpdate is the builder for updating Agentkyc entities.
type AgentkycUpdate struct {
	config
	hooks    []Hook
	mutation *AgentkycMutation
}

// Where adds a new predicate for the AgentkycUpdate builder.
func (au *AgentkycUpdate) Where(ps ...predicate.Agentkyc) *AgentkycUpdate {
	au.mutation.predicates = append(au.mutation.predicates, ps...)
	return au
}

// SetKYCDate sets the "KYCDate" field.
func (au *AgentkycUpdate) SetKYCDate(s string) *AgentkycUpdate {
	au.mutation.SetKYCDate(s)
	return au
}

// SetNillableKYCDate sets the "KYCDate" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableKYCDate(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetKYCDate(*s)
	}
	return au
}

// ClearKYCDate clears the value of the "KYCDate" field.
func (au *AgentkycUpdate) ClearKYCDate() *AgentkycUpdate {
	au.mutation.ClearKYCDate()
	return au
}

// SetKYCTime sets the "KYCTime" field.
func (au *AgentkycUpdate) SetKYCTime(s string) *AgentkycUpdate {
	au.mutation.SetKYCTime(s)
	return au
}

// SetNillableKYCTime sets the "KYCTime" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableKYCTime(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetKYCTime(*s)
	}
	return au
}

// ClearKYCTime clears the value of the "KYCTime" field.
func (au *AgentkycUpdate) ClearKYCTime() *AgentkycUpdate {
	au.mutation.ClearKYCTime()
	return au
}

// SetAgentID sets the "AgentID" field.
func (au *AgentkycUpdate) SetAgentID(s string) *AgentkycUpdate {
	au.mutation.SetAgentID(s)
	return au
}

// SetNillableAgentID sets the "AgentID" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableAgentID(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetAgentID(*s)
	}
	return au
}

// ClearAgentID clears the value of the "AgentID" field.
func (au *AgentkycUpdate) ClearAgentID() *AgentkycUpdate {
	au.mutation.ClearAgentID()
	return au
}

// SetAgentemail sets the "Agentemail" field.
func (au *AgentkycUpdate) SetAgentemail(s string) *AgentkycUpdate {
	au.mutation.SetAgentemail(s)
	return au
}

// SetNillableAgentemail sets the "Agentemail" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableAgentemail(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetAgentemail(*s)
	}
	return au
}

// ClearAgentemail clears the value of the "Agentemail" field.
func (au *AgentkycUpdate) ClearAgentemail() *AgentkycUpdate {
	au.mutation.ClearAgentemail()
	return au
}

// SetAgentNameLastname sets the "AgentNameLastname" field.
func (au *AgentkycUpdate) SetAgentNameLastname(s string) *AgentkycUpdate {
	au.mutation.SetAgentNameLastname(s)
	return au
}

// SetNillableAgentNameLastname sets the "AgentNameLastname" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableAgentNameLastname(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetAgentNameLastname(*s)
	}
	return au
}

// ClearAgentNameLastname clears the value of the "AgentNameLastname" field.
func (au *AgentkycUpdate) ClearAgentNameLastname() *AgentkycUpdate {
	au.mutation.ClearAgentNameLastname()
	return au
}

// SetKYCStatus sets the "KYCStatus" field.
func (au *AgentkycUpdate) SetKYCStatus(s string) *AgentkycUpdate {
	au.mutation.SetKYCStatus(s)
	return au
}

// SetNillableKYCStatus sets the "KYCStatus" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableKYCStatus(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetKYCStatus(*s)
	}
	return au
}

// ClearKYCStatus clears the value of the "KYCStatus" field.
func (au *AgentkycUpdate) ClearKYCStatus() *AgentkycUpdate {
	au.mutation.ClearKYCStatus()
	return au
}

// SetConsumerwalletid sets the "Consumerwalletid" field.
func (au *AgentkycUpdate) SetConsumerwalletid(s string) *AgentkycUpdate {
	au.mutation.SetConsumerwalletid(s)
	return au
}

// SetNillableConsumerwalletid sets the "Consumerwalletid" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableConsumerwalletid(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetConsumerwalletid(*s)
	}
	return au
}

// ClearConsumerwalletid clears the value of the "Consumerwalletid" field.
func (au *AgentkycUpdate) ClearConsumerwalletid() *AgentkycUpdate {
	au.mutation.ClearConsumerwalletid()
	return au
}

// SetKYCRespond sets the "KYCRespond" field.
func (au *AgentkycUpdate) SetKYCRespond(s string) *AgentkycUpdate {
	au.mutation.SetKYCRespond(s)
	return au
}

// SetNillableKYCRespond sets the "KYCRespond" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableKYCRespond(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetKYCRespond(*s)
	}
	return au
}

// ClearKYCRespond clears the value of the "KYCRespond" field.
func (au *AgentkycUpdate) ClearKYCRespond() *AgentkycUpdate {
	au.mutation.ClearKYCRespond()
	return au
}

// SetDOPARespond sets the "DOPARespond" field.
func (au *AgentkycUpdate) SetDOPARespond(s string) *AgentkycUpdate {
	au.mutation.SetDOPARespond(s)
	return au
}

// SetNillableDOPARespond sets the "DOPARespond" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableDOPARespond(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetDOPARespond(*s)
	}
	return au
}

// ClearDOPARespond clears the value of the "DOPARespond" field.
func (au *AgentkycUpdate) ClearDOPARespond() *AgentkycUpdate {
	au.mutation.ClearDOPARespond()
	return au
}

// SetAgentType sets the "AgentType" field.
func (au *AgentkycUpdate) SetAgentType(s string) *AgentkycUpdate {
	au.mutation.SetAgentType(s)
	return au
}

// SetNillableAgentType sets the "AgentType" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableAgentType(s *string) *AgentkycUpdate {
	if s != nil {
		au.SetAgentType(*s)
	}
	return au
}

// ClearAgentType clears the value of the "AgentType" field.
func (au *AgentkycUpdate) ClearAgentType() *AgentkycUpdate {
	au.mutation.ClearAgentType()
	return au
}

// SetFileimportID sets the "FileimportID" field.
func (au *AgentkycUpdate) SetFileimportID(i int) *AgentkycUpdate {
	au.mutation.ResetFileimportID()
	au.mutation.SetFileimportID(i)
	return au
}

// SetNillableFileimportID sets the "FileimportID" field if the given value is not nil.
func (au *AgentkycUpdate) SetNillableFileimportID(i *int) *AgentkycUpdate {
	if i != nil {
		au.SetFileimportID(*i)
	}
	return au
}

// AddFileimportID adds i to the "FileimportID" field.
func (au *AgentkycUpdate) AddFileimportID(i int) *AgentkycUpdate {
	au.mutation.AddFileimportID(i)
	return au
}

// ClearFileimportID clears the value of the "FileimportID" field.
func (au *AgentkycUpdate) ClearFileimportID() *AgentkycUpdate {
	au.mutation.ClearFileimportID()
	return au
}

// Mutation returns the AgentkycMutation object of the builder.
func (au *AgentkycUpdate) Mutation() *AgentkycMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AgentkycUpdate) Save(ctx context.Context) (int, error) {
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
			mutation, ok := m.(*AgentkycMutation)
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
func (au *AgentkycUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AgentkycUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AgentkycUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AgentkycUpdate) check() error {
	if v, ok := au.mutation.KYCDate(); ok {
		if err := agentkyc.KYCDateValidator(v); err != nil {
			return &ValidationError{Name: "KYCDate", err: fmt.Errorf("ent: validator failed for field \"KYCDate\": %w", err)}
		}
	}
	if v, ok := au.mutation.KYCTime(); ok {
		if err := agentkyc.KYCTimeValidator(v); err != nil {
			return &ValidationError{Name: "KYCTime", err: fmt.Errorf("ent: validator failed for field \"KYCTime\": %w", err)}
		}
	}
	if v, ok := au.mutation.AgentID(); ok {
		if err := agentkyc.AgentIDValidator(v); err != nil {
			return &ValidationError{Name: "AgentID", err: fmt.Errorf("ent: validator failed for field \"AgentID\": %w", err)}
		}
	}
	if v, ok := au.mutation.Agentemail(); ok {
		if err := agentkyc.AgentemailValidator(v); err != nil {
			return &ValidationError{Name: "Agentemail", err: fmt.Errorf("ent: validator failed for field \"Agentemail\": %w", err)}
		}
	}
	if v, ok := au.mutation.AgentNameLastname(); ok {
		if err := agentkyc.AgentNameLastnameValidator(v); err != nil {
			return &ValidationError{Name: "AgentNameLastname", err: fmt.Errorf("ent: validator failed for field \"AgentNameLastname\": %w", err)}
		}
	}
	if v, ok := au.mutation.KYCStatus(); ok {
		if err := agentkyc.KYCStatusValidator(v); err != nil {
			return &ValidationError{Name: "KYCStatus", err: fmt.Errorf("ent: validator failed for field \"KYCStatus\": %w", err)}
		}
	}
	if v, ok := au.mutation.Consumerwalletid(); ok {
		if err := agentkyc.ConsumerwalletidValidator(v); err != nil {
			return &ValidationError{Name: "Consumerwalletid", err: fmt.Errorf("ent: validator failed for field \"Consumerwalletid\": %w", err)}
		}
	}
	if v, ok := au.mutation.KYCRespond(); ok {
		if err := agentkyc.KYCRespondValidator(v); err != nil {
			return &ValidationError{Name: "KYCRespond", err: fmt.Errorf("ent: validator failed for field \"KYCRespond\": %w", err)}
		}
	}
	if v, ok := au.mutation.DOPARespond(); ok {
		if err := agentkyc.DOPARespondValidator(v); err != nil {
			return &ValidationError{Name: "DOPARespond", err: fmt.Errorf("ent: validator failed for field \"DOPARespond\": %w", err)}
		}
	}
	if v, ok := au.mutation.AgentType(); ok {
		if err := agentkyc.AgentTypeValidator(v); err != nil {
			return &ValidationError{Name: "AgentType", err: fmt.Errorf("ent: validator failed for field \"AgentType\": %w", err)}
		}
	}
	return nil
}

func (au *AgentkycUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agentkyc.Table,
			Columns: agentkyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agentkyc.FieldID,
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
	if value, ok := au.mutation.KYCDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCDate,
		})
	}
	if au.mutation.KYCDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCDate,
		})
	}
	if value, ok := au.mutation.KYCTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCTime,
		})
	}
	if au.mutation.KYCTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCTime,
		})
	}
	if value, ok := au.mutation.AgentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentID,
		})
	}
	if au.mutation.AgentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentID,
		})
	}
	if value, ok := au.mutation.Agentemail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentemail,
		})
	}
	if au.mutation.AgentemailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentemail,
		})
	}
	if value, ok := au.mutation.AgentNameLastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentNameLastname,
		})
	}
	if au.mutation.AgentNameLastnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentNameLastname,
		})
	}
	if value, ok := au.mutation.KYCStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCStatus,
		})
	}
	if au.mutation.KYCStatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCStatus,
		})
	}
	if value, ok := au.mutation.Consumerwalletid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldConsumerwalletid,
		})
	}
	if au.mutation.ConsumerwalletidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldConsumerwalletid,
		})
	}
	if value, ok := au.mutation.KYCRespond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCRespond,
		})
	}
	if au.mutation.KYCRespondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCRespond,
		})
	}
	if value, ok := au.mutation.DOPARespond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldDOPARespond,
		})
	}
	if au.mutation.DOPARespondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldDOPARespond,
		})
	}
	if value, ok := au.mutation.AgentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentType,
		})
	}
	if au.mutation.AgentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentType,
		})
	}
	if value, ok := au.mutation.FileimportID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agentkyc.FieldFileimportID,
		})
	}
	if value, ok := au.mutation.AddedFileimportID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agentkyc.FieldFileimportID,
		})
	}
	if au.mutation.FileimportIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: agentkyc.FieldFileimportID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agentkyc.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AgentkycUpdateOne is the builder for updating a single Agentkyc entity.
type AgentkycUpdateOne struct {
	config
	hooks    []Hook
	mutation *AgentkycMutation
}

// SetKYCDate sets the "KYCDate" field.
func (auo *AgentkycUpdateOne) SetKYCDate(s string) *AgentkycUpdateOne {
	auo.mutation.SetKYCDate(s)
	return auo
}

// SetNillableKYCDate sets the "KYCDate" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableKYCDate(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetKYCDate(*s)
	}
	return auo
}

// ClearKYCDate clears the value of the "KYCDate" field.
func (auo *AgentkycUpdateOne) ClearKYCDate() *AgentkycUpdateOne {
	auo.mutation.ClearKYCDate()
	return auo
}

// SetKYCTime sets the "KYCTime" field.
func (auo *AgentkycUpdateOne) SetKYCTime(s string) *AgentkycUpdateOne {
	auo.mutation.SetKYCTime(s)
	return auo
}

// SetNillableKYCTime sets the "KYCTime" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableKYCTime(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetKYCTime(*s)
	}
	return auo
}

// ClearKYCTime clears the value of the "KYCTime" field.
func (auo *AgentkycUpdateOne) ClearKYCTime() *AgentkycUpdateOne {
	auo.mutation.ClearKYCTime()
	return auo
}

// SetAgentID sets the "AgentID" field.
func (auo *AgentkycUpdateOne) SetAgentID(s string) *AgentkycUpdateOne {
	auo.mutation.SetAgentID(s)
	return auo
}

// SetNillableAgentID sets the "AgentID" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableAgentID(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetAgentID(*s)
	}
	return auo
}

// ClearAgentID clears the value of the "AgentID" field.
func (auo *AgentkycUpdateOne) ClearAgentID() *AgentkycUpdateOne {
	auo.mutation.ClearAgentID()
	return auo
}

// SetAgentemail sets the "Agentemail" field.
func (auo *AgentkycUpdateOne) SetAgentemail(s string) *AgentkycUpdateOne {
	auo.mutation.SetAgentemail(s)
	return auo
}

// SetNillableAgentemail sets the "Agentemail" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableAgentemail(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetAgentemail(*s)
	}
	return auo
}

// ClearAgentemail clears the value of the "Agentemail" field.
func (auo *AgentkycUpdateOne) ClearAgentemail() *AgentkycUpdateOne {
	auo.mutation.ClearAgentemail()
	return auo
}

// SetAgentNameLastname sets the "AgentNameLastname" field.
func (auo *AgentkycUpdateOne) SetAgentNameLastname(s string) *AgentkycUpdateOne {
	auo.mutation.SetAgentNameLastname(s)
	return auo
}

// SetNillableAgentNameLastname sets the "AgentNameLastname" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableAgentNameLastname(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetAgentNameLastname(*s)
	}
	return auo
}

// ClearAgentNameLastname clears the value of the "AgentNameLastname" field.
func (auo *AgentkycUpdateOne) ClearAgentNameLastname() *AgentkycUpdateOne {
	auo.mutation.ClearAgentNameLastname()
	return auo
}

// SetKYCStatus sets the "KYCStatus" field.
func (auo *AgentkycUpdateOne) SetKYCStatus(s string) *AgentkycUpdateOne {
	auo.mutation.SetKYCStatus(s)
	return auo
}

// SetNillableKYCStatus sets the "KYCStatus" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableKYCStatus(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetKYCStatus(*s)
	}
	return auo
}

// ClearKYCStatus clears the value of the "KYCStatus" field.
func (auo *AgentkycUpdateOne) ClearKYCStatus() *AgentkycUpdateOne {
	auo.mutation.ClearKYCStatus()
	return auo
}

// SetConsumerwalletid sets the "Consumerwalletid" field.
func (auo *AgentkycUpdateOne) SetConsumerwalletid(s string) *AgentkycUpdateOne {
	auo.mutation.SetConsumerwalletid(s)
	return auo
}

// SetNillableConsumerwalletid sets the "Consumerwalletid" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableConsumerwalletid(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetConsumerwalletid(*s)
	}
	return auo
}

// ClearConsumerwalletid clears the value of the "Consumerwalletid" field.
func (auo *AgentkycUpdateOne) ClearConsumerwalletid() *AgentkycUpdateOne {
	auo.mutation.ClearConsumerwalletid()
	return auo
}

// SetKYCRespond sets the "KYCRespond" field.
func (auo *AgentkycUpdateOne) SetKYCRespond(s string) *AgentkycUpdateOne {
	auo.mutation.SetKYCRespond(s)
	return auo
}

// SetNillableKYCRespond sets the "KYCRespond" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableKYCRespond(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetKYCRespond(*s)
	}
	return auo
}

// ClearKYCRespond clears the value of the "KYCRespond" field.
func (auo *AgentkycUpdateOne) ClearKYCRespond() *AgentkycUpdateOne {
	auo.mutation.ClearKYCRespond()
	return auo
}

// SetDOPARespond sets the "DOPARespond" field.
func (auo *AgentkycUpdateOne) SetDOPARespond(s string) *AgentkycUpdateOne {
	auo.mutation.SetDOPARespond(s)
	return auo
}

// SetNillableDOPARespond sets the "DOPARespond" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableDOPARespond(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetDOPARespond(*s)
	}
	return auo
}

// ClearDOPARespond clears the value of the "DOPARespond" field.
func (auo *AgentkycUpdateOne) ClearDOPARespond() *AgentkycUpdateOne {
	auo.mutation.ClearDOPARespond()
	return auo
}

// SetAgentType sets the "AgentType" field.
func (auo *AgentkycUpdateOne) SetAgentType(s string) *AgentkycUpdateOne {
	auo.mutation.SetAgentType(s)
	return auo
}

// SetNillableAgentType sets the "AgentType" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableAgentType(s *string) *AgentkycUpdateOne {
	if s != nil {
		auo.SetAgentType(*s)
	}
	return auo
}

// ClearAgentType clears the value of the "AgentType" field.
func (auo *AgentkycUpdateOne) ClearAgentType() *AgentkycUpdateOne {
	auo.mutation.ClearAgentType()
	return auo
}

// SetFileimportID sets the "FileimportID" field.
func (auo *AgentkycUpdateOne) SetFileimportID(i int) *AgentkycUpdateOne {
	auo.mutation.ResetFileimportID()
	auo.mutation.SetFileimportID(i)
	return auo
}

// SetNillableFileimportID sets the "FileimportID" field if the given value is not nil.
func (auo *AgentkycUpdateOne) SetNillableFileimportID(i *int) *AgentkycUpdateOne {
	if i != nil {
		auo.SetFileimportID(*i)
	}
	return auo
}

// AddFileimportID adds i to the "FileimportID" field.
func (auo *AgentkycUpdateOne) AddFileimportID(i int) *AgentkycUpdateOne {
	auo.mutation.AddFileimportID(i)
	return auo
}

// ClearFileimportID clears the value of the "FileimportID" field.
func (auo *AgentkycUpdateOne) ClearFileimportID() *AgentkycUpdateOne {
	auo.mutation.ClearFileimportID()
	return auo
}

// Mutation returns the AgentkycMutation object of the builder.
func (auo *AgentkycUpdateOne) Mutation() *AgentkycMutation {
	return auo.mutation
}

// Save executes the query and returns the updated Agentkyc entity.
func (auo *AgentkycUpdateOne) Save(ctx context.Context) (*Agentkyc, error) {
	var (
		err  error
		node *Agentkyc
	)
	if len(auo.hooks) == 0 {
		if err = auo.check(); err != nil {
			return nil, err
		}
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AgentkycMutation)
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
func (auo *AgentkycUpdateOne) SaveX(ctx context.Context) *Agentkyc {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AgentkycUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AgentkycUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AgentkycUpdateOne) check() error {
	if v, ok := auo.mutation.KYCDate(); ok {
		if err := agentkyc.KYCDateValidator(v); err != nil {
			return &ValidationError{Name: "KYCDate", err: fmt.Errorf("ent: validator failed for field \"KYCDate\": %w", err)}
		}
	}
	if v, ok := auo.mutation.KYCTime(); ok {
		if err := agentkyc.KYCTimeValidator(v); err != nil {
			return &ValidationError{Name: "KYCTime", err: fmt.Errorf("ent: validator failed for field \"KYCTime\": %w", err)}
		}
	}
	if v, ok := auo.mutation.AgentID(); ok {
		if err := agentkyc.AgentIDValidator(v); err != nil {
			return &ValidationError{Name: "AgentID", err: fmt.Errorf("ent: validator failed for field \"AgentID\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Agentemail(); ok {
		if err := agentkyc.AgentemailValidator(v); err != nil {
			return &ValidationError{Name: "Agentemail", err: fmt.Errorf("ent: validator failed for field \"Agentemail\": %w", err)}
		}
	}
	if v, ok := auo.mutation.AgentNameLastname(); ok {
		if err := agentkyc.AgentNameLastnameValidator(v); err != nil {
			return &ValidationError{Name: "AgentNameLastname", err: fmt.Errorf("ent: validator failed for field \"AgentNameLastname\": %w", err)}
		}
	}
	if v, ok := auo.mutation.KYCStatus(); ok {
		if err := agentkyc.KYCStatusValidator(v); err != nil {
			return &ValidationError{Name: "KYCStatus", err: fmt.Errorf("ent: validator failed for field \"KYCStatus\": %w", err)}
		}
	}
	if v, ok := auo.mutation.Consumerwalletid(); ok {
		if err := agentkyc.ConsumerwalletidValidator(v); err != nil {
			return &ValidationError{Name: "Consumerwalletid", err: fmt.Errorf("ent: validator failed for field \"Consumerwalletid\": %w", err)}
		}
	}
	if v, ok := auo.mutation.KYCRespond(); ok {
		if err := agentkyc.KYCRespondValidator(v); err != nil {
			return &ValidationError{Name: "KYCRespond", err: fmt.Errorf("ent: validator failed for field \"KYCRespond\": %w", err)}
		}
	}
	if v, ok := auo.mutation.DOPARespond(); ok {
		if err := agentkyc.DOPARespondValidator(v); err != nil {
			return &ValidationError{Name: "DOPARespond", err: fmt.Errorf("ent: validator failed for field \"DOPARespond\": %w", err)}
		}
	}
	if v, ok := auo.mutation.AgentType(); ok {
		if err := agentkyc.AgentTypeValidator(v); err != nil {
			return &ValidationError{Name: "AgentType", err: fmt.Errorf("ent: validator failed for field \"AgentType\": %w", err)}
		}
	}
	return nil
}

func (auo *AgentkycUpdateOne) sqlSave(ctx context.Context) (_node *Agentkyc, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   agentkyc.Table,
			Columns: agentkyc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: agentkyc.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Agentkyc.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.KYCDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCDate,
		})
	}
	if auo.mutation.KYCDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCDate,
		})
	}
	if value, ok := auo.mutation.KYCTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCTime,
		})
	}
	if auo.mutation.KYCTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCTime,
		})
	}
	if value, ok := auo.mutation.AgentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentID,
		})
	}
	if auo.mutation.AgentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentID,
		})
	}
	if value, ok := auo.mutation.Agentemail(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentemail,
		})
	}
	if auo.mutation.AgentemailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentemail,
		})
	}
	if value, ok := auo.mutation.AgentNameLastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentNameLastname,
		})
	}
	if auo.mutation.AgentNameLastnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentNameLastname,
		})
	}
	if value, ok := auo.mutation.KYCStatus(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCStatus,
		})
	}
	if auo.mutation.KYCStatusCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCStatus,
		})
	}
	if value, ok := auo.mutation.Consumerwalletid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldConsumerwalletid,
		})
	}
	if auo.mutation.ConsumerwalletidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldConsumerwalletid,
		})
	}
	if value, ok := auo.mutation.KYCRespond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldKYCRespond,
		})
	}
	if auo.mutation.KYCRespondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldKYCRespond,
		})
	}
	if value, ok := auo.mutation.DOPARespond(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldDOPARespond,
		})
	}
	if auo.mutation.DOPARespondCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldDOPARespond,
		})
	}
	if value, ok := auo.mutation.AgentType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: agentkyc.FieldAgentType,
		})
	}
	if auo.mutation.AgentTypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: agentkyc.FieldAgentType,
		})
	}
	if value, ok := auo.mutation.FileimportID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agentkyc.FieldFileimportID,
		})
	}
	if value, ok := auo.mutation.AddedFileimportID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: agentkyc.FieldFileimportID,
		})
	}
	if auo.mutation.FileimportIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: agentkyc.FieldFileimportID,
		})
	}
	_node = &Agentkyc{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{agentkyc.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
