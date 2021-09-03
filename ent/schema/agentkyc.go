package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Agentkyc holds the schema definition for the Agentkyc entity.
type Agentkyc struct {
	ent.Schema
}

// Fields of the Agentkyc.
func (Agentkyc) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "agent_kyc"},
	}
}

// Fields of the Consumer.
func (Agentkyc) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("KYCDate").MaxLen(45).Optional().Nillable().StorageKey("KYCDate"),
		//field.Time("KYCDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("KYCDate"),
		field.String("KYCTime").MaxLen(45).Optional().Nillable().StorageKey("KYCTime"),
		field.String("AgentID").MaxLen(45).Optional().Nillable().StorageKey("AgentID"),
		field.String("Agentemail").MaxLen(45).Optional().Nillable().StorageKey("Agentemail"),
		field.String("AgentNameLastname").MaxLen(45).Optional().Nillable().StorageKey("AgentNameLastname"),
		field.String("KYCStatus").MaxLen(45).Optional().Nillable().StorageKey("KYCStatus"),
		field.String("Consumerwalletid").MaxLen(45).Optional().Nillable().StorageKey("Consumerwalletid"),
		field.String("KYCRespond").MaxLen(45).Optional().Nillable().StorageKey("KYCRespond"),
		field.String("DOPARespond").MaxLen(45).Optional().Nillable().StorageKey("DOPARespond"),
		field.String("AgentType").MaxLen(45).Optional().Nillable().StorageKey("AgentType"),
		//field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Edges of the Agentkyc.
func (Agentkyc) Edges() []ent.Edge {
	return nil
}
