package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Writelog holds the schema definition for the Writelog entity.
type Writelog struct {
	ent.Schema
}

// Fields of the Writelog.
// func (Writelog) Fields() []ent.Field {
// 	return nil
// }
func (Writelog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "log"},
	}
}

func (Writelog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("admin_id").StorageKey("admin_id"),
		field.String("resource").MaxLen(255).Optional().Nillable().StorageKey("resource"),
		field.String("actions").MaxLen(255).Optional().Nillable().StorageKey("actions"),
		field.Time("created_at").SchemaType(map[string]string{dialect.MySQL: "created_at"}).Optional().Nillable().StorageKey("created_at"),
	}
}

// Edges of the Writelog.
func (Writelog) Edges() []ent.Edge {
	return nil
}
