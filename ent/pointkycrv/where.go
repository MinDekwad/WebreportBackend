// Code generated by entc, DO NOT EDIT.

package pointkycrv

import (
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// DateTimeGen applies equality check predicate on the "DateTimeGen" field. It's identical to DateTimeGenEQ.
func DateTimeGen(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateTimeGen), v))
	})
}

// WalletID applies equality check predicate on the "WalletID" field. It's identical to WalletIDEQ.
func WalletID(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletID), v))
	})
}

// StatusGen applies equality check predicate on the "StatusGen" field. It's identical to StatusGenEQ.
func StatusGen(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusGen), v))
	})
}

// KYCDate applies equality check predicate on the "KYCDate" field. It's identical to KYCDateEQ.
func KYCDate(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKYCDate), v))
	})
}

// RVDate applies equality check predicate on the "RVDate" field. It's identical to RVDateEQ.
func RVDate(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRVDate), v))
	})
}

// Type applies equality check predicate on the "Type" field. It's identical to TypeEQ.
func Type(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// Point applies equality check predicate on the "Point" field. It's identical to PointEQ.
func Point(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoint), v))
	})
}

// FileimportID applies equality check predicate on the "FileimportID" field. It's identical to FileimportIDEQ.
func FileimportID(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileimportID), v))
	})
}

// DateTimeGenEQ applies the EQ predicate on the "DateTimeGen" field.
func DateTimeGenEQ(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDateTimeGen), v))
	})
}

// DateTimeGenNEQ applies the NEQ predicate on the "DateTimeGen" field.
func DateTimeGenNEQ(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDateTimeGen), v))
	})
}

// DateTimeGenIn applies the In predicate on the "DateTimeGen" field.
func DateTimeGenIn(vs ...time.Time) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDateTimeGen), v...))
	})
}

// DateTimeGenNotIn applies the NotIn predicate on the "DateTimeGen" field.
func DateTimeGenNotIn(vs ...time.Time) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDateTimeGen), v...))
	})
}

// DateTimeGenGT applies the GT predicate on the "DateTimeGen" field.
func DateTimeGenGT(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDateTimeGen), v))
	})
}

// DateTimeGenGTE applies the GTE predicate on the "DateTimeGen" field.
func DateTimeGenGTE(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDateTimeGen), v))
	})
}

// DateTimeGenLT applies the LT predicate on the "DateTimeGen" field.
func DateTimeGenLT(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDateTimeGen), v))
	})
}

// DateTimeGenLTE applies the LTE predicate on the "DateTimeGen" field.
func DateTimeGenLTE(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDateTimeGen), v))
	})
}

// DateTimeGenIsNil applies the IsNil predicate on the "DateTimeGen" field.
func DateTimeGenIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDateTimeGen)))
	})
}

// DateTimeGenNotNil applies the NotNil predicate on the "DateTimeGen" field.
func DateTimeGenNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDateTimeGen)))
	})
}

// WalletIDEQ applies the EQ predicate on the "WalletID" field.
func WalletIDEQ(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWalletID), v))
	})
}

// WalletIDNEQ applies the NEQ predicate on the "WalletID" field.
func WalletIDNEQ(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWalletID), v))
	})
}

// WalletIDIn applies the In predicate on the "WalletID" field.
func WalletIDIn(vs ...string) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
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
func WalletIDNotIn(vs ...string) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
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
func WalletIDGT(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWalletID), v))
	})
}

// WalletIDGTE applies the GTE predicate on the "WalletID" field.
func WalletIDGTE(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWalletID), v))
	})
}

// WalletIDLT applies the LT predicate on the "WalletID" field.
func WalletIDLT(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWalletID), v))
	})
}

// WalletIDLTE applies the LTE predicate on the "WalletID" field.
func WalletIDLTE(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWalletID), v))
	})
}

// WalletIDContains applies the Contains predicate on the "WalletID" field.
func WalletIDContains(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldWalletID), v))
	})
}

// WalletIDHasPrefix applies the HasPrefix predicate on the "WalletID" field.
func WalletIDHasPrefix(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldWalletID), v))
	})
}

// WalletIDHasSuffix applies the HasSuffix predicate on the "WalletID" field.
func WalletIDHasSuffix(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldWalletID), v))
	})
}

// WalletIDIsNil applies the IsNil predicate on the "WalletID" field.
func WalletIDIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWalletID)))
	})
}

// WalletIDNotNil applies the NotNil predicate on the "WalletID" field.
func WalletIDNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWalletID)))
	})
}

// WalletIDEqualFold applies the EqualFold predicate on the "WalletID" field.
func WalletIDEqualFold(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldWalletID), v))
	})
}

// WalletIDContainsFold applies the ContainsFold predicate on the "WalletID" field.
func WalletIDContainsFold(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldWalletID), v))
	})
}

