// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/pointcsv"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Pointcsv is the model entity for the Pointcsv schema.
type Pointcsv struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WalletID holds the value of the "WalletID" field.
	WalletID string `json:"WalletID,omitempty"`
	// CreateDate holds the value of the "CreateDate" field.
	CreateDate time.Time `json:"CreateDate,omitempty"`
	// Adjustamount holds the value of the "Adjustamount" field.
	Adjustamount int `json:"Adjustamount,omitempty"`
	// Note holds the value of the "Note" field.
	Note *string `json:"Note,omitempty"`
	// PointTranDate holds the value of the "PointTranDate" field.
	PointTranDate time.Time `json:"PointTranDate,omitempty"`
	// ActionExport holds the value of the "ActionExport" field.
	ActionExport int `json:"ActionExport,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pointcsv) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pointcsv.FieldID, pointcsv.FieldAdjustamount, pointcsv.FieldActionExport:
			values[i] = &sql.NullInt64{}
		case pointcsv.FieldWalletID, pointcsv.FieldNote:
			values[i] = &sql.NullString{}
		case pointcsv.FieldCreateDate, pointcsv.FieldPointTranDate:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pointcsv", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pointcsv fields.
func (po *Pointcsv) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pointcsv.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case pointcsv.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletID", values[i])
			} else if value.Valid {
				po.WalletID = value.String
			}
		case pointcsv.FieldCreateDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field CreateDate", values[i])
			} else if value.Valid {
				po.CreateDate = value.Time
			}
		case pointcsv.FieldAdjustamount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Adjustamount", values[i])
			} else if value.Valid {
				po.Adjustamount = int(value.Int64)
			}
		case pointcsv.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Note", values[i])
			} else if value.Valid {
				po.Note = new(string)
				*po.Note = value.String
			}
		case pointcsv.FieldPointTranDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field PointTranDate", values[i])
			} else if value.Valid {
				po.PointTranDate = value.Time
			}
		case pointcsv.FieldActionExport:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ActionExport", values[i])
			} else if value.Valid {
				po.ActionExport = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Pointcsv.
// Note that you need to call Pointcsv.Unwrap() before calling this method if this Pointcsv
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Pointcsv) Update() *PointcsvUpdateOne {
	return (&PointcsvClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Pointcsv entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Pointcsv) Unwrap() *Pointcsv {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pointcsv is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Pointcsv) String() string {
	var builder strings.Builder
	builder.WriteString("Pointcsv(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", WalletID=")
	builder.WriteString(po.WalletID)
	builder.WriteString(", CreateDate=")
	builder.WriteString(po.CreateDate.Format(time.ANSIC))
	builder.WriteString(", Adjustamount=")
	builder.WriteString(fmt.Sprintf("%v", po.Adjustamount))
	if v := po.Note; v != nil {
		builder.WriteString(", Note=")
		builder.WriteString(*v)
	}
	builder.WriteString(", PointTranDate=")
	builder.WriteString(po.PointTranDate.Format(time.ANSIC))
	builder.WriteString(", ActionExport=")
	builder.WriteString(fmt.Sprintf("%v", po.ActionExport))
	builder.WriteByte(')')
	return builder.String()
}

// Pointcsvs is a parsable slice of Pointcsv.
type Pointcsvs []*Pointcsv

func (po Pointcsvs) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
