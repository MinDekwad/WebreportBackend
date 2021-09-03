// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-api-report2/ent/predicate"
	"go-api-report2/ent/userprofile"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserprofileUpdate is the builder for updating Userprofile entities.
type UserprofileUpdate struct {
	config
	hooks    []Hook
	mutation *UserprofileMutation
}

// Where adds a new predicate for the UserprofileUpdate builder.
func (uu *UserprofileUpdate) Where(ps ...predicate.Userprofile) *UserprofileUpdate {
	uu.mutation.predicates = append(uu.mutation.predicates, ps...)
	return uu
}

// SetUserId sets the "UserId" field.
func (uu *UserprofileUpdate) SetUserId(s string) *UserprofileUpdate {
	uu.mutation.SetUserId(s)
	return uu
}

// SetNillableUserId sets the "UserId" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableUserId(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetUserId(*s)
	}
	return uu
}

// ClearUserId clears the value of the "UserId" field.
func (uu *UserprofileUpdate) ClearUserId() *UserprofileUpdate {
	uu.mutation.ClearUserId()
	return uu
}

// SetFirstname sets the "Firstname" field.
func (uu *UserprofileUpdate) SetFirstname(s string) *UserprofileUpdate {
	uu.mutation.SetFirstname(s)
	return uu
}

// SetNillableFirstname sets the "Firstname" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableFirstname(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetFirstname(*s)
	}
	return uu
}

// ClearFirstname clears the value of the "Firstname" field.
func (uu *UserprofileUpdate) ClearFirstname() *UserprofileUpdate {
	uu.mutation.ClearFirstname()
	return uu
}

// SetLastname sets the "Lastname" field.
func (uu *UserprofileUpdate) SetLastname(s string) *UserprofileUpdate {
	uu.mutation.SetLastname(s)
	return uu
}

// SetNillableLastname sets the "Lastname" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableLastname(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetLastname(*s)
	}
	return uu
}

// ClearLastname clears the value of the "Lastname" field.
func (uu *UserprofileUpdate) ClearLastname() *UserprofileUpdate {
	uu.mutation.ClearLastname()
	return uu
}

// SetPhoneNo sets the "PhoneNo" field.
func (uu *UserprofileUpdate) SetPhoneNo(s string) *UserprofileUpdate {
	uu.mutation.SetPhoneNo(s)
	return uu
}

// SetNillablePhoneNo sets the "PhoneNo" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillablePhoneNo(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetPhoneNo(*s)
	}
	return uu
}

// ClearPhoneNo clears the value of the "PhoneNo" field.
func (uu *UserprofileUpdate) ClearPhoneNo() *UserprofileUpdate {
	uu.mutation.ClearPhoneNo()
	return uu
}

// SetEmail sets the "Email" field.
func (uu *UserprofileUpdate) SetEmail(s string) *UserprofileUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableEmail(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "Email" field.
func (uu *UserprofileUpdate) ClearEmail() *UserprofileUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetCitizenId sets the "CitizenId" field.
func (uu *UserprofileUpdate) SetCitizenId(s string) *UserprofileUpdate {
	uu.mutation.SetCitizenId(s)
	return uu
}

// SetNillableCitizenId sets the "CitizenId" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableCitizenId(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetCitizenId(*s)
	}
	return uu
}

// ClearCitizenId clears the value of the "CitizenId" field.
func (uu *UserprofileUpdate) ClearCitizenId() *UserprofileUpdate {
	uu.mutation.ClearCitizenId()
	return uu
}

// SetBirthDate sets the "BirthDate" field.
func (uu *UserprofileUpdate) SetBirthDate(s string) *UserprofileUpdate {
	uu.mutation.SetBirthDate(s)
	return uu
}

// SetNillableBirthDate sets the "BirthDate" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableBirthDate(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetBirthDate(*s)
	}
	return uu
}

