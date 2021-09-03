package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// ReportWallet holds the schema definition for the ReportWallet entity.
type ReportWallet struct {
	ent.Schema
}

// Annotations of the Wallet.
func (ReportWallet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "wallet_daily"},
	}
}

// Fields of the ReportWallet.
func (ReportWallet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Unique().Immutable().StorageKey("uid"),
		field.String("walletid").MaxLen(45).StorageKey("WalletID"),
		field.String("WalletTypeName").MaxLen(45).Optional().Nillable().StorageKey("WalletTypeName"),
		field.String("WalletPhoneno").MaxLen(45).Optional().Nillable().StorageKey("WalletPhoneno"),
		field.String("WalletName").MaxLen(255).Optional().Nillable().StorageKey("WalletName"),
		field.String("CitizenId").MaxLen(13).Optional().Nillable().StorageKey("CitizenId"), // fix maxlen 13 ถ้าไม่กรอกให้เป็นค่าว่าง
		field.String("Status").MaxLen(45).Optional().Nillable().StorageKey("Status"),
		field.Time("dateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		field.Float("Balance").Optional().Nillable().StorageKey("Balance"),
		field.String("Email").MaxLen(255).Optional().Nillable().StorageKey("Email"),
		field.String("IsForgetPin").MaxLen(45).Optional().Nillable().StorageKey("IsForgetPin"),
		field.String("ATMCard").MaxLen(45).Optional().Nillable().StorageKey("ATMCard"),
		field.String("AccountNo").MaxLen(45).Optional().Nillable().StorageKey("AccountNo"),
		field.String("AddressDetail").Optional().Nillable().StorageKey("AddressDetail"),
		field.String("Street").MaxLen(255).Optional().Nillable().StorageKey("Street"),
		field.String("District").MaxLen(255).Optional().Nillable().StorageKey("District"),
		field.String("SubDistrict").MaxLen(255).Optional().Nillable().StorageKey("SubDistrict"),
		field.String("Province").MaxLen(45).Optional().Nillable().StorageKey("Province"),
		field.String("PostalCode").MaxLen(45).Optional().Nillable().StorageKey("PostalCode"),
		field.Time("RegisterDateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("RegisterDateTime"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),

		// field.Int("id").Unique().Immutable().StorageKey("uid"),
		// field.String("WalletID").MaxLen(45),
		// field.String("WalletTypeName").MaxLen(45).Optional().Nillable().StorageKey("WalletTypeName"),
		// field.String("WalletPhoneno").MaxLen(45).Optional().Nillable().StorageKey("WalletPhoneno"),
		// field.String("WalletName").MaxLen(255).Optional().Nillable().StorageKey("WalletName"),
		// field.String("CitizenId").MaxLen(100).Optional().Nillable().StorageKey("CitizenId"),
		// field.String("Status").MaxLen(45).Optional().Nillable().StorageKey("Status"),
		// field.Time("DateTime").SchemaType(map[string]string{dialect.MySQL: "datetime"}).Optional().Nillable().StorageKey("DateTime"),
		// field.Float("Balance").Optional().Nillable().StorageKey("Balance"),
	}
}

// Edges of the ReportWallet.
func (ReportWallet) Edges() []ent.Edge {
	return nil
}
