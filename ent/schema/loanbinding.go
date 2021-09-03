package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Loanbinding holds the schema definition for the Loanbinding entity.
type Loanbinding struct {
	ent.Schema
}

// Fields of the Loanbinding.
func (Loanbinding) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "wallet_loanbinding"},
	}
}

// Fields of the Loanbinding.
// func (Loanbinding) Fields() []ent.Field {
// 	return nil
// }

// Fields of the Loanbinding.
func (Loanbinding) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("Status").MaxLen(45).Optional().Nillable().StorageKey("Status"),
		field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		field.String("WalletId").MaxLen(45).Optional().Nillable().StorageKey("WalletId"),
		field.String("AccountReference").MaxLen(150).Optional().Nillable().StorageKey("AccountReference"),
		field.String("LoanAccountNo").MaxLen(150).Optional().Nillable().StorageKey("LoanAccountNo"),
		field.Time("RequestDateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("RequestDateTime"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Edges of the Loanbinding.
func (Loanbinding) Edges() []ent.Edge {
	return nil
}
