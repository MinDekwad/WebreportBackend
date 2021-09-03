package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pointcsv holds the schema definition for the Pointcsv entity.
type Pointcsv struct {
	ent.Schema
}

// Fields of the Pointcsv.
func (Pointcsv) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "point_csv"},
	}
}

// Fields of the Pointcsv.
func (Pointcsv) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("WalletID").MaxLen(255).Optional().StorageKey("WalletID"),
		field.Time("CreateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("CreateDate"),
		field.Int("Adjustamount").Optional().StorageKey("Adjustamount"),
		field.String("Note").MaxLen(255).Optional().Nillable().StorageKey("Note"),
		// field.Int("PointTranID").Optional().StorageKey("PointTranID"),
		field.Time("PointTranDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("PointTranDate"),
		field.Int("ActionExport").Optional().StorageKey("ActionExport"),
		//field.Int("Num").Optional().StorageKey("Num"),
	}
}

// Fields of the Pointcsv.
// func (Pointcsv) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Pointcsv.
func (Pointcsv) Edges() []ent.Edge {
	return nil
}
