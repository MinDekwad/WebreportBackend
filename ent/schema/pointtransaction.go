package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pointtransaction holds the schema definition for the Pointtransaction entity.
type Pointtransaction struct {
	ent.Schema
}

// Fields of the Pointtransaction.
func (Pointtransaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "point_transaction"},
	}
}

// Fields of the Pointtransaction.
// func (Pointtransaction) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Pointtransaction.
func (Pointtransaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Time("Date").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("Date"),
		field.String("WalletID").MaxLen(225).Optional().StorageKey("WalletID"),
		field.String("TransactionName").MaxLen(255).Optional().StorageKey("TransactionName"),
		field.Int("Point").Optional().StorageKey("Point"),
		//field.Time("ExpireDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("ExpireDate"),
		field.Text("Reference").Optional().StorageKey("Reference"),
	}
}

// Edges of the Pointtransaction.
func (Pointtransaction) Edges() []ent.Edge {
	return nil
}
