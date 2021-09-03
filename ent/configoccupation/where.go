// Code generated by entc, DO NOT EDIT.

package configoccupation

import (
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// OccupationName applies equality check predicate on the "OccupationName" field. It's identical to OccupationNameEQ.
func OccupationName(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccupationName), v))
	})
}

// Rank applies equality check predicate on the "Rank" field. It's identical to RankEQ.
func Rank(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankTmp applies equality check predicate on the "RankTmp" field. It's identical to RankTmpEQ.
func RankTmp(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRankTmp), v))
	})
}

// UpdateDate applies equality check predicate on the "UpdateDate" field. It's identical to UpdateDateEQ.
func UpdateDate(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateDate), v))
	})
}

// ApproveBy applies equality check predicate on the "ApproveBy" field. It's identical to ApproveByEQ.
func ApproveBy(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApproveBy), v))
	})
}

// ApproveDate applies equality check predicate on the "ApproveDate" field. It's identical to ApproveDateEQ.
func ApproveDate(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApproveDate), v))
	})
}

// OccupationNameEQ applies the EQ predicate on the "OccupationName" field.
func OccupationNameEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOccupationName), v))
	})
}

// OccupationNameNEQ applies the NEQ predicate on the "OccupationName" field.
func OccupationNameNEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOccupationName), v))
	})
}

// OccupationNameIn applies the In predicate on the "OccupationName" field.
func OccupationNameIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func OccupationNameNotIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func OccupationNameGT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOccupationName), v))
	})
}

// OccupationNameGTE applies the GTE predicate on the "OccupationName" field.
func OccupationNameGTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOccupationName), v))
	})
}

// OccupationNameLT applies the LT predicate on the "OccupationName" field.
func OccupationNameLT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOccupationName), v))
	})
}

// OccupationNameLTE applies the LTE predicate on the "OccupationName" field.
func OccupationNameLTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOccupationName), v))
	})
}

// OccupationNameContains applies the Contains predicate on the "OccupationName" field.
func OccupationNameContains(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldOccupationName), v))
	})
}

// OccupationNameHasPrefix applies the HasPrefix predicate on the "OccupationName" field.
func OccupationNameHasPrefix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldOccupationName), v))
	})
}

// OccupationNameHasSuffix applies the HasSuffix predicate on the "OccupationName" field.
func OccupationNameHasSuffix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldOccupationName), v))
	})
}

// OccupationNameIsNil applies the IsNil predicate on the "OccupationName" field.
func OccupationNameIsNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOccupationName)))
	})
}

// OccupationNameNotNil applies the NotNil predicate on the "OccupationName" field.
func OccupationNameNotNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOccupationName)))
	})
}

// OccupationNameEqualFold applies the EqualFold predicate on the "OccupationName" field.
func OccupationNameEqualFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldOccupationName), v))
	})
}

// OccupationNameContainsFold applies the ContainsFold predicate on the "OccupationName" field.
func OccupationNameContainsFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldOccupationName), v))
	})
}

// RankEQ applies the EQ predicate on the "Rank" field.
func RankEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRank), v))
	})
}

// RankNEQ applies the NEQ predicate on the "Rank" field.
func RankNEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRank), v))
	})
}

// RankIn applies the In predicate on the "Rank" field.
func RankIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func RankNotIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func RankGT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRank), v))
	})
}

// RankGTE applies the GTE predicate on the "Rank" field.
func RankGTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRank), v))
	})
}

// RankLT applies the LT predicate on the "Rank" field.
func RankLT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRank), v))
	})
}

// RankLTE applies the LTE predicate on the "Rank" field.
func RankLTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRank), v))
	})
}

// RankContains applies the Contains predicate on the "Rank" field.
func RankContains(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRank), v))
	})
}

// RankHasPrefix applies the HasPrefix predicate on the "Rank" field.
func RankHasPrefix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRank), v))
	})
}

