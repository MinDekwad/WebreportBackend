package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Transactionfactorhistory holds the schema definition for the Transactionfactorhistory entity.
type Transactionfactorhistory struct {
	ent.Schema
}

// Annotations of the Transactionfactorhistory.
func (Transactionfactorhistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "transaction_factor_history"},
	}
}

// Fields of the Transactionfactorhistory.
func (Transactionfactorhistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("WalletID").MaxLen(45).StorageKey("WalletID"),
		//field.String("TransactionFactorName").MaxLen(255).Optional().StorageKey("TransactionFactorName"),
		//field.Int("TransactionFactorID").Optional().StorageKey("TransactionFactorID"),
		field.Int("RankTransactionFactor").Optional().StorageKey("RankTransactionFactor"),
		//field.String("DateCalRank").MaxLen(10).Optional().StorageKey("DateCalRank"),
		field.Time("DateCalRank").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateCalRank"),
	}
}

// Edges of the Transactionfactorhistory.
// func (Transactionfactorhistory) Edges() []ent.Edge {
// 	return nil
// }

// Edges of the StatementEndingBalanc.
func (Transactionfactorhistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Transactionfactor", Transactionfactor.Type).Ref("transactionhistory").Unique(),
	}
}
