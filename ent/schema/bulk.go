package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Bulk holds the schema definition for the Bulk entity.
type Bulk struct {
	ent.Schema
}

// Fields of the Bulk.
func (Bulk) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "bulktransaction"},
	}
}

// Fields of the Bulk.
// func (Bulk) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Bulk.
func (Bulk) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		//field.String("KYCDate").MaxLen(45).Optional().Nillable().StorageKey("KYCDate"),
		//field.Time("KYCDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("KYCDate"),
		//field.String("KYCTime").MaxLen(45).Optional().Nillable().StorageKey("KYCTime"),
		field.Float("bulkCreditSameday").Optional().Nillable().StorageKey("BulkCreditSameday"),
		field.Float("bulkCreditSamedayFee").Optional().Nillable().StorageKey("BulkCreditSamedayFee"),
		field.Float("transfertobankaccount").Optional().Nillable().StorageKey("Transfertobankaccount"),
		field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
	}
}

// Edges of the Bulk.
func (Bulk) Edges() []ent.Edge {
	return nil
}
