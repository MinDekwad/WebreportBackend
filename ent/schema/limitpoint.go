package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Limitpoint holds the schema definition for the Limitpoint entity.
type Limitpoint struct {
	ent.Schema
}

// Fields of the Limitpoint.
func (Limitpoint) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "limit_point"},
	}
}

// Fields of the Limitpoint.
// func (Limitpoint) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Limitpoint.
func (Limitpoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("LimitPoint").Optional().StorageKey("LimitPoint"),
	}
}

// Edges of the Limitpoint.
func (Limitpoint) Edges() []ent.Edge {
	return nil
}
