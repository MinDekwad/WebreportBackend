package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pendingloanbinding holds the schema definition for the Pendingloanbinding entity.
type Pendingloanbinding struct {
	ent.Schema
}

// Fields of the Pendingloanbinding.
func (Pendingloanbinding) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pending_loan_binding"},
	}
}

func (Pendingloanbinding) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("WalletID").MaxLen(45).Optional().StorageKey("WalletID"),
		field.String("NameLB").MaxLen(255).Optional().StorageKey("NameLB"),
		field.Bool("StatusGenLB").Default(false).Optional().StorageKey("StatusGenLB"),
		field.Int("PointLB").Default(0).Optional().StorageKey("PointLB"),
		field.Time("DateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		field.Time("DateGenLB").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateGenLB"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
		field.String("CAWalletID").MaxLen(255).Optional().Nillable().StorageKey("CAWalletID"),
		field.String("CAPort").MaxLen(255).Optional().Nillable().StorageKey("CAPort"),
		field.String("MainBranch").MaxLen(255).Optional().Nillable().StorageKey("MainBranch"),
	}
}

// Fields of the Pendingloanbinding.
// func (Pendingloanbinding) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Pendingloanbinding.
func (Pendingloanbinding) Edges() []ent.Edge {
	return nil
}
