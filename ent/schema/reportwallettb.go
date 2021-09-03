package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Reportwallettb holds the schema definition for the Reportwallettb entity.
type Reportwallettb struct {
	ent.Schema
}

// Annotations of the Reportwallettb.
func (Reportwallettb) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "report_wallet"},
	}
}

// Fields of the Reportwallettb.
func (Reportwallettb) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("walletid").MaxLen(45).StorageKey("walletid"),
		field.String("WalletTypeName").MaxLen(45).Optional().Nillable().StorageKey("WalletTypeName"),
		field.String("WalletPhoneno").MaxLen(45).Optional().Nillable().StorageKey("WalletPhoneno"),
		field.String("WalletName").MaxLen(255).Optional().Nillable().StorageKey("WalletName"),
		field.String("CitizenId").MaxLen(13).Optional().Nillable().StorageKey("CitizenId"), // fix maxlen 13 ถ้าไม่กรอกให้เป็นค่าว่าง
		field.String("Status").MaxLen(45).Optional().Nillable().StorageKey("Status"),
		field.Time("RegisterDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("RegisterDate"),
		field.Int("GroupUser").Optional().Nillable().StorageKey("GroupUser"),
		field.String("UserAgent").MaxLen(45).Optional().Nillable().StorageKey("UserAgent"),
		field.Time("KYC_Date").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("KYC_Date"),
		field.String("ATMCard").MaxLen(45).Optional().Nillable().StorageKey("ATMCard"),
		field.String("AccountNo").MaxLen(45).Optional().Nillable().StorageKey("AccountNo"),
		field.String("AddressDetail").Optional().Nillable().StorageKey("AddressDetail"),
		field.String("Street").Optional().Nillable().StorageKey("Street"),
		field.String("District").Optional().Nillable().StorageKey("District"),
		field.String("SubDistrict").Optional().Nillable().StorageKey("SubDistrict"),
		field.String("Province").MaxLen(45).Optional().Nillable().StorageKey("Province"),
		field.String("PostalCode").MaxLen(45).Optional().Nillable().StorageKey("PostalCode"),
		field.String("isKYC").MaxLen(1).Optional().Nillable().StorageKey("isKYC"),
		field.Time("UpdateDate").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("UpdateDate"),
	}
}

// Fields of the Reportwallettb.
// func (Reportwallettb) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Reportwallettb.
func (Reportwallettb) Edges() []ent.Edge {
	return nil
}
