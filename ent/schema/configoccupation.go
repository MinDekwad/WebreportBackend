package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Configoccupation holds the schema definition for the Configoccupation entity.
type Configoccupation struct {
	ent.Schema
}

// Fields of the Configoccupation.
func (Configoccupation) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "config_occupation"},
	}
}

// Fields of the Configoccupation.
// func (Configoccupation) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Configoccupation.
func (Configoccupation) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("OccupationName").MaxLen(255).Optional().StorageKey("OccupationName"),
		field.String("Rank").MaxLen(1).Optional().StorageKey("Rank"),
		field.String("RankTmp").MaxLen(1).Optional().StorageKey("RankTmp"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("UpdateDate"),
		field.String("ApproveBy").MaxLen(255).Optional().Nillable().StorageKey("ApproveBy"),
		field.Time("ApproveDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("ApproveDate"),
	}
}

// Edges of the Configoccupation.
func (Configoccupation) Edges() []ent.Edge {
	return nil
}
