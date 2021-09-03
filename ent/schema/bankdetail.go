package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Bankdetail holds the schema definition for the Bankdetail entity.
type Bankdetail struct {
	ent.Schema
}

// Fields of the Bankdetail.
func (Bankdetail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "bank_detail"},
	}
}

// Fields of the Bankdetail.
func (Bankdetail) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("Bank_AccountNo").MaxLen(45).Optional().Nillable().StorageKey("Bank_AccountNo"),
		field.String("Bank_Name").MaxLen(45).Optional().Nillable().StorageKey("Bank_Name"),
		field.String("Bank_AccountName").MaxLen(45).Optional().Nillable().StorageKey("Bank_AccountName"),
	}
}

// Edges of the Bankdetail.
func (Bankdetail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("statements", StatementEndingBalanc.Type).StorageKey(edge.Column("Bank_UID")),
	}
}
