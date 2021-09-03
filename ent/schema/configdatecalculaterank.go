package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Configdatecalculaterank holds the schema definition for the Configdatecalculaterank entity.
type Configdatecalculaterank struct {
	ent.Schema
}

// Fields of the Configdatecalculaterank.
func (Configdatecalculaterank) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "config_date_calculate_rank"},
	}
}

// Fields of the Configdatecalculaterank.
// func (Configdatecalculaterank) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Configdatecalculaterank.
func (Configdatecalculaterank) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("Rank").MaxLen(1).Optional().StorageKey("Rank"),
		field.Int("NumDateCalculateRank").Optional().StorageKey("NumDateCalculateRank"),
		field.String("NumDateCalculateRankTmp").MaxLen(11).Optional().StorageKey("NumDateCalculateRankTmp"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("UpdateDate"),
	}
}

// Edges of the Configdatecalculaterank.
func (Configdatecalculaterank) Edges() []ent.Edge {
	return nil
}
