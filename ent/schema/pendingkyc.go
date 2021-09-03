package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pendingkyc holds the schema definition for the Pendingkyc entity.
type Pendingkyc struct {
	ent.Schema
}

// Fields of the Pendingkyc.
func (Pendingkyc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pending_kyc"},
	}
}

// Fields of the Pendingkyc.
func (Pendingkyc) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("WalletID").MaxLen(45).Optional().StorageKey("WalletID"),
		field.String("Name").MaxLen(255).Optional().StorageKey("Name"),
		field.String("AgentID").MaxLen(45).Optional().Nillable().StorageKey("AgentID"),
		field.String("AgentNameLastname").MaxLen(45).Optional().Nillable().StorageKey("AgentNameLastname"),
		field.String("KYCDate").MaxLen(45).Optional().Nillable().StorageKey("KYCDate"),
		field.Time("DateGen").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateGen"),
		field.Bool("StatusGen").Default(false).Optional().StorageKey("StatusGen"),
		field.Int("Point").Default(0).Optional().StorageKey("Point"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Fields of the Pendingkyc.
// func (Pendingkyc) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Pendingkyc.
func (Pendingkyc) Edges() []ent.Edge {
	return nil
}
