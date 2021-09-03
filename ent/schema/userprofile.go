package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Userprofile holds the schema definition for the Userprofile entity.
type Userprofile struct {
	ent.Schema
}

// Annotations of the Userprofile.
func (Userprofile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_profile"},
	}
}

// Fields of the Userprofile.
func (Userprofile) Fields() []ent.Field {
	return []ent.Field{
		//field.Int("id").StorageKey("uid"),
		field.Int("id").Unique().Immutable().StorageKey("id"),
		field.String("UserId").MaxLen(255).Optional().StorageKey("UserId"),
		field.String("Firstname").MaxLen(255).Optional().StorageKey("Firstname"),
		field.String("Lastname").MaxLen(255).Optional().StorageKey("Lastname"),
		field.String("PhoneNo").MaxLen(15).Optional().StorageKey("PhoneNo"),
		field.String("Email").MaxLen(255).Optional().StorageKey("Email"),
		field.String("CitizenId").MaxLen(45).Optional().StorageKey("CitizenId"),
		field.String("BirthDate").MaxLen(45).Optional().StorageKey("BirthDate"),
		field.String("Gender").MaxLen(10).Optional().StorageKey("Gender"),
		// field.String("Address").MaxLen(255).Optional().StorageKey("Address"),
		// field.Time("CreateAt").SchemaType(map[string]string{dialect.MySQL: "CreateAt"}).Optional().Nillable().StorageKey("CreateAt"),
		// field.Time("UpdateAt").SchemaType(map[string]string{dialect.MySQL: "CreateAt"}).Optional().Nillable().StorageKey("UpdateAt"),
		// field.String("CreateBy").MaxLen(255).Optional().StorageKey("CreateBy"),
		// field.String("CreaterType").MaxLen(255).Optional().StorageKey("CreaterType"),
		// field.String("AddressId").MaxLen(255).Optional().StorageKey("AddressId"),
		field.String("BusinessAddress").MaxLen(255).Optional().StorageKey("BusinessAddress"),
		field.Int("OccupationId").Optional().StorageKey("OccupationId"),
		field.Int("FileimportID").Optional().Nillable().StorageKey("FileimportID"),
	}
}

// Fields of the Userprofile.
// func (Userprofile) Fields() []ent.Field {
// 	return nil
// }

// Edges of the Userprofile.
func (Userprofile) Edges() []ent.Edge {
	return nil
}