// ClearBirthDate clears the value of the "BirthDate" field.
func (uu *UserprofileUpdate) ClearBirthDate() *UserprofileUpdate {
	uu.mutation.ClearBirthDate()
	return uu
}

// SetGender sets the "Gender" field.
func (uu *UserprofileUpdate) SetGender(s string) *UserprofileUpdate {
	uu.mutation.SetGender(s)
	return uu
}

// SetNillableGender sets the "Gender" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableGender(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetGender(*s)
	}
	return uu
}

// ClearGender clears the value of the "Gender" field.
func (uu *UserprofileUpdate) ClearGender() *UserprofileUpdate {
	uu.mutation.ClearGender()
	return uu
}

// SetBusinessAddress sets the "BusinessAddress" field.
func (uu *UserprofileUpdate) SetBusinessAddress(s string) *UserprofileUpdate {
	uu.mutation.SetBusinessAddress(s)
	return uu
}

// SetNillableBusinessAddress sets the "BusinessAddress" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableBusinessAddress(s *string) *UserprofileUpdate {
	if s != nil {
		uu.SetBusinessAddress(*s)
	}
	return uu
}

// ClearBusinessAddress clears the value of the "BusinessAddress" field.
func (uu *UserprofileUpdate) ClearBusinessAddress() *UserprofileUpdate {
	uu.mutation.ClearBusinessAddress()
	return uu
}

// SetOccupationId sets the "OccupationId" field.
func (uu *UserprofileUpdate) SetOccupationId(i int) *UserprofileUpdate {
	uu.mutation.ResetOccupationId()
	uu.mutation.SetOccupationId(i)
	return uu
}

// SetNillableOccupationId sets the "OccupationId" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableOccupationId(i *int) *UserprofileUpdate {
	if i != nil {
		uu.SetOccupationId(*i)
	}
	return uu
}

// AddOccupationId adds i to the "OccupationId" field.
func (uu *UserprofileUpdate) AddOccupationId(i int) *UserprofileUpdate {
	uu.mutation.AddOccupationId(i)
	return uu
}

// ClearOccupationId clears the value of the "OccupationId" field.
func (uu *UserprofileUpdate) ClearOccupationId() *UserprofileUpdate {
	uu.mutation.ClearOccupationId()
	return uu
}

// SetFileimportID sets the "FileimportID" field.
func (uu *UserprofileUpdate) SetFileimportID(i int) *UserprofileUpdate {
	uu.mutation.ResetFileimportID()
	uu.mutation.SetFileimportID(i)
	return uu
}

// SetNillableFileimportID sets the "FileimportID" field if the given value is not nil.
func (uu *UserprofileUpdate) SetNillableFileimportID(i *int) *UserprofileUpdate {
	if i != nil {
		uu.SetFileimportID(*i)
	}
	return uu
}

// AddFileimportID adds i to the "FileimportID" field.
func (uu *UserprofileUpdate) AddFileimportID(i int) *UserprofileUpdate {
	uu.mutation.AddFileimportID(i)
	return uu
}

// ClearFileimportID clears the value of the "FileimportID" field.
func (uu *UserprofileUpdate) ClearFileimportID() *UserprofileUpdate {
	uu.mutation.ClearFileimportID()
	return uu
}

