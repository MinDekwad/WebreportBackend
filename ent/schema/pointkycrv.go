package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pointkycrv holds the schema definition for the Pointkycrv entity.
type Pointkycrv struct {
	ent.Schema
}

// Fields of the Pointtransaction.
func (Pointkycrv) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "point_kyc_rv"},
	}
}

// Fields of the Pointkycrv.
// func (Pointkycrv) Fields() []ent.Field {
// 	return nil
// }

func (Pointkycrv) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Time("DateTimeGen").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("DateTimeGen"),
		field.String("WalletID").MaxLen(45).Optional().StorageKey("WalletID"),
		field.String("StatusGen").MaxLen(1).Optional().StorageKey("StatusGen"),
		field.Time("KYCDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("KYCDate"),
		field.Time("RVDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("RVDate"),
		field.String("Type").MaxLen(3).Optional().StorageKey("Type"),
		field.Int("Point").Optional().StorageKey("Point"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Edges of the Pointkycrv.
func (Pointkycrv) Edges() []ent.Edge {
	return nil
}
