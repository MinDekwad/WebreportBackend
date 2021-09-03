package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Configpoint holds the schema definition for the Configpoint entity.
type Configpoint struct {
	ent.Schema
}

// Fields of the Configpoint.
func (Configpoint) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "config_point"},
	}
}

// Fields of the Configpoint.
// func (Configpoint) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Configpoint.
func (Configpoint) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("TransactionName").MaxLen(255).Optional().StorageKey("TransactionName"),
		field.String("TransactionType").MaxLen(255).Optional().Nillable().StorageKey("TransactionType"),
		field.String("PaymentChannel").MaxLen(255).Optional().Nillable().StorageKey("PaymentChannel"),
		field.String("PaymentType").MaxLen(255).Optional().Nillable().StorageKey("PaymentType"),
		field.String("DummyWallet").MaxLen(255).Optional().Nillable().StorageKey("DummyWallet"),
		field.Int("Amount").Optional().StorageKey("Amount"),
		field.Int("Point").Optional().StorageKey("Point"),
		field.Int("Expire").Optional().StorageKey("Expire"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("UpdateDate"),
		field.Time("ExpireDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("ExpireDate"),
		//field.Int("StatusTransaction").Optional().StorageKey("StatusTransaction"),
		field.String("StatusTransaction").MaxLen(10).Optional().Nillable().StorageKey("StatusTransaction"),
	}
}

// Edges of the Configpoint.
func (Configpoint) Edges() []ent.Edge {
	return nil
}
