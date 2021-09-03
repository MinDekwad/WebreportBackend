package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Fileimport holds the schema definition for the Fileimport entity.
type Fileimport struct {
	ent.Schema
}

// Annotations of the Fileimport.
func (Fileimport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "fileimport"},
	}
}

// Fields of the Fileimport.
func (Fileimport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("filename").MaxLen(150).Optional().Nillable().StorageKey("filename"),
		field.String("filetype").MaxLen(45).Optional().Nillable().StorageKey("filetype"),
		field.Time("importdate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("importdate"),
		field.String("status").MaxLen(10).Optional().Nillable().StorageKey("status"),
	}
}

// Fields of the Fileimport.
// func (Fileimport) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Fileimport.
func (Fileimport) Edges() []ent.Edge {
	return nil
}
