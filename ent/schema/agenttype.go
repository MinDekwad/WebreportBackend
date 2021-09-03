package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Agenttype holds the schema definition for the Agenttype entity.
type Agenttype struct {
	ent.Schema
}

// Fields of the Agenttype.
func (Agenttype) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "agent_type"},
	}
}

// Fields of the Agenttype.
func (Agenttype) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").StorageKey("uid"),
		field.String("agentid").MaxLen(45).Optional().Nillable().StorageKey("agentid"),
		field.String("agentname").MaxLen(45).Optional().Nillable().StorageKey("agentname"),
		field.String("agenttype").MaxLen(45).Optional().Nillable().StorageKey("agenttype"),
	}
}

// Edges of the Agenttype.
func (Agenttype) Edges() []ent.Edge {
	return nil
}
