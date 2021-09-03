package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Ranking holds the schema definition for the Ranking entity.
type Ranking struct {
	ent.Schema
}

// Annotations of the Ranking.
func (Ranking) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "ranking"},
	}
}

// Fields of the Ranking.
func (Ranking) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").Unique().Immutable().StorageKey("id"),
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("WalletID").MaxLen(45).StorageKey("WalletID"),
		field.String("Name").MaxLen(255).Optional().StorageKey("Name"),
		field.String("TaxID").MaxLen(13).Optional().StorageKey("TaxID"), // fix maxlen 13 ถ้าไม่กรอกให้เป็นค่าว่าง
		field.String("ProvinceNameTH").MaxLen(255).Optional().StorageKey("ProvinceNameTH"),
		field.String("DistrictNameTH").MaxLen(255).Optional().StorageKey("DistrictNameTH"),
		field.String("DistrictNameEN").MaxLen(255).Optional().StorageKey("DistrictNameEN"),
		field.String("OccupationName").MaxLen(255).Optional().StorageKey("OccupationName"),
		field.Int("LastRank").Optional().StorageKey("LastRank"),
		field.Int("CurrentRank").Optional().StorageKey("CurrentRank"),
		field.String("StatusRanking").MaxLen(10).Optional().StorageKey("StatusRanking"),
		//field.Int("StateRank").Optional().StorageKey("StateRank"),
		field.String("LastDateCalRank").MaxLen(10).Optional().StorageKey("LastDateCalRank"),
		field.String("NextDateCalRank").MaxLen(10).Optional().StorageKey("NextDateCalRank"),
		field.Int("StateCal").Optional().StorageKey("StateCal"),
		field.String("ZipCode").MaxLen(5).Optional().Nillable().StorageKey("ZipCode"),
		field.Int("TransactionFactorRank").Optional().StorageKey("TransactionFactorRank"),
		field.Time("RegisDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("RegisDate"),
		field.String("SubDistrict").MaxLen(255).Optional().StorageKey("SubDistrict"),
		field.String("Phoneno").MaxLen(45).Optional().Nillable().StorageKey("Phoneno"),
		field.String("AddressDetail").MaxLen(255).Optional().Nillable().StorageKey("AddressDetail"),
		field.String("Street").MaxLen(255).Optional().Nillable().StorageKey("Street"),
	}
}

// Fields of the Ranking.
// func (Ranking) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Ranking.
func (Ranking) Edges() []ent.Edge {
	return nil
}
