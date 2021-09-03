package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Watchlisttype holds the schema definition for the Watchlisttype entity.
type Watchlisttype struct {
	ent.Schema
}

// Annotations of the Watchlisttype.
func (Watchlisttype) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "watchlist_type"},
	}
}

// Fields of the Watchlisttype.
func (Watchlisttype) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("TypeID").
			//Unique().Positive().
			Optional().StorageKey("TypeID"),
		field.String("TypeName").MaxLen(255).StorageKey("TypeName"),
		field.String("TypeDescription").MaxLen(255).StorageKey("TypeDescription"),
	}
}

// Fields of the Watchlisttype.
// func (Watchlisttype) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Watchlisttype.
func (Watchlisttype) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("watchlist", Watchlist.Type).StorageKey(edge.Column("TypeID")),
	}
}
