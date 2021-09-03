package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// StatementEndingBalanc holds the schema definition for the StatementEndingBalanc entity.
type StatementEndingBalanc struct {
	ent.Schema
}

// Fields of the StatementEndingBalanc.
func (StatementEndingBalanc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "statement_ending_balance"},
	}
}

// Fields of the StatementEndingBalanc.
func (StatementEndingBalanc) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		// field.Int("BankUID").StorageKey("BankUID"),
		field.Float("Statement_Balance").Optional().Nillable().StorageKey("Statement_Balance"),
		//field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("Statement_Date"),
		field.Time("Statement_Date").SchemaType(map[string]string{dialect.MySQL: "Statement_Date"}).Optional().Nillable().StorageKey("Statement_Date"),
	}
}

// Edges of the StatementEndingBalanc.
func (StatementEndingBalanc) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
		// and reference it to the "cars" edge (in User schema)
		// explicitly using the `Ref` method.
		edge.From("bank", Bankdetail.Type).
			Ref("statements").
			// setting the edge to unique, ensure
			// that a car can have only one owner.
			Unique(),
	}
}
