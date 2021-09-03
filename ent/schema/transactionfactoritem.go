package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Transactionfactoritem holds the schema definition for the Transactionfactoritem entity.
type Transactionfactoritem struct {
	ent.Schema
}

// Annotations of the Transactionfactoritem.
func (Transactionfactoritem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "transaction_factor_item"},
	}
}

// Fields of the Transactionfactoritem.
func (Transactionfactoritem) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").StorageKey("uid"),
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("TransactionFactorID").Optional().StorageKey("TransactionFactorID"),
		field.Float("Min").Optional().StorageKey("Min"),
		field.Float("Max").Optional().StorageKey("Max"),
		field.Int("Rank").Optional().StorageKey("Rank"),
	}
}

// Fields of the Transactionfactoritem.
// func (Transactionfactoritem) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Transactionfactoritem.
func (Transactionfactoritem) Edges() []ent.Edge {
	return nil
}
