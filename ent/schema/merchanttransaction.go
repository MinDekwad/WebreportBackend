package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// MerchantTransaction holds the schema definition for the MerchantTransaction entity.
type MerchantTransaction struct {
	ent.Schema
}

// Annotations of the Merchant Transaction.
func (MerchantTransaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "merchant_transaction"},
	}
}

// Fields of the MerchantTransaction.
// func (MerchantTransaction) Fields() []ent.Field {
// 	return nil
// }

// Fields of the MerchantTransaction.
func (MerchantTransaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("transaction_id").MaxLen(45).Optional().Nillable().StorageKey("TransactionID"),
		field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		field.Float("amount").Optional().Nillable().StorageKey("Amount"),
		field.String("PaymentType").MaxLen(45).Optional().Nillable().StorageKey("PaymentType"),
		field.String("PaymentChannel").MaxLen(45).Optional().Nillable().StorageKey("PaymentChannel"),
		field.String("Status").MaxLen(45).Optional().Nillable().StorageKey("Status"),
		field.String("MerchantID").MaxLen(45).Optional().Nillable().StorageKey("MerchantID"),
		field.String("TerminalID").MaxLen(45).Optional().Nillable().StorageKey("TerminalID"),
		field.String("MerchantFullName").MaxLen(255).Optional().Nillable().StorageKey("MerchantFullName"),
		field.String("FromAccount").MaxLen(45).Optional().Nillable().StorageKey("FromAccount"),
		//field.String("FromName").MaxLen(45).Optional().Nillable().StorageKey("FromName"),
		field.String("SettlementAccount").MaxLen(45).Optional().Nillable().StorageKey("SettlementAccount"),
		field.Float("MDR_FEETHB").Optional().Nillable().StorageKey("MDR_FEETHB"),
		field.Float("TransactionFEETHB").Optional().Nillable().StorageKey("TransactionFEETHB"),
		field.Float("TotalFeeincVAT").Optional().Nillable().StorageKey("TotalFeeincVAT"),
		field.Float("VATTHB").Optional().Nillable().StorageKey("VATTHB"),
		field.Float("TotalFeeExcVAT").Optional().Nillable().StorageKey("TotalFeeExcVAT"),
		field.Float("WHTTHB").Optional().Nillable().StorageKey("WHTTHB"),
		field.Float("NetPayTHB").Optional().Nillable().StorageKey("NetPayTHB"),
		field.String("TransactionType").MaxLen(45).Optional().Nillable().StorageKey("TransactionType"),
		field.String("BankCode").MaxLen(45).Optional().Nillable().StorageKey("BankCode"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Edges of the MerchantTransaction.
func (MerchantTransaction) Edges() []ent.Edge {
	return nil
}