// Mutation returns the UserprofileMutation object of the builder.
func (uu *UserprofileUpdate) Mutation() *UserprofileMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserprofileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		if err = uu.check(); err != nil {
			return 0, err
		}
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserprofileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uu.check(); err != nil {
				return 0, err
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserprofileUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserprofileUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserprofileUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserprofileUpdate) check() error {
	if v, ok := uu.mutation.UserId(); ok {
		if err := userprofile.UserIdValidator(v); err != nil {
			return &ValidationError{Name: "UserId", err: fmt.Errorf("ent: validator failed for field \"UserId\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Firstname(); ok {
		if err := userprofile.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "Firstname", err: fmt.Errorf("ent: validator failed for field \"Firstname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Lastname(); ok {
		if err := userprofile.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "Lastname", err: fmt.Errorf("ent: validator failed for field \"Lastname\": %w", err)}
		}
	}
	if v, ok := uu.mutation.PhoneNo(); ok {
		if err := userprofile.PhoneNoValidator(v); err != nil {
			return &ValidationError{Name: "PhoneNo", err: fmt.Errorf("ent: validator failed for field \"PhoneNo\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Email(); ok {
		if err := userprofile.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf("ent: validator failed for field \"Email\": %w", err)}
		}
	}
	if v, ok := uu.mutation.CitizenId(); ok {
		if err := userprofile.CitizenIdValidator(v); err != nil {
			return &ValidationError{Name: "CitizenId", err: fmt.Errorf("ent: validator failed for field \"CitizenId\": %w", err)}
		}
	}
	if v, ok := uu.mutation.BirthDate(); ok {
		if err := userprofile.BirthDateValidator(v); err != nil {
			return &ValidationError{Name: "BirthDate", err: fmt.Errorf("ent: validator failed for field \"BirthDate\": %w", err)}
		}
	}
	if v, ok := uu.mutation.Gender(); ok {
		if err := userprofile.GenderValidator(v); err != nil {
			return &ValidationError{Name: "Gender", err: fmt.Errorf("ent: validator failed for field \"Gender\": %w", err)}
		}
	}
	if v, ok := uu.mutation.BusinessAddress(); ok {
		if err := userprofile.BusinessAddressValidator(v); err != nil {
			return &ValidationError{Name: "BusinessAddress", err: fmt.Errorf("ent: validator failed for field \"BusinessAddress\": %w", err)}
		}
	}
	return nil
}

func (uu *UserprofileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userprofile.Table,
			Columns: userprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userprofile.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UserId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldUserId,
		})
	}
	if uu.mutation.UserIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldUserId,
		})
	}
	if value, ok := uu.mutation.Firstname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldFirstname,
		})
	}
	if uu.mutation.FirstnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldFirstname,
		})
	}
	if value, ok := uu.mutation.Lastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldLastname,
		})
	}
	if uu.mutation.LastnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldLastname,
		})
	}
	if value, ok := uu.mutation.PhoneNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldPhoneNo,
		})
	}
	if uu.mutation.PhoneNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldPhoneNo,
		})
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldEmail,
		})
	}
	if uu.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldEmail,
		})
	}
	if value, ok := uu.mutation.CitizenId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldCitizenId,
		})
	}
	if uu.mutation.CitizenIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldCitizenId,
		})
	}
	if value, ok := uu.mutation.BirthDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldBirthDate,
		})
	}
	if uu.mutation.BirthDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldBirthDate,
		})
	}
	if value, ok := uu.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldGender,
		})
	}
	if uu.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldGender,
		})
	}
	if value, ok := uu.mutation.BusinessAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldBusinessAddress,
		})
	}
	if uu.mutation.BusinessAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldBusinessAddress,
		})
	}
	if value, ok := uu.mutation.OccupationId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldOccupationId,
		})
	}
	if value, ok := uu.mutation.AddedOccupationId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldOccupationId,
		})
	}
	if uu.mutation.OccupationIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: userprofile.FieldOccupationId,
		})
	}
	if value, ok := uu.mutation.FileimportID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldFileimportID,
		})
	}
	if value, ok := uu.mutation.AddedFileimportID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldFileimportID,
		})
	}
	if uu.mutation.FileimportIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: userprofile.FieldFileimportID,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserprofileUpdateOne is the builder for updating a single Userprofile entity.
type UserprofileUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserprofileMutation
}

// SetUserId sets the "UserId" field.
func (uuo *UserprofileUpdateOne) SetUserId(s string) *UserprofileUpdateOne {
	uuo.mutation.SetUserId(s)
	return uuo
}

