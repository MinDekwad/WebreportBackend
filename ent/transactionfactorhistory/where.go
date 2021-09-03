// Code generated by entc, DO NOT EDIT.

package transactionfactorhistory

import (
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// WalletID applies equality check predicate on the "WalletID" field. It's identical to WalletIDEQ.
func WalletID(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletID), v))
	})
}

// RankTransactionFactor applies equality check predicate on the "RankTransactionFactor" field. It's identical to RankTransactionFactorEQ.
func RankTransactionFactor(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRankTransactionFactor), v))
	})
}

// DateCalRank applies equality check predicate on the "DateCalRank" field. It's identical to DateCalRankEQ.
func DateCalRank(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCalRank), v))
	})
}

// WalletIDEQ applies the EQ predicate on the "WalletID" field.
func WalletIDEQ(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletID), v))
	})
}

// WalletIDNEQ applies the NEQ predicate on the "WalletID" field.
func WalletIDNEQ(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWalletID), v))
	})
}

// WalletIDIn applies the In predicate on the "WalletID" field.
func WalletIDIn(vs ...string) predicate.Transactionfactorhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func WalletIDNotIn(vs ...string) predicate.Transactionfactorhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func WalletIDGT(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWalletID), v))
	})
}

// WalletIDGTE applies the GTE predicate on the "WalletID" field.
func WalletIDGTE(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWalletID), v))
	})
}

// WalletIDLT applies the LT predicate on the "WalletID" field.
func WalletIDLT(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWalletID), v))
	})
}

// WalletIDLTE applies the LTE predicate on the "WalletID" field.
func WalletIDLTE(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWalletID), v))
	})
}

// WalletIDContains applies the Contains predicate on the "WalletID" field.
func WalletIDContains(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWalletID), v))
	})
}

// WalletIDHasPrefix applies the HasPrefix predicate on the "WalletID" field.
func WalletIDHasPrefix(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWalletID), v))
	})
}

// WalletIDHasSuffix applies the HasSuffix predicate on the "WalletID" field.
func WalletIDHasSuffix(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWalletID), v))
	})
}

// WalletIDEqualFold applies the EqualFold predicate on the "WalletID" field.
func WalletIDEqualFold(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWalletID), v))
	})
}

// WalletIDContainsFold applies the ContainsFold predicate on the "WalletID" field.
func WalletIDContainsFold(v string) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWalletID), v))
	})
}

// RankTransactionFactorEQ applies the EQ predicate on the "RankTransactionFactor" field.
func RankTransactionFactorEQ(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRankTransactionFactor), v))
	})
}

// RankTransactionFactorNEQ applies the NEQ predicate on the "RankTransactionFactor" field.
func RankTransactionFactorNEQ(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRankTransactionFactor), v))
	})
}

// RankTransactionFactorIn applies the In predicate on the "RankTransactionFactor" field.
func RankTransactionFactorIn(vs ...int) predicate.Transactionfactorhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRankTransactionFactor), v...))
	})
}

// RankTransactionFactorNotIn applies the NotIn predicate on the "RankTransactionFactor" field.
func RankTransactionFactorNotIn(vs ...int) predicate.Transactionfactorhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRankTransactionFactor), v...))
	})
}

// RankTransactionFactorGT applies the GT predicate on the "RankTransactionFactor" field.
func RankTransactionFactorGT(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRankTransactionFactor), v))
	})
}

// RankTransactionFactorGTE applies the GTE predicate on the "RankTransactionFactor" field.
func RankTransactionFactorGTE(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRankTransactionFactor), v))
	})
}

// RankTransactionFactorLT applies the LT predicate on the "RankTransactionFactor" field.
func RankTransactionFactorLT(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRankTransactionFactor), v))
	})
}

// RankTransactionFactorLTE applies the LTE predicate on the "RankTransactionFactor" field.
func RankTransactionFactorLTE(v int) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRankTransactionFactor), v))
	})
}

// RankTransactionFactorIsNil applies the IsNil predicate on the "RankTransactionFactor" field.
func RankTransactionFactorIsNil() predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRankTransactionFactor)))
	})
}

// RankTransactionFactorNotNil applies the NotNil predicate on the "RankTransactionFactor" field.
func RankTransactionFactorNotNil() predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRankTransactionFactor)))
	})
}

// DateCalRankEQ applies the EQ predicate on the "DateCalRank" field.
func DateCalRankEQ(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankNEQ applies the NEQ predicate on the "DateCalRank" field.
func DateCalRankNEQ(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankIn applies the In predicate on the "DateCalRank" field.
func DateCalRankIn(vs ...time.Time) predicate.Transactionfactorhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func DateCalRankNotIn(vs ...time.Time) predicate.Transactionfactorhistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func DateCalRankGT(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankGTE applies the GTE predicate on the "DateCalRank" field.
func DateCalRankGTE(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankLT applies the LT predicate on the "DateCalRank" field.
func DateCalRankLT(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankLTE applies the LTE predicate on the "DateCalRank" field.
func DateCalRankLTE(v time.Time) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateCalRank), v))
	})
}

// DateCalRankIsNil applies the IsNil predicate on the "DateCalRank" field.
func DateCalRankIsNil() predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDateCalRank)))
	})
}

// DateCalRankNotNil applies the NotNil predicate on the "DateCalRank" field.
func DateCalRankNotNil() predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDateCalRank)))
	})
}

// HasTransactionfactor applies the HasEdge predicate on the "Transactionfactor" edge.
func HasTransactionfactor() predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionfactorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransactionfactorTable, TransactionfactorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionfactorWith applies the HasEdge predicate on the "Transactionfactor" edge with a given conditions (other predicates).
func HasTransactionfactorWith(preds ...predicate.Transactionfactor) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionfactorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TransactionfactorTable, TransactionfactorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transactionfactorhistory) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transactionfactorhistory) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
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
func Not(p predicate.Transactionfactorhistory) predicate.Transactionfactorhistory {
	return predicate.Transactionfactorhistory(func(s *sql.Selector) {
		p(s.Not())
	})
}