// StatusGenEQ applies the EQ predicate on the "StatusGen" field.
func StatusGenEQ(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusGen), v))
	})
}

// StatusGenNEQ applies the NEQ predicate on the "StatusGen" field.
func StatusGenNEQ(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusGen), v))
	})
}

// StatusGenIn applies the In predicate on the "StatusGen" field.
func StatusGenIn(vs ...string) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusGen), v...))
	})
}

// StatusGenNotIn applies the NotIn predicate on the "StatusGen" field.
func StatusGenNotIn(vs ...string) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusGen), v...))
	})
}

// StatusGenGT applies the GT predicate on the "StatusGen" field.
func StatusGenGT(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusGen), v))
	})
}

// StatusGenGTE applies the GTE predicate on the "StatusGen" field.
func StatusGenGTE(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusGen), v))
	})
}

// StatusGenLT applies the LT predicate on the "StatusGen" field.
func StatusGenLT(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusGen), v))
	})
}

// StatusGenLTE applies the LTE predicate on the "StatusGen" field.
func StatusGenLTE(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusGen), v))
	})
}

// StatusGenContains applies the Contains predicate on the "StatusGen" field.
func StatusGenContains(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatusGen), v))
	})
}

// StatusGenHasPrefix applies the HasPrefix predicate on the "StatusGen" field.
func StatusGenHasPrefix(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatusGen), v))
	})
}

// StatusGenHasSuffix applies the HasSuffix predicate on the "StatusGen" field.
func StatusGenHasSuffix(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatusGen), v))
	})
}

// StatusGenIsNil applies the IsNil predicate on the "StatusGen" field.
func StatusGenIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatusGen)))
	})
}

// StatusGenNotNil applies the NotNil predicate on the "StatusGen" field.
func StatusGenNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatusGen)))
	})
}

// StatusGenEqualFold applies the EqualFold predicate on the "StatusGen" field.
func StatusGenEqualFold(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatusGen), v))
	})
}

// StatusGenContainsFold applies the ContainsFold predicate on the "StatusGen" field.
func StatusGenContainsFold(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatusGen), v))
	})
}

// KYCDateEQ applies the EQ predicate on the "KYCDate" field.
func KYCDateEQ(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKYCDate), v))
	})
}

// KYCDateNEQ applies the NEQ predicate on the "KYCDate" field.
func KYCDateNEQ(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKYCDate), v))
	})
}

// KYCDateIn applies the In predicate on the "KYCDate" field.
func KYCDateIn(vs ...time.Time) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldKYCDate), v...))
	})
}

// KYCDateNotIn applies the NotIn predicate on the "KYCDate" field.
func KYCDateNotIn(vs ...time.Time) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldKYCDate), v...))
	})
}

// KYCDateGT applies the GT predicate on the "KYCDate" field.
func KYCDateGT(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKYCDate), v))
	})
}

// KYCDateGTE applies the GTE predicate on the "KYCDate" field.
func KYCDateGTE(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKYCDate), v))
	})
}

// KYCDateLT applies the LT predicate on the "KYCDate" field.
func KYCDateLT(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKYCDate), v))
	})
}

// KYCDateLTE applies the LTE predicate on the "KYCDate" field.
func KYCDateLTE(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKYCDate), v))
	})
}

// KYCDateIsNil applies the IsNil predicate on the "KYCDate" field.
func KYCDateIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldKYCDate)))
	})
}

// KYCDateNotNil applies the NotNil predicate on the "KYCDate" field.
func KYCDateNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldKYCDate)))
	})
}

// RVDateEQ applies the EQ predicate on the "RVDate" field.
func RVDateEQ(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRVDate), v))
	})
}

// RVDateNEQ applies the NEQ predicate on the "RVDate" field.
func RVDateNEQ(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRVDate), v))
	})
}

// RVDateIn applies the In predicate on the "RVDate" field.
func RVDateIn(vs ...time.Time) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldRVDate), v...))
	})
}

// RVDateNotIn applies the NotIn predicate on the "RVDate" field.
func RVDateNotIn(vs ...time.Time) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldRVDate), v...))
	})
}

// RVDateGT applies the GT predicate on the "RVDate" field.
func RVDateGT(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRVDate), v))
	})
}

// RVDateGTE applies the GTE predicate on the "RVDate" field.
func RVDateGTE(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRVDate), v))
	})
}

// RVDateLT applies the LT predicate on the "RVDate" field.
func RVDateLT(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRVDate), v))
	})
}

// RVDateLTE applies the LTE predicate on the "RVDate" field.
func RVDateLTE(v time.Time) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRVDate), v))
	})
}

// RVDateIsNil applies the IsNil predicate on the "RVDate" field.
func RVDateIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRVDate)))
	})
}

