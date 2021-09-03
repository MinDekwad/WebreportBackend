// Code generated by entc, DO NOT EDIT.

package logexport

import (
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserName applies equality check predicate on the "UserName" field. It's identical to UserNameEQ.
func UserName(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// FileName applies equality check predicate on the "FileName" field. It's identical to FileNameEQ.
func FileName(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileName), v))
	})
}

// ExportDate applies equality check predicate on the "ExportDate" field. It's identical to ExportDateEQ.
func ExportDate(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExportDate), v))
	})
}

// UserNameEQ applies the EQ predicate on the "UserName" field.
func UserNameEQ(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserName), v))
	})
}

// UserNameNEQ applies the NEQ predicate on the "UserName" field.
func UserNameNEQ(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserName), v))
	})
}

// UserNameIn applies the In predicate on the "UserName" field.
func UserNameIn(vs ...string) predicate.Logexport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserName), v...))
	})
}

// UserNameNotIn applies the NotIn predicate on the "UserName" field.
func UserNameNotIn(vs ...string) predicate.Logexport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserName), v...))
	})
}

// UserNameGT applies the GT predicate on the "UserName" field.
func UserNameGT(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserName), v))
	})
}

// UserNameGTE applies the GTE predicate on the "UserName" field.
func UserNameGTE(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserName), v))
	})
}

// UserNameLT applies the LT predicate on the "UserName" field.
func UserNameLT(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserName), v))
	})
}

// UserNameLTE applies the LTE predicate on the "UserName" field.
func UserNameLTE(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserName), v))
	})
}

// UserNameContains applies the Contains predicate on the "UserName" field.
func UserNameContains(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserName), v))
	})
}

// UserNameHasPrefix applies the HasPrefix predicate on the "UserName" field.
func UserNameHasPrefix(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserName), v))
	})
}

// UserNameHasSuffix applies the HasSuffix predicate on the "UserName" field.
func UserNameHasSuffix(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserName), v))
	})
}

// UserNameIsNil applies the IsNil predicate on the "UserName" field.
func UserNameIsNil() predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserName)))
	})
}

// UserNameNotNil applies the NotNil predicate on the "UserName" field.
func UserNameNotNil() predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserName)))
	})
}

// UserNameEqualFold applies the EqualFold predicate on the "UserName" field.
func UserNameEqualFold(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserName), v))
	})
}

// UserNameContainsFold applies the ContainsFold predicate on the "UserName" field.
func UserNameContainsFold(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserName), v))
	})
}

// FileNameEQ applies the EQ predicate on the "FileName" field.
func FileNameEQ(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileName), v))
	})
}

// FileNameNEQ applies the NEQ predicate on the "FileName" field.
func FileNameNEQ(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFileName), v))
	})
}

// FileNameIn applies the In predicate on the "FileName" field.
func FileNameIn(vs ...string) predicate.Logexport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFileName), v...))
	})
}

// FileNameNotIn applies the NotIn predicate on the "FileName" field.
func FileNameNotIn(vs ...string) predicate.Logexport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFileName), v...))
	})
}

// FileNameGT applies the GT predicate on the "FileName" field.
func FileNameGT(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFileName), v))
	})
}

// FileNameGTE applies the GTE predicate on the "FileName" field.
func FileNameGTE(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFileName), v))
	})
}

// FileNameLT applies the LT predicate on the "FileName" field.
func FileNameLT(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFileName), v))
	})
}

// FileNameLTE applies the LTE predicate on the "FileName" field.
func FileNameLTE(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFileName), v))
	})
}

// FileNameContains applies the Contains predicate on the "FileName" field.
func FileNameContains(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFileName), v))
	})
}

// FileNameHasPrefix applies the HasPrefix predicate on the "FileName" field.
func FileNameHasPrefix(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFileName), v))
	})
}

// FileNameHasSuffix applies the HasSuffix predicate on the "FileName" field.
func FileNameHasSuffix(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFileName), v))
	})
}

// FileNameIsNil applies the IsNil predicate on the "FileName" field.
func FileNameIsNil() predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFileName)))
	})
}

// FileNameNotNil applies the NotNil predicate on the "FileName" field.
func FileNameNotNil() predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFileName)))
	})
}

// FileNameEqualFold applies the EqualFold predicate on the "FileName" field.
func FileNameEqualFold(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFileName), v))
	})
}

// FileNameContainsFold applies the ContainsFold predicate on the "FileName" field.
func FileNameContainsFold(v string) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFileName), v))
	})
}

// ExportDateEQ applies the EQ predicate on the "ExportDate" field.
func ExportDateEQ(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldExportDate), v))
	})
}

// ExportDateNEQ applies the NEQ predicate on the "ExportDate" field.
func ExportDateNEQ(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldExportDate), v))
	})
}

// ExportDateIn applies the In predicate on the "ExportDate" field.
func ExportDateIn(vs ...time.Time) predicate.Logexport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldExportDate), v...))
	})
}

// ExportDateNotIn applies the NotIn predicate on the "ExportDate" field.
func ExportDateNotIn(vs ...time.Time) predicate.Logexport {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Logexport(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldExportDate), v...))
	})
}

// ExportDateGT applies the GT predicate on the "ExportDate" field.
func ExportDateGT(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldExportDate), v))
	})
}

// ExportDateGTE applies the GTE predicate on the "ExportDate" field.
func ExportDateGTE(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldExportDate), v))
	})
}

// ExportDateLT applies the LT predicate on the "ExportDate" field.
func ExportDateLT(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldExportDate), v))
	})
}

// ExportDateLTE applies the LTE predicate on the "ExportDate" field.
func ExportDateLTE(v time.Time) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldExportDate), v))
	})
}

// ExportDateIsNil applies the IsNil predicate on the "ExportDate" field.
func ExportDateIsNil() predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldExportDate)))
	})
}

// ExportDateNotNil applies the NotNil predicate on the "ExportDate" field.
func ExportDateNotNil() predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldExportDate)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Logexport) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Logexport) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Logexport) predicate.Logexport {
	return predicate.Logexport(func(s *sql.Selector) {
		p(s.Not())
	})
}
