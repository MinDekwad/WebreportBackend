// Code generated by entc, DO NOT EDIT.

package occupationhistory

import (
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// WalletID applies equality check predicate on the "WalletID" field. It's identical to WalletIDEQ.
func WalletID(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletID), v))
	})
}

// OccupationName applies equality check predicate on the "OccupationName" field. It's identical to OccupationNameEQ.
func OccupationName(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccupationName), v))
	})
}

// RankOccupation applies equality check predicate on the "RankOccupation" field. It's identical to RankOccupationEQ.
func RankOccupation(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRankOccupation), v))
	})
}

// DateCalRank applies equality check predicate on the "DateCalRank" field. It's identical to DateCalRankEQ.
func DateCalRank(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCalRank), v))
	})
}

// WalletIDEQ applies the EQ predicate on the "WalletID" field.
func WalletIDEQ(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletID), v))
	})
}

// WalletIDNEQ applies the NEQ predicate on the "WalletID" field.
func WalletIDNEQ(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWalletID), v))
	})
}

// WalletIDIn applies the In predicate on the "WalletID" field.
func WalletIDIn(vs ...string) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWalletID), v...))
	})
}

// WalletIDNotIn applies the NotIn predicate on the "WalletID" field.
func WalletIDNotIn(vs ...string) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWalletID), v...))
	})
}

// WalletIDGT applies the GT predicate on the "WalletID" field.
func WalletIDGT(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWalletID), v))
	})
}

// WalletIDGTE applies the GTE predicate on the "WalletID" field.
func WalletIDGTE(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWalletID), v))
	})
}

// WalletIDLT applies the LT predicate on the "WalletID" field.
func WalletIDLT(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWalletID), v))
	})
}

// WalletIDLTE applies the LTE predicate on the "WalletID" field.
func WalletIDLTE(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWalletID), v))
	})
}

// WalletIDContains applies the Contains predicate on the "WalletID" field.
func WalletIDContains(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWalletID), v))
	})
}

// WalletIDHasPrefix applies the HasPrefix predicate on the "WalletID" field.
func WalletIDHasPrefix(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWalletID), v))
	})
}

// WalletIDHasSuffix applies the HasSuffix predicate on the "WalletID" field.
func WalletIDHasSuffix(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWalletID), v))
	})
}

// WalletIDEqualFold applies the EqualFold predicate on the "WalletID" field.
func WalletIDEqualFold(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWalletID), v))
	})
}

// WalletIDContainsFold applies the ContainsFold predicate on the "WalletID" field.
func WalletIDContainsFold(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWalletID), v))
	})
}

// OccupationNameEQ applies the EQ predicate on the "OccupationName" field.
func OccupationNameEQ(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccupationName), v))
	})
}

// OccupationNameNEQ applies the NEQ predicate on the "OccupationName" field.
func OccupationNameNEQ(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOccupationName), v))
	})
}

// OccupationNameIn applies the In predicate on the "OccupationName" field.
func OccupationNameIn(vs ...string) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldOccupationName), v...))
	})
}

// OccupationNameNotIn applies the NotIn predicate on the "OccupationName" field.
func OccupationNameNotIn(vs ...string) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldOccupationName), v...))
	})
}

// OccupationNameGT applies the GT predicate on the "OccupationName" field.
func OccupationNameGT(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOccupationName), v))
	})
}

// OccupationNameGTE applies the GTE predicate on the "OccupationName" field.
func OccupationNameGTE(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOccupationName), v))
	})
}

// OccupationNameLT applies the LT predicate on the "OccupationName" field.
func OccupationNameLT(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOccupationName), v))
	})
}

// OccupationNameLTE applies the LTE predicate on the "OccupationName" field.
func OccupationNameLTE(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOccupationName), v))
	})
}

// OccupationNameContains applies the Contains predicate on the "OccupationName" field.
func OccupationNameContains(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOccupationName), v))
	})
}