// RVDateNotNil applies the NotNil predicate on the "RVDate" field.
func RVDateNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRVDate)))
	})
}

// TypeEQ applies the EQ predicate on the "Type" field.
func TypeEQ(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldType), v))
	})
}

// TypeNEQ applies the NEQ predicate on the "Type" field.
func TypeNEQ(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldType), v))
	})
}

// TypeIn applies the In predicate on the "Type" field.
func TypeIn(vs ...string) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldType), v...))
	})
}

// TypeNotIn applies the NotIn predicate on the "Type" field.
func TypeNotIn(vs ...string) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldType), v...))
	})
}

// TypeGT applies the GT predicate on the "Type" field.
func TypeGT(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldType), v))
	})
}

// TypeGTE applies the GTE predicate on the "Type" field.
func TypeGTE(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldType), v))
	})
}

// TypeLT applies the LT predicate on the "Type" field.
func TypeLT(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldType), v))
	})
}

// TypeLTE applies the LTE predicate on the "Type" field.
func TypeLTE(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldType), v))
	})
}

// TypeContains applies the Contains predicate on the "Type" field.
func TypeContains(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldType), v))
	})
}

// TypeHasPrefix applies the HasPrefix predicate on the "Type" field.
func TypeHasPrefix(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldType), v))
	})
}

// TypeHasSuffix applies the HasSuffix predicate on the "Type" field.
func TypeHasSuffix(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldType), v))
	})
}

// TypeIsNil applies the IsNil predicate on the "Type" field.
func TypeIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldType)))
	})
}

// TypeNotNil applies the NotNil predicate on the "Type" field.
func TypeNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldType)))
	})
}

// TypeEqualFold applies the EqualFold predicate on the "Type" field.
func TypeEqualFold(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldType), v))
	})
}

// TypeContainsFold applies the ContainsFold predicate on the "Type" field.
func TypeContainsFold(v string) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldType), v))
	})
}

// PointEQ applies the EQ predicate on the "Point" field.
func PointEQ(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPoint), v))
	})
}

// PointNEQ applies the NEQ predicate on the "Point" field.
func PointNEQ(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPoint), v))
	})
}

// PointIn applies the In predicate on the "Point" field.
func PointIn(vs ...int) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPoint), v...))
	})
}

// PointNotIn applies the NotIn predicate on the "Point" field.
func PointNotIn(vs ...int) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPoint), v...))
	})
}

// PointGT applies the GT predicate on the "Point" field.
func PointGT(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPoint), v))
	})
}

// PointGTE applies the GTE predicate on the "Point" field.
func PointGTE(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPoint), v))
	})
}

// PointLT applies the LT predicate on the "Point" field.
func PointLT(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPoint), v))
	})
}

// PointLTE applies the LTE predicate on the "Point" field.
func PointLTE(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPoint), v))
	})
}

// PointIsNil applies the IsNil predicate on the "Point" field.
func PointIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPoint)))
	})
}

// PointNotNil applies the NotNil predicate on the "Point" field.
func PointNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPoint)))
	})
}

// FileimportIDEQ applies the EQ predicate on the "FileimportID" field.
func FileimportIDEQ(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFileimportID), v))
	})
}

// FileimportIDNEQ applies the NEQ predicate on the "FileimportID" field.
func FileimportIDNEQ(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFileimportID), v))
	})
}

// FileimportIDIn applies the In predicate on the "FileimportID" field.
func FileimportIDIn(vs ...int) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFileimportID), v...))
	})
}

// FileimportIDNotIn applies the NotIn predicate on the "FileimportID" field.
func FileimportIDNotIn(vs ...int) predicate.Pointkycrv {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pointkycrv(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFileimportID), v...))
	})
}

// FileimportIDGT applies the GT predicate on the "FileimportID" field.
func FileimportIDGT(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFileimportID), v))
	})
}

// FileimportIDGTE applies the GTE predicate on the "FileimportID" field.
func FileimportIDGTE(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFileimportID), v))
	})
}

// FileimportIDLT applies the LT predicate on the "FileimportID" field.
func FileimportIDLT(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFileimportID), v))
	})
}

// FileimportIDLTE applies the LTE predicate on the "FileimportID" field.
func FileimportIDLTE(v int) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFileimportID), v))
	})
}

// FileimportIDIsNil applies the IsNil predicate on the "FileimportID" field.
func FileimportIDIsNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldFileimportID)))
	})
}

// FileimportIDNotNil applies the NotNil predicate on the "FileimportID" field.
func FileimportIDNotNil() predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldFileimportID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pointkycrv) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pointkycrv) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
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
func Not(p predicate.Pointkycrv) predicate.Pointkycrv {
	return predicate.Pointkycrv(func(s *sql.Selector) {
		p(s.Not())
	})
}
