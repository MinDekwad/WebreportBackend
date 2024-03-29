// Code generated by entc, DO NOT EDIT.

package transactionfactor

import (
	"go-api-report2/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// IDIsNil applies the IsNil predicate on the ID field.
func IDIsNil(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldID)))
	})
}

// IDNotNil applies the NotNil predicate on the ID field.
func IDNotNil(id int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldID)))
	})
}

// TransactionFactorName applies equality check predicate on the "TransactionFactorName" field. It's identical to TransactionFactorNameEQ.
func TransactionFactorName(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionType applies equality check predicate on the "TransactionType" field. It's identical to TransactionTypeEQ.
func TransactionType(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionType), v))
	})
}

// PaymentChannel applies equality check predicate on the "PaymentChannel" field. It's identical to PaymentChannelEQ.
func PaymentChannel(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentChannel), v))
	})
}

// PaymentType applies equality check predicate on the "PaymentType" field. It's identical to PaymentTypeEQ.
func PaymentType(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentType), v))
	})
}

// NumDay applies equality check predicate on the "NumDay" field. It's identical to NumDayEQ.
func NumDay(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumDay), v))
	})
}

// Date applies equality check predicate on the "Date" field. It's identical to DateEQ.
func Date(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// UpdateDate applies equality check predicate on the "UpdateDate" field. It's identical to UpdateDateEQ.
func UpdateDate(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateDate), v))
	})
}

// StatusApprove applies equality check predicate on the "StatusApprove" field. It's identical to StatusApproveEQ.
func StatusApprove(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusApprove), v))
	})
}

// TransactionFactorNameEQ applies the EQ predicate on the "TransactionFactorName" field.
func TransactionFactorNameEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameNEQ applies the NEQ predicate on the "TransactionFactorName" field.
func TransactionFactorNameNEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameIn applies the In predicate on the "TransactionFactorName" field.
func TransactionFactorNameIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransactionFactorName), v...))
	})
}

// TransactionFactorNameNotIn applies the NotIn predicate on the "TransactionFactorName" field.
func TransactionFactorNameNotIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransactionFactorName), v...))
	})
}

// TransactionFactorNameGT applies the GT predicate on the "TransactionFactorName" field.
func TransactionFactorNameGT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameGTE applies the GTE predicate on the "TransactionFactorName" field.
func TransactionFactorNameGTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameLT applies the LT predicate on the "TransactionFactorName" field.
func TransactionFactorNameLT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameLTE applies the LTE predicate on the "TransactionFactorName" field.
func TransactionFactorNameLTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameContains applies the Contains predicate on the "TransactionFactorName" field.
func TransactionFactorNameContains(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameHasPrefix applies the HasPrefix predicate on the "TransactionFactorName" field.
func TransactionFactorNameHasPrefix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameHasSuffix applies the HasSuffix predicate on the "TransactionFactorName" field.
func TransactionFactorNameHasSuffix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameIsNil applies the IsNil predicate on the "TransactionFactorName" field.
func TransactionFactorNameIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTransactionFactorName)))
	})
}

// TransactionFactorNameNotNil applies the NotNil predicate on the "TransactionFactorName" field.
func TransactionFactorNameNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTransactionFactorName)))
	})
}

// TransactionFactorNameEqualFold applies the EqualFold predicate on the "TransactionFactorName" field.
func TransactionFactorNameEqualFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionFactorNameContainsFold applies the ContainsFold predicate on the "TransactionFactorName" field.
func TransactionFactorNameContainsFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionFactorName), v))
	})
}

// TransactionTypeEQ applies the EQ predicate on the "TransactionType" field.
func TransactionTypeEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeNEQ applies the NEQ predicate on the "TransactionType" field.
func TransactionTypeNEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeIn applies the In predicate on the "TransactionType" field.
func TransactionTypeIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTransactionType), v...))
	})
}

// TransactionTypeNotIn applies the NotIn predicate on the "TransactionType" field.
func TransactionTypeNotIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTransactionType), v...))
	})
}

// TransactionTypeGT applies the GT predicate on the "TransactionType" field.
func TransactionTypeGT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeGTE applies the GTE predicate on the "TransactionType" field.
func TransactionTypeGTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeLT applies the LT predicate on the "TransactionType" field.
func TransactionTypeLT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeLTE applies the LTE predicate on the "TransactionType" field.
func TransactionTypeLTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeContains applies the Contains predicate on the "TransactionType" field.
func TransactionTypeContains(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeHasPrefix applies the HasPrefix predicate on the "TransactionType" field.
func TransactionTypeHasPrefix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeHasSuffix applies the HasSuffix predicate on the "TransactionType" field.
func TransactionTypeHasSuffix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeIsNil applies the IsNil predicate on the "TransactionType" field.
func TransactionTypeIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTransactionType)))
	})
}

// TransactionTypeNotNil applies the NotNil predicate on the "TransactionType" field.
func TransactionTypeNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTransactionType)))
	})
}

// TransactionTypeEqualFold applies the EqualFold predicate on the "TransactionType" field.
func TransactionTypeEqualFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTransactionType), v))
	})
}