// OccupationNameHasPrefix applies the HasPrefix predicate on the "OccupationName" field.
func OccupationNameHasPrefix(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOccupationName), v))
	})
}

// OccupationNameHasSuffix applies the HasSuffix predicate on the "OccupationName" field.
func OccupationNameHasSuffix(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOccupationName), v))
	})
}

// OccupationNameIsNil applies the IsNil predicate on the "OccupationName" field.
func OccupationNameIsNil() predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOccupationName)))
	})
}

// OccupationNameNotNil applies the NotNil predicate on the "OccupationName" field.
func OccupationNameNotNil() predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOccupationName)))
	})
}

// OccupationNameEqualFold applies the EqualFold predicate on the "OccupationName" field.
func OccupationNameEqualFold(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOccupationName), v))
	})
}

// OccupationNameContainsFold applies the ContainsFold predicate on the "OccupationName" field.
func OccupationNameContainsFold(v string) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOccupationName), v))
	})
}

// RankOccupationEQ applies the EQ predicate on the "RankOccupation" field.
func RankOccupationEQ(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRankOccupation), v))
	})
}

// RankOccupationNEQ applies the NEQ predicate on the "RankOccupation" field.
func RankOccupationNEQ(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRankOccupation), v))
	})
}

// RankOccupationIn applies the In predicate on the "RankOccupation" field.
func RankOccupationIn(vs ...int) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRankOccupation), v...))
	})
}

// RankOccupationNotIn applies the NotIn predicate on the "RankOccupation" field.
func RankOccupationNotIn(vs ...int) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRankOccupation), v...))
	})
}

// RankOccupationGT applies the GT predicate on the "RankOccupation" field.
func RankOccupationGT(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRankOccupation), v))
	})
}

// RankOccupationGTE applies the GTE predicate on the "RankOccupation" field.
func RankOccupationGTE(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRankOccupation), v))
	})
}

// RankOccupationLT applies the LT predicate on the "RankOccupation" field.
func RankOccupationLT(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRankOccupation), v))
	})
}

// RankOccupationLTE applies the LTE predicate on the "RankOccupation" field.
func RankOccupationLTE(v int) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRankOccupation), v))
	})
}

// RankOccupationIsNil applies the IsNil predicate on the "RankOccupation" field.
func RankOccupationIsNil() predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRankOccupation)))
	})
}

// RankOccupationNotNil applies the NotNil predicate on the "RankOccupation" field.
func RankOccupationNotNil() predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRankOccupation)))
	})
}

// DateCalRankEQ applies the EQ predicate on the "DateCalRank" field.
func DateCalRankEQ(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankNEQ applies the NEQ predicate on the "DateCalRank" field.
func DateCalRankNEQ(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankIn applies the In predicate on the "DateCalRank" field.
func DateCalRankIn(vs ...time.Time) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateCalRank), v...))
	})
}

// DateCalRankNotIn applies the NotIn predicate on the "DateCalRank" field.
func DateCalRankNotIn(vs ...time.Time) predicate.Occupationhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Occupationhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateCalRank), v...))
	})
}

// DateCalRankGT applies the GT predicate on the "DateCalRank" field.
func DateCalRankGT(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankGTE applies the GTE predicate on the "DateCalRank" field.
func DateCalRankGTE(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankLT applies the LT predicate on the "DateCalRank" field.
func DateCalRankLT(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankLTE applies the LTE predicate on the "DateCalRank" field.
func DateCalRankLTE(v time.Time) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankIsNil applies the IsNil predicate on the "DateCalRank" field.
func DateCalRankIsNil() predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDateCalRank)))
	})
}

// DateCalRankNotNil applies the NotNil predicate on the "DateCalRank" field.
func DateCalRankNotNil() predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDateCalRank)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Occupationhistory) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Occupationhistory) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
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
func Not(p predicate.Occupationhistory) predicate.Occupationhistory {
	return predicate.Occupationhistory(func(s *sql.Selector) {
		p(s.Not())
	})
}