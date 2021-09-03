// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/ranking"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Ranking is the model entity for the Ranking schema.
type Ranking struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WalletID holds the value of the "WalletID" field.
	WalletID string `json:"WalletID,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// TaxID holds the value of the "TaxID" field.
	TaxID string `json:"TaxID,omitempty"`
	// ProvinceNameTH holds the value of the "ProvinceNameTH" field.
	ProvinceNameTH string `json:"ProvinceNameTH,omitempty"`
	// DistrictNameTH holds the value of the "DistrictNameTH" field.
	DistrictNameTH string `json:"DistrictNameTH,omitempty"`
	// DistrictNameEN holds the value of the "DistrictNameEN" field.
	DistrictNameEN string `json:"DistrictNameEN,omitempty"`
	// OccupationName holds the value of the "OccupationName" field.
	OccupationName string `json:"OccupationName,omitempty"`
	// LastRank holds the value of the "LastRank" field.
	LastRank int `json:"LastRank,omitempty"`
	// CurrentRank holds the value of the "CurrentRank" field.
	CurrentRank int `json:"CurrentRank,omitempty"`
	// StatusRanking holds the value of the "StatusRanking" field.
	StatusRanking string `json:"StatusRanking,omitempty"`
	// LastDateCalRank holds the value of the "LastDateCalRank" field.
	LastDateCalRank string `json:"LastDateCalRank,omitempty"`
	// NextDateCalRank holds the value of the "NextDateCalRank" field.
	NextDateCalRank string `json:"NextDateCalRank,omitempty"`
	// StateCal holds the value of the "StateCal" field.
	StateCal int `json:"StateCal,omitempty"`
	// ZipCode holds the value of the "ZipCode" field.
	ZipCode *string `json:"ZipCode,omitempty"`
	// TransactionFactorRank holds the value of the "TransactionFactorRank" field.
	TransactionFactorRank int `json:"TransactionFactorRank,omitempty"`
	// RegisDate holds the value of the "RegisDate" field.
	RegisDate *time.Time `json:"RegisDate,omitempty"`
	// SubDistrict holds the value of the "SubDistrict" field.
	SubDistrict string `json:"SubDistrict,omitempty"`
	// Phoneno holds the value of the "Phoneno" field.
	Phoneno *string `json:"Phoneno,omitempty"`
	// AddressDetail holds the value of the "AddressDetail" field.
	AddressDetail *string `json:"AddressDetail,omitempty"`
	// Street holds the value of the "Street" field.
	Street *string `json:"Street,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ranking) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case ranking.FieldID, ranking.FieldLastRank, ranking.FieldCurrentRank, ranking.FieldStateCal, ranking.FieldTransactionFactorRank:
			values[i] = &sql.NullInt64{}
		case ranking.FieldWalletID, ranking.FieldName, ranking.FieldTaxID, ranking.FieldProvinceNameTH, ranking.FieldDistrictNameTH, ranking.FieldDistrictNameEN, ranking.FieldOccupationName, ranking.FieldStatusRanking, ranking.FieldLastDateCalRank, ranking.FieldNextDateCalRank, ranking.FieldZipCode, ranking.FieldSubDistrict, ranking.FieldPhoneno, ranking.FieldAddressDetail, ranking.FieldStreet:
			values[i] = &sql.NullString{}
		case ranking.FieldRegisDate:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Ranking", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ranking fields.
