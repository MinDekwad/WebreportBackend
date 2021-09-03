package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Watchlisthistory holds the schema definition for the Watchlisthistory entity.
type Watchlisthistory struct {
	ent.Schema
}

// Annotations of the Watchlisthistory.
func (Watchlisthistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "watchlist_history"},
	}
}

// Fields of the Watchlisthistory.
func (Watchlisthistory) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("Name").MaxLen(255).StorageKey("Name"),
		field.String("TaxID").MaxLen(13).StorageKey("TaxID"),
		field.String("TypeName").MaxLen(255).StorageKey("TypeName"),
		field.Int("RankWatchlist").Optional().StorageKey("RankWatchlist"),
		//field.String("DateCalRank").MaxLen(10).Optional().StorageKey("DateCalRank"),
		field.Int("StatusDel").Optional().StorageKey("StatusDel"),
		field.Time("DateCalRank").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateCalRank"),
	}
}

// Fields of the Watchlisthistory.
// func (Watchlisthistory) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Watchlisthistory.
func (Watchlisthistory) Edges() []ent.Edge {
	return nil
}
