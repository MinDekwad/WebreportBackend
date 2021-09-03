package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pointpendinglbtransaction holds the schema definition for the Pointpendinglbtransaction entity.
type Pointpendinglbtransaction struct {
	ent.Schema
}

// Fields of the Pointpendinglbtransaction.
func (Pointpendinglbtransaction) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "point_pending_lb_transaction"},
	}
}

func (Pointpendinglbtransaction) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("WalletID").MaxLen(45).Optional().StorageKey("WalletID"),
		field.Int("PointLB").Default(0).Optional().Nillable().StorageKey("PointLB"),
		field.Time("DateExportLB").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateExportLB"),
		field.Time("DateGenLB").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateGenLB"),
		field.String("NoteLB").MaxLen(255).Optional().Nillable().StorageKey("NoteLB"),
		field.Bool("StatusExportLB").Default(false).Optional().StorageKey("StatusExportLB"),
		field.Time("LBDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("LBDate"),
	}
}

// Fields of the Pointpendinglbtransaction.
// func (Pointpendinglbtransaction) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Pointpendinglbtransaction.
func (Pointpendinglbtransaction) Edges() []ent.Edge {
	return nil
}