// RankHasSuffix applies the HasSuffix predicate on the "Rank" field.
func RankHasSuffix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRank), v))
	})
}

// RankIsNil applies the IsNil predicate on the "Rank" field.
func RankIsNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRank)))
	})
}

// RankNotNil applies the NotNil predicate on the "Rank" field.
func RankNotNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRank)))
	})
}

// RankEqualFold applies the EqualFold predicate on the "Rank" field.
func RankEqualFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRank), v))
	})
}

// RankContainsFold applies the ContainsFold predicate on the "Rank" field.
func RankContainsFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRank), v))
	})
}

// RankTmpEQ applies the EQ predicate on the "RankTmp" field.
func RankTmpEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRankTmp), v))
	})
}

// RankTmpNEQ applies the NEQ predicate on the "RankTmp" field.
func RankTmpNEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRankTmp), v))
	})
}

// RankTmpIn applies the In predicate on the "RankTmp" field.
func RankTmpIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRankTmp), v...))
	})
}

// RankTmpNotIn applies the NotIn predicate on the "RankTmp" field.
func RankTmpNotIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRankTmp), v...))
	})
}

// RankTmpGT applies the GT predicate on the "RankTmp" field.
func RankTmpGT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRankTmp), v))
	})
}

// RankTmpGTE applies the GTE predicate on the "RankTmp" field.
func RankTmpGTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRankTmp), v))
	})
}

// RankTmpLT applies the LT predicate on the "RankTmp" field.
func RankTmpLT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRankTmp), v))
	})
}

// RankTmpLTE applies the LTE predicate on the "RankTmp" field.
func RankTmpLTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRankTmp), v))
	})
}

// RankTmpContains applies the Contains predicate on the "RankTmp" field.
func RankTmpContains(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldRankTmp), v))
	})
}

// RankTmpHasPrefix applies the HasPrefix predicate on the "RankTmp" field.
func RankTmpHasPrefix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldRankTmp), v))
	})
}

// RankTmpHasSuffix applies the HasSuffix predicate on the "RankTmp" field.
func RankTmpHasSuffix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldRankTmp), v))
	})
}

// RankTmpIsNil applies the IsNil predicate on the "RankTmp" field.
func RankTmpIsNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRankTmp)))
	})
}

// RankTmpNotNil applies the NotNil predicate on the "RankTmp" field.
func RankTmpNotNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRankTmp)))
	})
}

// RankTmpEqualFold applies the EqualFold predicate on the "RankTmp" field.
func RankTmpEqualFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldRankTmp), v))
	})
}

// RankTmpContainsFold applies the ContainsFold predicate on the "RankTmp" field.
func RankTmpContainsFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldRankTmp), v))
	})
}

// UpdateDateEQ applies the EQ predicate on the "UpdateDate" field.
func UpdateDateEQ(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateNEQ applies the NEQ predicate on the "UpdateDate" field.
func UpdateDateNEQ(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateIn applies the In predicate on the "UpdateDate" field.
func UpdateDateIn(vs ...time.Time) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateDate), v...))
	})
}

// UpdateDateNotIn applies the NotIn predicate on the "UpdateDate" field.
func UpdateDateNotIn(vs ...time.Time) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateDate), v...))
	})
}

// UpdateDateGT applies the GT predicate on the "UpdateDate" field.
func UpdateDateGT(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateGTE applies the GTE predicate on the "UpdateDate" field.
func UpdateDateGTE(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateLT applies the LT predicate on the "UpdateDate" field.
func UpdateDateLT(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateLTE applies the LTE predicate on the "UpdateDate" field.
func UpdateDateLTE(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateIsNil applies the IsNil predicate on the "UpdateDate" field.
func UpdateDateIsNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateDate)))
	})
}

// UpdateDateNotNil applies the NotNil predicate on the "UpdateDate" field.
func UpdateDateNotNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateDate)))
	})
}

// ApproveByEQ applies the EQ predicate on the "ApproveBy" field.
func ApproveByEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApproveBy), v))
	})
}

