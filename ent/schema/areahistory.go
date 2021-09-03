package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Areahistory holds the schema definition for the Areahistory entity.
type Areahistory struct {
	ent.Schema
}

// Annotations of the Areahistory.
func (Areahistory) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "area_history"},
	}
}

// Fields of the Areahistory.
func (Areahistory) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("WalletID").MaxLen(45).StorageKey("WalletID"),
		field.String("ProvinceNameTH").MaxLen(255).Optional().StorageKey("ProvinceNameTH"),
		field.String("DistrictNameTH").MaxLen(255).Optional().StorageKey("DistrictNameTH"),
		field.String("SubDistrict").MaxLen(255).Nillable().Optional().StorageKey("SubDistrict"),
		field.Int("RankArea").Optional().StorageKey("RankArea"),
		//field.String("DateCalRank").MaxLen(10).Optional().StorageKey("DateCalRank"),
		field.Time("DateCalRank").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateCalRank"),
	}
}

// Fields of the Areahistory.
// func (Areahistory) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Areahistory.
func (Areahistory) Edges() []ent.Edge {
	return nil
}
