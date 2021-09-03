package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Watchlist holds the schema definition for the Watchlist entity.
type Watchlist struct {
	ent.Schema
}

// Annotations of the Watchlist.
func (Watchlist) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "watchlist"},
	}
}

// Fields of the Watchlist.
func (Watchlist) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("Name").MaxLen(255).StorageKey("Name"),
		field.String("TaxID").MaxLen(20).StorageKey("TaxID"),
		//field.Int("TypeID").Optional().StorageKey("TypeID"),
		field.Int("RankWatchlist").Optional().StorageKey("RankWatchlist"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
		field.Bool("IsDeleted").Default(false).Optional().StorageKey("IsDeleted"),
		field.Time("ImportDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("ImportDate"),
		field.String("UserUpload").MaxLen(255).Optional().Nillable().StorageKey("UserUpload"),
	}
}

// Fields of the Watchlist.
// func (Watchlist) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Watchlist.
func (Watchlist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("related", Watchlisttype.Type).Ref("watchlist").Unique(),
	}
}