func (r *Ranking) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ranking.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case ranking.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletID", values[i])
			} else if value.Valid {
				r.WalletID = value.String
			}
		case ranking.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				r.Name = value.String
			}
		case ranking.FieldTaxID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TaxID", values[i])
			} else if value.Valid {
				r.TaxID = value.String
			}
		case ranking.FieldProvinceNameTH:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ProvinceNameTH", values[i])
			} else if value.Valid {
				r.ProvinceNameTH = value.String
			}
		case ranking.FieldDistrictNameTH:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DistrictNameTH", values[i])
			} else if value.Valid {
				r.DistrictNameTH = value.String
			}
		case ranking.FieldDistrictNameEN:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field DistrictNameEN", values[i])
			} else if value.Valid {
				r.DistrictNameEN = value.String
			}
		case ranking.FieldOccupationName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OccupationName", values[i])
			} else if value.Valid {
				r.OccupationName = value.String
			}
		case ranking.FieldLastRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field LastRank", values[i])
			} else if value.Valid {
				r.LastRank = int(value.Int64)
			}
		case ranking.FieldCurrentRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field CurrentRank", values[i])
			} else if value.Valid {
				r.CurrentRank = int(value.Int64)
			}
		case ranking.FieldStatusRanking:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field StatusRanking", values[i])
			} else if value.Valid {
				r.StatusRanking = value.String
			}
		case ranking.FieldLastDateCalRank:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LastDateCalRank", values[i])
			} else if value.Valid {
				r.LastDateCalRank = value.String
			}
		case ranking.FieldNextDateCalRank:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field NextDateCalRank", values[i])
			} else if value.Valid {
				r.NextDateCalRank = value.String
			}
		case ranking.FieldStateCal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field StateCal", values[i])
			} else if value.Valid {
				r.StateCal = int(value.Int64)
			}
		case ranking.FieldZipCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ZipCode", values[i])
			} else if value.Valid {
				r.ZipCode = new(string)
				*r.ZipCode = value.String
			}
		case ranking.FieldTransactionFactorRank:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionFactorRank", values[i])
			} else if value.Valid {
				r.TransactionFactorRank = int(value.Int64)
			}
		case ranking.FieldRegisDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field RegisDate", values[i])
			} else if value.Valid {
				r.RegisDate = new(time.Time)
				*r.RegisDate = value.Time
			}
		case ranking.FieldSubDistrict:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field SubDistrict", values[i])
			} else if value.Valid {
				r.SubDistrict = value.String
			}
		case ranking.FieldPhoneno:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Phoneno", values[i])
			} else if value.Valid {
				r.Phoneno = new(string)
				*r.Phoneno = value.String
			}
		case ranking.FieldAddressDetail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field AddressDetail", values[i])
			} else if value.Valid {
				r.AddressDetail = new(string)
				*r.AddressDetail = value.String
			}
		case ranking.FieldStreet:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Street", values[i])
			} else if value.Valid {
				r.Street = new(string)
				*r.Street = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Ranking.
// Note that you need to call Ranking.Unwrap() before calling this method if this Ranking
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Ranking) Update() *RankingUpdateOne {
	return (&RankingClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Ranking entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Ranking) Unwrap() *Ranking {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ranking is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Ranking) String() string {
	var builder strings.Builder
	builder.WriteString("Ranking(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", WalletID=")
	builder.WriteString(r.WalletID)
	builder.WriteString(", Name=")
	builder.WriteString(r.Name)
	builder.WriteString(", TaxID=")
	builder.WriteString(r.TaxID)
	builder.WriteString(", ProvinceNameTH=")
	builder.WriteString(r.ProvinceNameTH)
	builder.WriteString(", DistrictNameTH=")
	builder.WriteString(r.DistrictNameTH)
	builder.WriteString(", DistrictNameEN=")
	builder.WriteString(r.DistrictNameEN)
	builder.WriteString(", OccupationName=")
	builder.WriteString(r.OccupationName)
	builder.WriteString(", LastRank=")
	builder.WriteString(fmt.Sprintf("%v", r.LastRank))
	builder.WriteString(", CurrentRank=")
	builder.WriteString(fmt.Sprintf("%v", r.CurrentRank))
	builder.WriteString(", StatusRanking=")
	builder.WriteString(r.StatusRanking)
	builder.WriteString(", LastDateCalRank=")
	builder.WriteString(r.LastDateCalRank)
	builder.WriteString(", NextDateCalRank=")
	builder.WriteString(r.NextDateCalRank)
	builder.WriteString(", StateCal=")
	builder.WriteString(fmt.Sprintf("%v", r.StateCal))
	if v := r.ZipCode; v != nil {
		builder.WriteString(", ZipCode=")
		builder.WriteString(*v)
	}
	builder.WriteString(", TransactionFactorRank=")
	builder.WriteString(fmt.Sprintf("%v", r.TransactionFactorRank))
	if v := r.RegisDate; v != nil {
		builder.WriteString(", RegisDate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", SubDistrict=")
	builder.WriteString(r.SubDistrict)
	if v := r.Phoneno; v != nil {
		builder.WriteString(", Phoneno=")
		builder.WriteString(*v)
	}
	if v := r.AddressDetail; v != nil {
		builder.WriteString(", AddressDetail=")
		builder.WriteString(*v)
	}
	if v := r.Street; v != nil {
		builder.WriteString(", Street=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Rankings is a parsable slice of Ranking.
type Rankings []*Ranking

func (r Rankings) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
