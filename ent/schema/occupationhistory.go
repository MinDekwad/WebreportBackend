package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Occupationhistory holds the schema definition for the Occupationhistory entity.
type Occupationhistory struct {
	ent.Schema
}

// Annotations of the Occupationhistory.
func (Occupationhistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "occupation_history"},
	}
}

// Fields of the Occupationhistory.
func (Occupationhistory) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("WalletID").MaxLen(45).StorageKey("WalletID"),
		field.String("OccupationName").MaxLen(255).Optional().StorageKey("OccupationName"),
		field.Int("RankOccupation").Optional().StorageKey("RankOccupation"),
		//field.String("DateCalRank").MaxLen(10).Optional().StorageKey("DateCalRank"),
		field.Time("DateCalRank").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateCalRank"),
	}
}

// Fields of the Occupationhistory.
// func (Occupationhistory) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Occupationhistory.
func (Occupationhistory) Edges() []ent.Edge {
	return nil
}
