// Code generated by entc, DO NOT EDIT.

package transactionfactoritemtmp

import (
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Min applies equality check predicate on the "Min" field. It's identical to MinEQ.
func Min(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMin), v))
	})
}

// Max applies equality check predicate on the "Max" field. It's identical to MaxEQ.
func Max(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMax), v))
	})
}

// Rank applies equality check predicate on the "Rank" field. It's identical to RankEQ.
func Rank(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// MinEQ applies the EQ predicate on the "Min" field.
func MinEQ(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMin), v))
	})
}

// MinNEQ applies the NEQ predicate on the "Min" field.
func MinNEQ(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMin), v))
	})
}

// MinIn applies the In predicate on the "Min" field.
func MinIn(vs ...float64) predicate.Transactionfactoritemtmp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMin), v...))
	})
}

// MinNotIn applies the NotIn predicate on the "Min" field.
func MinNotIn(vs ...float64) predicate.Transactionfactoritemtmp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMin), v...))
	})
}

// MinGT applies the GT predicate on the "Min" field.
func MinGT(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMin), v))
	})
}

// MinGTE applies the GTE predicate on the "Min" field.
func MinGTE(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMin), v))
	})
}

// MinLT applies the LT predicate on the "Min" field.
func MinLT(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMin), v))
	})
}

// MinLTE applies the LTE predicate on the "Min" field.
func MinLTE(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMin), v))
	})
}

// MinIsNil applies the IsNil predicate on the "Min" field.
func MinIsNil() predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMin)))
	})
}

// MinNotNil applies the NotNil predicate on the "Min" field.
func MinNotNil() predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMin)))
	})
}

// MaxEQ applies the EQ predicate on the "Max" field.
func MaxEQ(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMax), v))
	})
}

// MaxNEQ applies the NEQ predicate on the "Max" field.
func MaxNEQ(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMax), v))
	})
}

// MaxIn applies the In predicate on the "Max" field.
func MaxIn(vs ...float64) predicate.Transactionfactoritemtmp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMax), v...))
	})
}

// MaxNotIn applies the NotIn predicate on the "Max" field.
func MaxNotIn(vs ...float64) predicate.Transactionfactoritemtmp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMax), v...))
	})
}

// MaxGT applies the GT predicate on the "Max" field.
func MaxGT(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMax), v))
	})
}

// MaxGTE applies the GTE predicate on the "Max" field.
func MaxGTE(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMax), v))
	})
}

// MaxLT applies the LT predicate on the "Max" field.
func MaxLT(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMax), v))
	})
}

// MaxLTE applies the LTE predicate on the "Max" field.
func MaxLTE(v float64) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMax), v))
	})
}

// MaxIsNil applies the IsNil predicate on the "Max" field.
func MaxIsNil() predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMax)))
	})
}

// MaxNotNil applies the NotNil predicate on the "Max" field.
func MaxNotNil() predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMax)))
	})
}

// RankEQ applies the EQ predicate on the "Rank" field.
func RankEQ(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "Rank" field.
func RankNEQ(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "Rank" field.
func RankIn(vs ...int) predicate.Transactionfactoritemtmp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRank), v...))
	})
}

// RankNotIn applies the NotIn predicate on the "Rank" field.
func RankNotIn(vs ...int) predicate.Transactionfactoritemtmp {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRank), v...))
	})
}

// RankGT applies the GT predicate on the "Rank" field.
func RankGT(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "Rank" field.
func RankGTE(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "Rank" field.
func RankLT(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "Rank" field.
func RankLTE(v int) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// RankIsNil applies the IsNil predicate on the "Rank" field.
func RankIsNil() predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRank)))
	})
}

// RankNotNil applies the NotNil predicate on the "Rank" field.
func RankNotNil() predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRank)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transactionfactoritemtmp) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transactionfactoritemtmp) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
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
func Not(p predicate.Transactionfactoritemtmp) predicate.Transactionfactoritemtmp {
	return predicate.Transactionfactoritemtmp(func(s *sql.Selector) {
		p(s.Not())
	})
}
