package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Consumer holds the schema definition for the Consumer entity.
type Consumer struct {
	ent.Schema
}

// Annotations of the User.
func (Consumer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "consumer_transaction"},
	}
}

// Fields of the Consumer.
func (Consumer) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").StorageKey("uid"),
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		//field.String("TransactionID").MaxLen(45),
		field.String("transaction_id").MaxLen(45).Optional().Nillable().StorageKey("TransactionID"),
		field.String("TransactionStatus").MaxLen(45).Optional().Nillable().StorageKey("TransactionStatus"),
		field.String("TransactionType").MaxLen(45).Optional().Nillable().StorageKey("TransactionType"),
		field.String("PaymentChannel").MaxLen(45).Optional().Nillable().StorageKey("PaymentChannel"),
		field.String("PaymentType").MaxLen(45).Optional().Nillable().StorageKey("PaymentType"),
		field.String("TypeCode").MaxLen(45).Optional().Nillable().StorageKey("TypeCode"),
		field.String("ApprovalCode").MaxLen(45).Optional().Nillable().StorageKey("ApprovalCode"),
		field.String("BillerID").MaxLen(45).Optional().Nillable().StorageKey("BillerID"),
		field.String("ref1").MaxLen(45).Optional().Nillable().StorageKey("ref1"),
		field.String("ref2").MaxLen(45).Optional().Nillable().StorageKey("ref2"),
		field.String("ref3").MaxLen(45).Optional().Nillable().StorageKey("ref3"),
		field.Float("amount").Optional().Nillable().StorageKey("Amount"),
		field.Float("fee").Optional().Nillable().StorageKey("fee"),
		field.Float("total").Optional().Nillable().StorageKey("total"),
		field.String("FromReference").MaxLen(45).Optional().Nillable().StorageKey("FromReference"),
		field.String("FromPhoneNo").MaxLen(45).Optional().Nillable().StorageKey("FromPhoneNo"),
		field.String("FromName").MaxLen(250).Optional().Nillable().StorageKey("FromName"),
		field.String("ToAccount").MaxLen(45).Optional().Nillable().StorageKey("ToAccount"),
		field.String("ToAccountPhoneNo").MaxLen(45).Optional().Nillable().StorageKey("ToAccountPhoneNo"),
		field.String("ToAccountName").MaxLen(250).Optional().Nillable().StorageKey("ToAccountName"),
		field.String("BankCode").MaxLen(45).Optional().Nillable().StorageKey("BankCode"),
		field.String("TerminalId").MaxLen(45).Optional().Nillable().StorageKey("TerminalId"),
		field.String("TerminalType").MaxLen(45).Optional().Nillable().StorageKey("TerminalType"),
		field.String("ToAccount105").MaxLen(45).Optional().Nillable().StorageKey("ToAccount105"),
		field.String("FromReference105").MaxLen(45).Optional().Nillable().StorageKey("FromReference105"),
		field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		field.String("PartnerRef").MaxLen(50).Optional().Nillable().StorageKey("PartnerRef"),
		field.String("ResponseCode").MaxLen(10).Optional().Nillable().StorageKey("ResponseCode"),
		field.String("ResponseDescription").MaxLen(255).Optional().Nillable().StorageKey("ResponseDescription"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Edges of the Consumer.
func (Consumer) Edges() []ent.Edge {
	return nil
}
