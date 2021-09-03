package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pointpendingkyctransaction holds the schema definition for the Pointpendingkyctransaction entity.
type Pointpendingkyctransaction struct {
	ent.Schema
}

// Fields of the Pointpendingkyctransaction.
func (Pointpendingkyctransaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "point_pending_kyc_transaction"},
	}
}

func (Pointpendingkyctransaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("WalletId").MaxLen(45).Optional().StorageKey("WalletId"),
		field.Int("Point").Default(0).Optional().Nillable().StorageKey("Point"),
		field.Time("DateExport").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateExport"),
		field.Time("DateGen").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateGen"),
		field.String("Note").MaxLen(255).Optional().Nillable().StorageKey("Note"),
		field.Bool("StatusExport").Default(false).Optional().StorageKey("StatusExport"),
		field.String("KYCDate").MaxLen(45).Optional().Nillable().StorageKey("KYCDate"),
	}
}

// Fields of the Pointpendingkyctransaction.
// func (Pointpendingkyctransaction) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Pointpendingkyctransaction.
func (Pointpendingkyctransaction) Edges() []ent.Edge {
	return nil
}
