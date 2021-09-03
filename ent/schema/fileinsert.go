package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Fileinsert holds the schema definition for the Fileinsert entity.
type Fileinsert struct {
	ent.Schema
}

// Annotations of the Fileinsert.
func (Fileinsert) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "fileinsert"},
	}
}

// Fields of the Fileinsert.
func (Fileinsert) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("name").MaxLen(150).Optional().Nillable().StorageKey("name"),
		field.Time("importdate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("importdate"),
		field.Bool("IsSuccess").Default(false).Optional().StorageKey("IsSuccess"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("UpdateDate"),
	}
}

// Edges of the Fileinsert.
func (Fileinsert) Edges() []ent.Edge {
	return nil
}
