// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/pointkycrv"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Pointkycrv is the model entity for the Pointkycrv schema.
type Pointkycrv struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateTimeGen holds the value of the "DateTimeGen" field.
	DateTimeGen time.Time `json:"DateTimeGen,omitempty"`
	// WalletID holds the value of the "WalletID" field.
	WalletID string `json:"WalletID,omitempty"`
	// StatusGen holds the value of the "StatusGen" field.
	StatusGen string `json:"StatusGen,omitempty"`
	// KYCDate holds the value of the "KYCDate" field.
	KYCDate time.Time `json:"KYCDate,omitempty"`
	// RVDate holds the value of the "RVDate" field.
	RVDate time.Time `json:"RVDate,omitempty"`
	// Type holds the value of the "Type" field.
	Type string `json:"Type,omitempty"`
	// Point holds the value of the "Point" field.
	Point int `json:"Point,omitempty"`
	// FileimportID holds the value of the "FileimportID" field.
	FileimportID *int `json:"FileimportID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pointkycrv) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pointkycrv.FieldID, pointkycrv.FieldPoint, pointkycrv.FieldFileimportID:
			values[i] = &sql.NullInt64{}
		case pointkycrv.FieldWalletID, pointkycrv.FieldStatusGen, pointkycrv.FieldType:
			values[i] = &sql.NullString{}
		case pointkycrv.FieldDateTimeGen, pointkycrv.FieldKYCDate, pointkycrv.FieldRVDate:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pointkycrv", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pointkycrv fields.
func (po *Pointkycrv) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pointkycrv.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case pointkycrv.FieldDateTimeGen:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DateTimeGen", values[i])
			} else if value.Valid {
				po.DateTimeGen = value.Time
			}
		case pointkycrv.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletID", values[i])
			} else if value.Valid {
				po.WalletID = value.String
			}
		case pointkycrv.FieldStatusGen:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field StatusGen", values[i])
			} else if value.Valid {
				po.StatusGen = value.String
			}
		case pointkycrv.FieldKYCDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field KYCDate", values[i])
			} else if value.Valid {
				po.KYCDate = value.Time
			}
		case pointkycrv.FieldRVDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field RVDate", values[i])
			} else if value.Valid {
				po.RVDate = value.Time
			}
		case pointkycrv.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Type", values[i])
			} else if value.Valid {
				po.Type = value.String
			}
		case pointkycrv.FieldPoint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Point", values[i])
			} else if value.Valid {
				po.Point = int(value.Int64)
			}
		case pointkycrv.FieldFileimportID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field FileimportID", values[i])
			} else if value.Valid {
				po.FileimportID = new(int)
				*po.FileimportID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Pointkycrv.
// Note that you need to call Pointkycrv.Unwrap() before calling this method if this Pointkycrv
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pointkycrv) Update() *PointkycrvUpdateOne {
	return (&PointkycrvClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Pointkycrv entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pointkycrv) Unwrap() *Pointkycrv {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pointkycrv is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pointkycrv) String() string {
	var builder strings.Builder
	builder.WriteString("Pointkycrv(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", DateTimeGen=")
	builder.WriteString(po.DateTimeGen.Format(time.ANSIC))
	builder.WriteString(", WalletID=")
	builder.WriteString(po.WalletID)
	builder.WriteString(", StatusGen=")
	builder.WriteString(po.StatusGen)
	builder.WriteString(", KYCDate=")
	builder.WriteString(po.KYCDate.Format(time.ANSIC))
	builder.WriteString(", RVDate=")
	builder.WriteString(po.RVDate.Format(time.ANSIC))
	builder.WriteString(", Type=")
	builder.WriteString(po.Type)
	builder.WriteString(", Point=")
	builder.WriteString(fmt.Sprintf("%v", po.Point))
	if v := po.FileimportID; v != nil {
		builder.WriteString(", FileimportID=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Pointkycrvs is a parsable slice of Pointkycrv.
type Pointkycrvs []*Pointkycrv

func (po Pointkycrvs) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