// SetNillableUserId sets the "UserId" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableUserId(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetUserId(*s)
	}
	return uuo
}

// ClearUserId clears the value of the "UserId" field.
func (uuo *UserprofileUpdateOne) ClearUserId() *UserprofileUpdateOne {
	uuo.mutation.ClearUserId()
	return uuo
}

// SetFirstname sets the "Firstname" field.
func (uuo *UserprofileUpdateOne) SetFirstname(s string) *UserprofileUpdateOne {
	uuo.mutation.SetFirstname(s)
	return uuo
}

// SetNillableFirstname sets the "Firstname" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableFirstname(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetFirstname(*s)
	}
	return uuo
}

// ClearFirstname clears the value of the "Firstname" field.
func (uuo *UserprofileUpdateOne) ClearFirstname() *UserprofileUpdateOne {
	uuo.mutation.ClearFirstname()
	return uuo
}

// SetLastname sets the "Lastname" field.
func (uuo *UserprofileUpdateOne) SetLastname(s string) *UserprofileUpdateOne {
	uuo.mutation.SetLastname(s)
	return uuo
}

// SetNillableLastname sets the "Lastname" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableLastname(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetLastname(*s)
	}
	return uuo
}

// ClearLastname clears the value of the "Lastname" field.
func (uuo *UserprofileUpdateOne) ClearLastname() *UserprofileUpdateOne {
	uuo.mutation.ClearLastname()
	return uuo
}

// SetPhoneNo sets the "PhoneNo" field.
func (uuo *UserprofileUpdateOne) SetPhoneNo(s string) *UserprofileUpdateOne {
	uuo.mutation.SetPhoneNo(s)
	return uuo
}

// SetNillablePhoneNo sets the "PhoneNo" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillablePhoneNo(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetPhoneNo(*s)
	}
	return uuo
}

// ClearPhoneNo clears the value of the "PhoneNo" field.
func (uuo *UserprofileUpdateOne) ClearPhoneNo() *UserprofileUpdateOne {
	uuo.mutation.ClearPhoneNo()
	return uuo
}

// SetEmail sets the "Email" field.
func (uuo *UserprofileUpdateOne) SetEmail(s string) *UserprofileUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "Email" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableEmail(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "Email" field.
func (uuo *UserprofileUpdateOne) ClearEmail() *UserprofileUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetCitizenId sets the "CitizenId" field.
func (uuo *UserprofileUpdateOne) SetCitizenId(s string) *UserprofileUpdateOne {
	uuo.mutation.SetCitizenId(s)
	return uuo
}

// SetNillableCitizenId sets the "CitizenId" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableCitizenId(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetCitizenId(*s)
	}
	return uuo
}

// ClearCitizenId clears the value of the "CitizenId" field.
func (uuo *UserprofileUpdateOne) ClearCitizenId() *UserprofileUpdateOne {
	uuo.mutation.ClearCitizenId()
	return uuo
}

// SetBirthDate sets the "BirthDate" field.
func (uuo *UserprofileUpdateOne) SetBirthDate(s string) *UserprofileUpdateOne {
	uuo.mutation.SetBirthDate(s)
	return uuo
}

// SetNillableBirthDate sets the "BirthDate" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableBirthDate(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetBirthDate(*s)
	}
	return uuo
}

// ClearBirthDate clears the value of the "BirthDate" field.
func (uuo *UserprofileUpdateOne) ClearBirthDate() *UserprofileUpdateOne {
	uuo.mutation.ClearBirthDate()
	return uuo
}

// SetGender sets the "Gender" field.
func (uuo *UserprofileUpdateOne) SetGender(s string) *UserprofileUpdateOne {
	uuo.mutation.SetGender(s)
	return uuo
}

// SetNillableGender sets the "Gender" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableGender(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetGender(*s)
	}
	return uuo
}

// ClearGender clears the value of the "Gender" field.
func (uuo *UserprofileUpdateOne) ClearGender() *UserprofileUpdateOne {
	uuo.mutation.ClearGender()
	return uuo
}