// ApproveByNEQ applies the NEQ predicate on the "ApproveBy" field.
func ApproveByNEQ(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApproveBy), v))
	})
}

// ApproveByIn applies the In predicate on the "ApproveBy" field.
func ApproveByIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApproveBy), v...))
	})
}

// ApproveByNotIn applies the NotIn predicate on the "ApproveBy" field.
func ApproveByNotIn(vs ...string) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApproveBy), v...))
	})
}

// ApproveByGT applies the GT predicate on the "ApproveBy" field.
func ApproveByGT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApproveBy), v))
	})
}

// ApproveByGTE applies the GTE predicate on the "ApproveBy" field.
func ApproveByGTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApproveBy), v))
	})
}

// ApproveByLT applies the LT predicate on the "ApproveBy" field.
func ApproveByLT(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApproveBy), v))
	})
}

// ApproveByLTE applies the LTE predicate on the "ApproveBy" field.
func ApproveByLTE(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApproveBy), v))
	})
}

// ApproveByContains applies the Contains predicate on the "ApproveBy" field.
func ApproveByContains(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldApproveBy), v))
	})
}

// ApproveByHasPrefix applies the HasPrefix predicate on the "ApproveBy" field.
func ApproveByHasPrefix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldApproveBy), v))
	})
}

// ApproveByHasSuffix applies the HasSuffix predicate on the "ApproveBy" field.
func ApproveByHasSuffix(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldApproveBy), v))
	})
}

// ApproveByIsNil applies the IsNil predicate on the "ApproveBy" field.
func ApproveByIsNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldApproveBy)))
	})
}

// ApproveByNotNil applies the NotNil predicate on the "ApproveBy" field.
func ApproveByNotNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldApproveBy)))
	})
}

// ApproveByEqualFold applies the EqualFold predicate on the "ApproveBy" field.
func ApproveByEqualFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldApproveBy), v))
	})
}

// ApproveByContainsFold applies the ContainsFold predicate on the "ApproveBy" field.
func ApproveByContainsFold(v string) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldApproveBy), v))
	})
}

// ApproveDateEQ applies the EQ predicate on the "ApproveDate" field.
func ApproveDateEQ(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldApproveDate), v))
	})
}

// ApproveDateNEQ applies the NEQ predicate on the "ApproveDate" field.
func ApproveDateNEQ(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldApproveDate), v))
	})
}

// ApproveDateIn applies the In predicate on the "ApproveDate" field.
func ApproveDateIn(vs ...time.Time) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldApproveDate), v...))
	})
}

// ApproveDateNotIn applies the NotIn predicate on the "ApproveDate" field.
func ApproveDateNotIn(vs ...time.Time) predicate.Configoccupation {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Configoccupation(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldApproveDate), v...))
	})
}

// ApproveDateGT applies the GT predicate on the "ApproveDate" field.
func ApproveDateGT(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldApproveDate), v))
	})
}

// ApproveDateGTE applies the GTE predicate on the "ApproveDate" field.
func ApproveDateGTE(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldApproveDate), v))
	})
}

// ApproveDateLT applies the LT predicate on the "ApproveDate" field.
func ApproveDateLT(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldApproveDate), v))
	})
}

// ApproveDateLTE applies the LTE predicate on the "ApproveDate" field.
func ApproveDateLTE(v time.Time) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldApproveDate), v))
	})
}

// ApproveDateIsNil applies the IsNil predicate on the "ApproveDate" field.
func ApproveDateIsNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldApproveDate)))
	})
}

// ApproveDateNotNil applies the NotNil predicate on the "ApproveDate" field.
func ApproveDateNotNil() predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldApproveDate)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Configoccupation) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Configoccupation) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
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
func Not(p predicate.Configoccupation) predicate.Configoccupation {
	return predicate.Configoccupation(func(s *sql.Selector) {
		p(s.Not())
	})
}