// TransactionTypeContainsFold applies the ContainsFold predicate on the "TransactionType" field.
func TransactionTypeContainsFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTransactionType), v))
	})
}

// PaymentChannelEQ applies the EQ predicate on the "PaymentChannel" field.
func PaymentChannelEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelNEQ applies the NEQ predicate on the "PaymentChannel" field.
func PaymentChannelNEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelIn applies the In predicate on the "PaymentChannel" field.
func PaymentChannelIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaymentChannel), v...))
	})
}

// PaymentChannelNotIn applies the NotIn predicate on the "PaymentChannel" field.
func PaymentChannelNotIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaymentChannel), v...))
	})
}

// PaymentChannelGT applies the GT predicate on the "PaymentChannel" field.
func PaymentChannelGT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelGTE applies the GTE predicate on the "PaymentChannel" field.
func PaymentChannelGTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelLT applies the LT predicate on the "PaymentChannel" field.
func PaymentChannelLT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelLTE applies the LTE predicate on the "PaymentChannel" field.
func PaymentChannelLTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelContains applies the Contains predicate on the "PaymentChannel" field.
func PaymentChannelContains(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelHasPrefix applies the HasPrefix predicate on the "PaymentChannel" field.
func PaymentChannelHasPrefix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelHasSuffix applies the HasSuffix predicate on the "PaymentChannel" field.
func PaymentChannelHasSuffix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelIsNil applies the IsNil predicate on the "PaymentChannel" field.
func PaymentChannelIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentChannel)))
	})
}

// PaymentChannelNotNil applies the NotNil predicate on the "PaymentChannel" field.
func PaymentChannelNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentChannel)))
	})
}

// PaymentChannelEqualFold applies the EqualFold predicate on the "PaymentChannel" field.
func PaymentChannelEqualFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaymentChannel), v))
	})
}

// PaymentChannelContainsFold applies the ContainsFold predicate on the "PaymentChannel" field.
func PaymentChannelContainsFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaymentChannel), v))
	})
}

// PaymentTypeEQ applies the EQ predicate on the "PaymentType" field.
func PaymentTypeEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeNEQ applies the NEQ predicate on the "PaymentType" field.
func PaymentTypeNEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeIn applies the In predicate on the "PaymentType" field.
func PaymentTypeIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPaymentType), v...))
	})
}

// PaymentTypeNotIn applies the NotIn predicate on the "PaymentType" field.
func PaymentTypeNotIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPaymentType), v...))
	})
}

// PaymentTypeGT applies the GT predicate on the "PaymentType" field.
func PaymentTypeGT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeGTE applies the GTE predicate on the "PaymentType" field.
func PaymentTypeGTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeLT applies the LT predicate on the "PaymentType" field.
func PaymentTypeLT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeLTE applies the LTE predicate on the "PaymentType" field.
func PaymentTypeLTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeContains applies the Contains predicate on the "PaymentType" field.
func PaymentTypeContains(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeHasPrefix applies the HasPrefix predicate on the "PaymentType" field.
func PaymentTypeHasPrefix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeHasSuffix applies the HasSuffix predicate on the "PaymentType" field.
func PaymentTypeHasSuffix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeIsNil applies the IsNil predicate on the "PaymentType" field.
func PaymentTypeIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPaymentType)))
	})
}

// PaymentTypeNotNil applies the NotNil predicate on the "PaymentType" field.
func PaymentTypeNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPaymentType)))
	})
}

// PaymentTypeEqualFold applies the EqualFold predicate on the "PaymentType" field.
func PaymentTypeEqualFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPaymentType), v))
	})
}

// PaymentTypeContainsFold applies the ContainsFold predicate on the "PaymentType" field.
func PaymentTypeContainsFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPaymentType), v))
	})
}

// NumDayEQ applies the EQ predicate on the "NumDay" field.
func NumDayEQ(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNumDay), v))
	})
}

// NumDayNEQ applies the NEQ predicate on the "NumDay" field.
func NumDayNEQ(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNumDay), v))
	})
}

// NumDayIn applies the In predicate on the "NumDay" field.
func NumDayIn(vs ...int) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNumDay), v...))
	})
}

// NumDayNotIn applies the NotIn predicate on the "NumDay" field.
func NumDayNotIn(vs ...int) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNumDay), v...))
	})
}

// NumDayGT applies the GT predicate on the "NumDay" field.
func NumDayGT(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNumDay), v))
	})
}

// NumDayGTE applies the GTE predicate on the "NumDay" field.
func NumDayGTE(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNumDay), v))
	})
}

// NumDayLT applies the LT predicate on the "NumDay" field.
func NumDayLT(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNumDay), v))
	})
}

// NumDayLTE applies the LTE predicate on the "NumDay" field.
func NumDayLTE(v int) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNumDay), v))
	})
}

// NumDayIsNil applies the IsNil predicate on the "NumDay" field.
func NumDayIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNumDay)))
	})
}

// NumDayNotNil applies the NotNil predicate on the "NumDay" field.
func NumDayNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNumDay)))
	})
}

// DateEQ applies the EQ predicate on the "Date" field.
func DateEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDate), v))
	})
}

