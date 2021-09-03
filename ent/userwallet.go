// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/userwallet"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Userwallet is the model entity for the Userwallet schema.
type Userwallet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Walletid holds the value of the "walletid" field.
	Walletid string `json:"walletid,omitempty"`
	// WalletTypeName holds the value of the "WalletTypeName" field.
	WalletTypeName *string `json:"WalletTypeName,omitempty"`
	// WalletPhoneno holds the value of the "WalletPhoneno" field.
	WalletPhoneno *string `json:"WalletPhoneno,omitempty"`
	// WalletName holds the value of the "WalletName" field.
	WalletName *string `json:"WalletName,omitempty"`
	// CitizenId holds the value of the "CitizenId" field.
	CitizenId *string `json:"CitizenId,omitempty"`
	// Status holds the value of the "Status" field.
	Status *string `json:"Status,omitempty"`
	// RegisterDate holds the value of the "RegisterDate" field.
	RegisterDate *time.Time `json:"RegisterDate,omitempty"`
	// GroupUser holds the value of the "GroupUser" field.
	GroupUser *int `json:"GroupUser,omitempty"`
	// UserAgent holds the value of the "UserAgent" field.
	UserAgent *string `json:"UserAgent,omitempty"`
	// KYCDate holds the value of the "KYC_Date" field.
	KYCDate *time.Time `json:"KYC_Date,omitempty"`
	// ATMCard holds the value of the "ATMCard" field.
	ATMCard *string `json:"ATMCard,omitempty"`
	// AccountNo holds the value of the "AccountNo" field.
	AccountNo *string `json:"AccountNo,omitempty"`
	// AddressDetail holds the value of the "AddressDetail" field.
	AddressDetail *string `json:"AddressDetail,omitempty"`
	// Street holds the value of the "Street" field.
	Street *string `json:"Street,omitempty"`
	// District holds the value of the "District" field.
	District *string `json:"District,omitempty"`
	// SubDistrict holds the value of the "SubDistrict" field.
	SubDistrict *string `json:"SubDistrict,omitempty"`
	// Province holds the value of the "Province" field.
	Province *string `json:"Province,omitempty"`
	// PostalCode holds the value of the "PostalCode" field.
	PostalCode *string `json:"PostalCode,omitempty"`
	// IsKYC holds the value of the "isKYC" field.
	IsKYC *string `json:"isKYC,omitempty"`
	// UpdateDate holds the value of the "UpdateDate" field.
	UpdateDate *time.Time `json:"UpdateDate,omitempty"`
	// OccupationId holds the value of the "OccupationId" field.
	OccupationId int `json:"OccupationId,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Userwallet) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case userwallet.FieldID, userwallet.FieldGroupUser, userwallet.FieldOccupationId:
			values[i] = &sql.NullInt64{}
		case userwallet.FieldWalletid, userwallet.FieldWalletTypeName, userwallet.FieldWalletPhoneno, userwallet.FieldWalletName, userwallet.FieldCitizenId, userwallet.FieldStatus, userwallet.FieldUserAgent, userwallet.FieldATMCard, userwallet.FieldAccountNo, userwallet.FieldAddressDetail, userwallet.FieldStreet, userwallet.FieldDistrict, userwallet.FieldSubDistrict, userwallet.FieldProvince, userwallet.FieldPostalCode, userwallet.FieldIsKYC:
			values[i] = &sql.NullString{}
		case userwallet.FieldRegisterDate, userwallet.FieldKYCDate, userwallet.FieldUpdateDate:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Userwallet", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Userwallet fields.
func (u *Userwallet) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case userwallet.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case userwallet.FieldWalletid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field walletid", values[i])
			} else if value.Valid {
				u.Walletid = value.String
			}
		case userwallet.FieldWalletTypeName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletTypeName", values[i])
			} else if value.Valid {
				u.WalletTypeName = new(string)
				*u.WalletTypeName = value.String
			}
		case userwallet.FieldWalletPhoneno:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletPhoneno", values[i])
			} else if value.Valid {
				u.WalletPhoneno = new(string)
				*u.WalletPhoneno = value.String
			}
		case userwallet.FieldWalletName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletName", values[i])
			} else if value.Valid {
				u.WalletName = new(string)
				*u.WalletName = value.String
			}
		case userwallet.FieldCitizenId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CitizenId", values[i])
			} else if value.Valid {
				u.CitizenId = new(string)
				*u.CitizenId = value.String
			}
		case userwallet.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				u.Status = new(string)
				*u.Status = value.String
			}
		case userwallet.FieldRegisterDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field RegisterDate", values[i])
			} else if value.Valid {
				u.RegisterDate = new(time.Time)
				*u.RegisterDate = value.Time
			}
		case userwallet.FieldGroupUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field GroupUser", values[i])
			} else if value.Valid {
				u.GroupUser = new(int)
				*u.GroupUser = int(value.Int64)
			}
		case userwallet.FieldUserAgent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field UserAgent", values[i])
			} else if value.Valid {
				u.UserAgent = new(string)
				*u.UserAgent = value.String
			}
		case userwallet.FieldKYCDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field KYC_Date", values[i])
			} else if value.Valid {
				u.KYCDate = new(time.Time)
				*u.KYCDate = value.Time
			}
		case userwallet.FieldATMCard:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ATMCard", values[i])
			} else if value.Valid {
				u.ATMCard = new(string)
				*u.ATMCard = value.String
			}
		case userwallet.FieldAccountNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field AccountNo", values[i])
			} else if value.Valid {
				u.AccountNo = new(string)
				*u.AccountNo = value.String
			}
		case userwallet.FieldAddressDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field AddressDetail", values[i])
			} else if value.Valid {
				u.AddressDetail = new(string)
				*u.AddressDetail = value.String
			}
		case userwallet.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Street", values[i])
			} else if value.Valid {
				u.Street = new(string)
				*u.Street = value.String
			}
		case userwallet.FieldDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field District", values[i])
			} else if value.Valid {
				u.District = new(string)
				*u.District = value.String
			}
		case userwallet.FieldSubDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field SubDistrict", values[i])
			} else if value.Valid {
				u.SubDistrict = new(string)
				*u.SubDistrict = value.String
			}
		case userwallet.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Province", values[i])
			} else if value.Valid {
				u.Province = new(string)
				*u.Province = value.String
			}
		case userwallet.FieldPostalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PostalCode", values[i])
			} else if value.Valid {
				u.PostalCode = new(string)
				*u.PostalCode = value.String
			}
		case userwallet.FieldIsKYC:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field isKYC", values[i])
			} else if value.Valid {
				u.IsKYC = new(string)
				*u.IsKYC = value.String
			}
		case userwallet.FieldUpdateDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field UpdateDate", values[i])
			} else if value.Valid {
				u.UpdateDate = new(time.Time)
				*u.UpdateDate = value.Time
			}
		case userwallet.FieldOccupationId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field OccupationId", values[i])
			} else if value.Valid {
				u.OccupationId = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Userwallet.
// Note that you need to call Userwallet.Unwrap() before calling this method if this Userwallet
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Userwallet) Update() *UserwalletUpdateOne {
	return (&UserwalletClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the Userwallet entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Userwallet) Unwrap() *Userwallet {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Userwallet is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Userwallet) String() string {
	var builder strings.Builder
	builder.WriteString("Userwallet(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", walletid=")
	builder.WriteString(u.Walletid)
	if v := u.WalletTypeName; v != nil {
		builder.WriteString(", WalletTypeName=")
		builder.WriteString(*v)
	}
	if v := u.WalletPhoneno; v != nil {
		builder.WriteString(", WalletPhoneno=")
		builder.WriteString(*v)
	}
	if v := u.WalletName; v != nil {
		builder.WriteString(", WalletName=")
		builder.WriteString(*v)
	}
	if v := u.CitizenId; v != nil {
		builder.WriteString(", CitizenId=")
		builder.WriteString(*v)
	}
	if v := u.Status; v != nil {
		builder.WriteString(", Status=")
		builder.WriteString(*v)
	}
	if v := u.RegisterDate; v != nil {
		builder.WriteString(", RegisterDate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.GroupUser; v != nil {
		builder.WriteString(", GroupUser=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := u.UserAgent; v != nil {
		builder.WriteString(", UserAgent=")
		builder.WriteString(*v)
	}
	if v := u.KYCDate; v != nil {
		builder.WriteString(", KYC_Date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := u.ATMCard; v != nil {
		builder.WriteString(", ATMCard=")
		builder.WriteString(*v)
	}
	if v := u.AccountNo; v != nil {
		builder.WriteString(", AccountNo=")
		builder.WriteString(*v)
	}
	if v := u.AddressDetail; v != nil {
		builder.WriteString(", AddressDetail=")
		builder.WriteString(*v)
	}
	if v := u.Street; v != nil {
		builder.WriteString(", Street=")
		builder.WriteString(*v)
	}
	if v := u.District; v != nil {
		builder.WriteString(", District=")
		builder.WriteString(*v)
	}
	if v := u.SubDistrict; v != nil {
		builder.WriteString(", SubDistrict=")
		builder.WriteString(*v)
	}
	if v := u.Province; v != nil {
		builder.WriteString(", Province=")
		builder.WriteString(*v)
	}
	if v := u.PostalCode; v != nil {
		builder.WriteString(", PostalCode=")
		builder.WriteString(*v)
	}
	if v := u.IsKYC; v != nil {
		builder.WriteString(", isKYC=")
		builder.WriteString(*v)
	}
	if v := u.UpdateDate; v != nil {
		builder.WriteString(", UpdateDate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", OccupationId=")
	builder.WriteString(fmt.Sprintf("%v", u.OccupationId))
	builder.WriteByte(')')
	return builder.String()
}

// Userwallets is a parsable slice of Userwallet.
type Userwallets []*Userwallet

func (u Userwallets) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}