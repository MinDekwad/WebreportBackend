package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Configarea holds the schema definition for the Configarea entity.
type Configarea struct {
	ent.Schema
}

// Fields of the Configarea.
func (Configarea) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "cofig_area"},
	}
}

// Fields of the Configarea.
func (Configarea) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("ProvinceID").Optional().StorageKey("ProvinceID"),
		field.String("ProvinceNameTH").MaxLen(255).Optional().StorageKey("ProvinceNameTH"),
		field.Int("DistrictID").Optional().StorageKey("DistrictID"),
		field.String("DistrictNameTH").MaxLen(255).Optional().StorageKey("DistrictNameTH"),
		field.String("DistrictNameEN").MaxLen(255).Optional().StorageKey("DistrictNameEN"),
		field.String("Rank").MaxLen(1).Optional().StorageKey("Rank"),
		field.String("RankTmp").MaxLen(1).Optional().StorageKey("RankTmp"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().StorageKey("UpdateDate"),
		field.String("ZipCode").MaxLen(1).Optional().StorageKey("ZipCode"),
		field.String("SubDistrictNameTH").MaxLen(1).Optional().StorageKey("SubDistrictNameTH"),
		field.String("ApproveBy").MaxLen(255).Optional().Nillable().StorageKey("ApproveBy"),
		field.Time("ApproveDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("ApproveDate"),
	}
}

// Fields of the Configarea.
// func (Configarea) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Configarea.
func (Configarea) Edges() []ent.Edge {
	return nil
}