// SetBusinessAddress sets the "BusinessAddress" field.
func (uuo *UserprofileUpdateOne) SetBusinessAddress(s string) *UserprofileUpdateOne {
	uuo.mutation.SetBusinessAddress(s)
	return uuo
}

// SetNillableBusinessAddress sets the "BusinessAddress" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableBusinessAddress(s *string) *UserprofileUpdateOne {
	if s != nil {
		uuo.SetBusinessAddress(*s)
	}
	return uuo
}

// ClearBusinessAddress clears the value of the "BusinessAddress" field.
func (uuo *UserprofileUpdateOne) ClearBusinessAddress() *UserprofileUpdateOne {
	uuo.mutation.ClearBusinessAddress()
	return uuo
}

// SetOccupationId sets the "OccupationId" field.
func (uuo *UserprofileUpdateOne) SetOccupationId(i int) *UserprofileUpdateOne {
	uuo.mutation.ResetOccupationId()
	uuo.mutation.SetOccupationId(i)
	return uuo
}

// SetNillableOccupationId sets the "OccupationId" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableOccupationId(i *int) *UserprofileUpdateOne {
	if i != nil {
		uuo.SetOccupationId(*i)
	}
	return uuo
}

// AddOccupationId adds i to the "OccupationId" field.
func (uuo *UserprofileUpdateOne) AddOccupationId(i int) *UserprofileUpdateOne {
	uuo.mutation.AddOccupationId(i)
	return uuo
}

// ClearOccupationId clears the value of the "OccupationId" field.
func (uuo *UserprofileUpdateOne) ClearOccupationId() *UserprofileUpdateOne {
	uuo.mutation.ClearOccupationId()
	return uuo
}

// SetFileimportID sets the "FileimportID" field.
func (uuo *UserprofileUpdateOne) SetFileimportID(i int) *UserprofileUpdateOne {
	uuo.mutation.ResetFileimportID()
	uuo.mutation.SetFileimportID(i)
	return uuo
}

// SetNillableFileimportID sets the "FileimportID" field if the given value is not nil.
func (uuo *UserprofileUpdateOne) SetNillableFileimportID(i *int) *UserprofileUpdateOne {
	if i != nil {
		uuo.SetFileimportID(*i)
	}
	return uuo
}

// AddFileimportID adds i to the "FileimportID" field.
func (uuo *UserprofileUpdateOne) AddFileimportID(i int) *UserprofileUpdateOne {
	uuo.mutation.AddFileimportID(i)
	return uuo
}

// ClearFileimportID clears the value of the "FileimportID" field.
func (uuo *UserprofileUpdateOne) ClearFileimportID() *UserprofileUpdateOne {
	uuo.mutation.ClearFileimportID()
	return uuo
}

// Mutation returns the UserprofileMutation object of the builder.
func (uuo *UserprofileUpdateOne) Mutation() *UserprofileMutation {
	return uuo.mutation
}