// DateNEQ applies the NEQ predicate on the "Date" field.
func DateNEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDate), v))
	})
}

// DateIn applies the In predicate on the "Date" field.
func DateIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDate), v...))
	})
}

// DateNotIn applies the NotIn predicate on the "Date" field.
func DateNotIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDate), v...))
	})
}

// DateGT applies the GT predicate on the "Date" field.
func DateGT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDate), v))
	})
}

// DateGTE applies the GTE predicate on the "Date" field.
func DateGTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDate), v))
	})
}

// DateLT applies the LT predicate on the "Date" field.
func DateLT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDate), v))
	})
}

// DateLTE applies the LTE predicate on the "Date" field.
func DateLTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDate), v))
	})
}

// DateContains applies the Contains predicate on the "Date" field.
func DateContains(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDate), v))
	})
}

// DateHasPrefix applies the HasPrefix predicate on the "Date" field.
func DateHasPrefix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDate), v))
	})
}

// DateHasSuffix applies the HasSuffix predicate on the "Date" field.
func DateHasSuffix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDate), v))
	})
}

// DateIsNil applies the IsNil predicate on the "Date" field.
func DateIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDate)))
	})
}

// DateNotNil applies the NotNil predicate on the "Date" field.
func DateNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDate)))
	})
}

// DateEqualFold applies the EqualFold predicate on the "Date" field.
func DateEqualFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDate), v))
	})
}

// DateContainsFold applies the ContainsFold predicate on the "Date" field.
func DateContainsFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDate), v))
	})
}

// UpdateDateEQ applies the EQ predicate on the "UpdateDate" field.
func UpdateDateEQ(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateNEQ applies the NEQ predicate on the "UpdateDate" field.
func UpdateDateNEQ(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateIn applies the In predicate on the "UpdateDate" field.
func UpdateDateIn(vs ...time.Time) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
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
func UpdateDateNotIn(vs ...time.Time) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
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
func UpdateDateGT(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateGTE applies the GTE predicate on the "UpdateDate" field.
func UpdateDateGTE(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateLT applies the LT predicate on the "UpdateDate" field.
func UpdateDateLT(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateLTE applies the LTE predicate on the "UpdateDate" field.
func UpdateDateLTE(v time.Time) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateDate), v))
	})
}

// UpdateDateIsNil applies the IsNil predicate on the "UpdateDate" field.
func UpdateDateIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdateDate)))
	})
}

// UpdateDateNotNil applies the NotNil predicate on the "UpdateDate" field.
func UpdateDateNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdateDate)))
	})
}

// StatusApproveEQ applies the EQ predicate on the "StatusApprove" field.
func StatusApproveEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveNEQ applies the NEQ predicate on the "StatusApprove" field.
func StatusApproveNEQ(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveIn applies the In predicate on the "StatusApprove" field.
func StatusApproveIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatusApprove), v...))
	})
}

// StatusApproveNotIn applies the NotIn predicate on the "StatusApprove" field.
func StatusApproveNotIn(vs ...string) predicate.Transactionfactor {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Transactionfactor(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatusApprove), v...))
	})
}

// StatusApproveGT applies the GT predicate on the "StatusApprove" field.
func StatusApproveGT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveGTE applies the GTE predicate on the "StatusApprove" field.
func StatusApproveGTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveLT applies the LT predicate on the "StatusApprove" field.
func StatusApproveLT(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveLTE applies the LTE predicate on the "StatusApprove" field.
func StatusApproveLTE(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveContains applies the Contains predicate on the "StatusApprove" field.
func StatusApproveContains(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveHasPrefix applies the HasPrefix predicate on the "StatusApprove" field.
func StatusApproveHasPrefix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveHasSuffix applies the HasSuffix predicate on the "StatusApprove" field.
func StatusApproveHasSuffix(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveIsNil applies the IsNil predicate on the "StatusApprove" field.
func StatusApproveIsNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStatusApprove)))
	})
}

// StatusApproveNotNil applies the NotNil predicate on the "StatusApprove" field.
func StatusApproveNotNil() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStatusApprove)))
	})
}

// StatusApproveEqualFold applies the EqualFold predicate on the "StatusApprove" field.
func StatusApproveEqualFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatusApprove), v))
	})
}

// StatusApproveContainsFold applies the ContainsFold predicate on the "StatusApprove" field.
func StatusApproveContainsFold(v string) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatusApprove), v))
	})
}

// HasTransactionhistory applies the HasEdge predicate on the "transactionhistory" edge.
func HasTransactionhistory() predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionhistoryTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionhistoryTable, TransactionhistoryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTransactionhistoryWith applies the HasEdge predicate on the "transactionhistory" edge with a given conditions (other predicates).
func HasTransactionhistoryWith(preds ...predicate.Transactionfactorhistory) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TransactionhistoryInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TransactionhistoryTable, TransactionhistoryColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Transactionfactor) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Transactionfactor) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
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
func Not(p predicate.Transactionfactor) predicate.Transactionfactor {
	return predicate.Transactionfactor(func(s *sql.Selector) {
		p(s.Not())
	})
}
