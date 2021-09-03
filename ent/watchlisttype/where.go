// Code generated by entc, DO NOT EDIT.

package watchlisttype

import (
	"go-api-report2/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TypeID applies equality check predicate on the "TypeID" field. It's identical to TypeIDEQ.
func TypeID(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeID), v))
	})
}

// TypeName applies equality check predicate on the "TypeName" field. It's identical to TypeNameEQ.
func TypeName(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeName), v))
	})
}

// TypeDescription applies equality check predicate on the "TypeDescription" field. It's identical to TypeDescriptionEQ.
func TypeDescription(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeDescription), v))
	})
}

// TypeIDEQ applies the EQ predicate on the "TypeID" field.
func TypeIDEQ(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeID), v))
	})
}

// TypeIDNEQ applies the NEQ predicate on the "TypeID" field.
func TypeIDNEQ(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeID), v))
	})
}

// TypeIDIn applies the In predicate on the "TypeID" field.
func TypeIDIn(vs ...int) predicate.Watchlisttype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Watchlisttype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypeID), v...))
	})
}

// TypeIDNotIn applies the NotIn predicate on the "TypeID" field.
func TypeIDNotIn(vs ...int) predicate.Watchlisttype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Watchlisttype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypeID), v...))
	})
}

// TypeIDGT applies the GT predicate on the "TypeID" field.
func TypeIDGT(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeID), v))
	})
}

// TypeIDGTE applies the GTE predicate on the "TypeID" field.
func TypeIDGTE(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeID), v))
	})
}

// TypeIDLT applies the LT predicate on the "TypeID" field.
func TypeIDLT(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeID), v))
	})
}

// TypeIDLTE applies the LTE predicate on the "TypeID" field.
func TypeIDLTE(v int) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeID), v))
	})
}

// TypeIDIsNil applies the IsNil predicate on the "TypeID" field.
func TypeIDIsNil() predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTypeID)))
	})
}

// TypeIDNotNil applies the NotNil predicate on the "TypeID" field.
func TypeIDNotNil() predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTypeID)))
	})
}

// TypeNameEQ applies the EQ predicate on the "TypeName" field.
func TypeNameEQ(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeName), v))
	})
}

// TypeNameNEQ applies the NEQ predicate on the "TypeName" field.
func TypeNameNEQ(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeName), v))
	})
}

// TypeNameIn applies the In predicate on the "TypeName" field.
func TypeNameIn(vs ...string) predicate.Watchlisttype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Watchlisttype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypeName), v...))
	})
}

// TypeNameNotIn applies the NotIn predicate on the "TypeName" field.
func TypeNameNotIn(vs ...string) predicate.Watchlisttype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Watchlisttype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypeName), v...))
	})
}

// TypeNameGT applies the GT predicate on the "TypeName" field.
func TypeNameGT(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeName), v))
	})
}

// TypeNameGTE applies the GTE predicate on the "TypeName" field.
func TypeNameGTE(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeName), v))
	})
}

// TypeNameLT applies the LT predicate on the "TypeName" field.
func TypeNameLT(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeName), v))
	})
}

// TypeNameLTE applies the LTE predicate on the "TypeName" field.
func TypeNameLTE(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeName), v))
	})
}

// TypeNameContains applies the Contains predicate on the "TypeName" field.
func TypeNameContains(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTypeName), v))
	})
}

// TypeNameHasPrefix applies the HasPrefix predicate on the "TypeName" field.
func TypeNameHasPrefix(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTypeName), v))
	})
}

// TypeNameHasSuffix applies the HasSuffix predicate on the "TypeName" field.
func TypeNameHasSuffix(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTypeName), v))
	})
}

// TypeNameEqualFold applies the EqualFold predicate on the "TypeName" field.
func TypeNameEqualFold(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTypeName), v))
	})
}

// TypeNameContainsFold applies the ContainsFold predicate on the "TypeName" field.
func TypeNameContainsFold(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTypeName), v))
	})
}

// TypeDescriptionEQ applies the EQ predicate on the "TypeDescription" field.
func TypeDescriptionEQ(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionNEQ applies the NEQ predicate on the "TypeDescription" field.
func TypeDescriptionNEQ(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionIn applies the In predicate on the "TypeDescription" field.
func TypeDescriptionIn(vs ...string) predicate.Watchlisttype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Watchlisttype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTypeDescription), v...))
	})
}

// TypeDescriptionNotIn applies the NotIn predicate on the "TypeDescription" field.
func TypeDescriptionNotIn(vs ...string) predicate.Watchlisttype {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Watchlisttype(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTypeDescription), v...))
	})
}

// TypeDescriptionGT applies the GT predicate on the "TypeDescription" field.
func TypeDescriptionGT(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionGTE applies the GTE predicate on the "TypeDescription" field.
func TypeDescriptionGTE(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionLT applies the LT predicate on the "TypeDescription" field.
func TypeDescriptionLT(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionLTE applies the LTE predicate on the "TypeDescription" field.
func TypeDescriptionLTE(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionContains applies the Contains predicate on the "TypeDescription" field.
func TypeDescriptionContains(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionHasPrefix applies the HasPrefix predicate on the "TypeDescription" field.
func TypeDescriptionHasPrefix(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionHasSuffix applies the HasSuffix predicate on the "TypeDescription" field.
func TypeDescriptionHasSuffix(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionEqualFold applies the EqualFold predicate on the "TypeDescription" field.
func TypeDescriptionEqualFold(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTypeDescription), v))
	})
}

// TypeDescriptionContainsFold applies the ContainsFold predicate on the "TypeDescription" field.
func TypeDescriptionContainsFold(v string) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTypeDescription), v))
	})
}

// HasWatchlist applies the HasEdge predicate on the "watchlist" edge.
func HasWatchlist() predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WatchlistTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WatchlistTable, WatchlistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasWatchlistWith applies the HasEdge predicate on the "watchlist" edge with a given conditions (other predicates).
func HasWatchlistWith(preds ...predicate.Watchlist) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(WatchlistInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, WatchlistTable, WatchlistColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Watchlisttype) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Watchlisttype) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
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
func Not(p predicate.Watchlisttype) predicate.Watchlisttype {
	return predicate.Watchlisttype(func(s *sql.Selector) {
		p(s.Not())
	})
}