// Save executes the query and returns the updated Userprofile entity.
func (uuo *UserprofileUpdateOne) Save(ctx context.Context) (*Userprofile, error) {
	var (
		err  error
		node *Userprofile
	)
	if len(uuo.hooks) == 0 {
		if err = uuo.check(); err != nil {
			return nil, err
		}
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserprofileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uuo.check(); err != nil {
				return nil, err
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			mut = uuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserprofileUpdateOne) SaveX(ctx context.Context) *Userprofile {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserprofileUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserprofileUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserprofileUpdateOne) check() error {
	if v, ok := uuo.mutation.UserId(); ok {
		if err := userprofile.UserIdValidator(v); err != nil {
			return &ValidationError{Name: "UserId", err: fmt.Errorf("ent: validator failed for field \"UserId\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Firstname(); ok {
		if err := userprofile.FirstnameValidator(v); err != nil {
			return &ValidationError{Name: "Firstname", err: fmt.Errorf("ent: validator failed for field \"Firstname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Lastname(); ok {
		if err := userprofile.LastnameValidator(v); err != nil {
			return &ValidationError{Name: "Lastname", err: fmt.Errorf("ent: validator failed for field \"Lastname\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.PhoneNo(); ok {
		if err := userprofile.PhoneNoValidator(v); err != nil {
			return &ValidationError{Name: "PhoneNo", err: fmt.Errorf("ent: validator failed for field \"PhoneNo\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Email(); ok {
		if err := userprofile.EmailValidator(v); err != nil {
			return &ValidationError{Name: "Email", err: fmt.Errorf("ent: validator failed for field \"Email\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.CitizenId(); ok {
		if err := userprofile.CitizenIdValidator(v); err != nil {
			return &ValidationError{Name: "CitizenId", err: fmt.Errorf("ent: validator failed for field \"CitizenId\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.BirthDate(); ok {
		if err := userprofile.BirthDateValidator(v); err != nil {
			return &ValidationError{Name: "BirthDate", err: fmt.Errorf("ent: validator failed for field \"BirthDate\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.Gender(); ok {
		if err := userprofile.GenderValidator(v); err != nil {
			return &ValidationError{Name: "Gender", err: fmt.Errorf("ent: validator failed for field \"Gender\": %w", err)}
		}
	}
	if v, ok := uuo.mutation.BusinessAddress(); ok {
		if err := userprofile.BusinessAddressValidator(v); err != nil {
			return &ValidationError{Name: "BusinessAddress", err: fmt.Errorf("ent: validator failed for field \"BusinessAddress\": %w", err)}
		}
	}
	return nil
}

func (uuo *UserprofileUpdateOne) sqlSave(ctx context.Context) (_node *Userprofile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userprofile.Table,
			Columns: userprofile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userprofile.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Userprofile.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UserId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldUserId,
		})
	}
	if uuo.mutation.UserIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldUserId,
		})
	}
	if value, ok := uuo.mutation.Firstname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldFirstname,
		})
	}
	if uuo.mutation.FirstnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldFirstname,
		})
	}
	if value, ok := uuo.mutation.Lastname(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldLastname,
		})
	}
	if uuo.mutation.LastnameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldLastname,
		})
	}
	if value, ok := uuo.mutation.PhoneNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldPhoneNo,
		})
	}
	if uuo.mutation.PhoneNoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldPhoneNo,
		})
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldEmail,
		})
	}
	if uuo.mutation.EmailCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldEmail,
		})
	}
	if value, ok := uuo.mutation.CitizenId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldCitizenId,
		})
	}
	if uuo.mutation.CitizenIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldCitizenId,
		})
	}
	if value, ok := uuo.mutation.BirthDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldBirthDate,
		})
	}
	if uuo.mutation.BirthDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldBirthDate,
		})
	}
	if value, ok := uuo.mutation.Gender(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldGender,
		})
	}
	if uuo.mutation.GenderCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldGender,
		})
	}
	if value, ok := uuo.mutation.BusinessAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userprofile.FieldBusinessAddress,
		})
	}
	if uuo.mutation.BusinessAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: userprofile.FieldBusinessAddress,
		})
	}
	if value, ok := uuo.mutation.OccupationId(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldOccupationId,
		})
	}
	if value, ok := uuo.mutation.AddedOccupationId(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldOccupationId,
		})
	}
	if uuo.mutation.OccupationIdCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: userprofile.FieldOccupationId,
		})
	}
	if value, ok := uuo.mutation.FileimportID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldFileimportID,
		})
	}
	if value, ok := uuo.mutation.AddedFileimportID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userprofile.FieldFileimportID,
		})
	}
	if uuo.mutation.FileimportIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: userprofile.FieldFileimportID,
		})
	}
	_node = &Userprofile{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userprofile.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
