package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Transactionfactoritemtmp holds the schema definition for the Transactionfactoritemtmp entity.
type Transactionfactoritemtmp struct {
	ent.Schema
}

// Annotations of the Transactionfactoritemtmp.
func (Transactionfactoritemtmp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "transaction_factor_item_tmp"},
	}
}

// Fields of the Transactionfactor.
func (Transactionfactoritemtmp) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").StorageKey("uid"),
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Float("Min").Optional().StorageKey("Min"),
		field.Float("Max").Optional().StorageKey("Max"),
		field.Int("Rank").Optional().StorageKey("Rank"),
	}
}

// Fields of the Transactionfactoritemtmp.
// func (Transactionfactoritemtmp) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Transactionfactoritemtmp.
func (Transactionfactoritemtmp) Edges() []ent.Edge {
	return nil
}
