package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transactionfactor holds the schema definition for the Transactionfactor entity.
type Transactionfactor struct {
	ent.Schema
}

// Annotations of the Transactionfactor.
func (Transactionfactor) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "transaction_factor"},
	}
}

// Fields of the Transactionfactor.
func (Transactionfactor) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").StorageKey("uid"),
		field.Int("id").Unique().Optional().Immutable().StorageKey("id"),
		field.String("TransactionFactorName").MaxLen(255).Optional().StorageKey("TransactionFactorName"),
		field.String("TransactionType").MaxLen(255).Optional().StorageKey("TransactionType"),
		field.String("PaymentChannel").MaxLen(255).Optional().StorageKey("PaymentChannel"),
		field.String("PaymentType").MaxLen(255).Optional().StorageKey("PaymentType"),
		field.Int("NumDay").Optional().StorageKey("NumDay"),
		field.String("Date").MaxLen(10).Optional().StorageKey("Date"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("UpdateDate"),
		field.String("StatusApprove").MaxLen(10).Optional().StorageKey("StatusApprove"),
		//field.Int("FactorItemID").Optional().StorageKey("FactorItemID"),
		// field.Float("Min").Optional().Nillable().StorageKey("Min"),
		// field.Float("Max").Optional().Nillable().StorageKey("Max"),
		// field.Int("Rank").Optional().StorageKey("Rank"),
	}
}

// Fields of the Transactionfactor.
// func (Transactionfactor) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Transactionfactor.
// func (Transactionfactor) Edges() []ent.Edge {
// 	return nil
// }

func (Transactionfactor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("transactionhistory", Transactionfactorhistory.Type).StorageKey(edge.Column("TransactionFactorID")),
	}
}
