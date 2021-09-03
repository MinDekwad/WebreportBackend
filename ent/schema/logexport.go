package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Logexport holds the schema definition for the Logexport entity.
type Logexport struct {
	ent.Schema
}

func (Logexport) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "log_export"},
	}
}

// Fields of the Logexport.
// func (Logexport) Fields() []ent.Field {
// 	return nil
// }

func (Logexport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		//field.Int("UserID").StorageKey("UserID"),
		field.String("UserName").MaxLen(255).Optional().StorageKey("UserName"),
		field.String("FileName").MaxLen(255).Optional().StorageKey("FileName"),
		field.Time("ExportDate").SchemaType(map[string]string{dialect.MySQL: "ExportDate"}).Optional().StorageKey("ExportDate"),
	}
}

// Edges of the Logexport.
func (Logexport) Edges() []ent.Edge {
	return nil
}
