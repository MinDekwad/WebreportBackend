// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/occupationhistory"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Occupationhistory is the model entity for the Occupationhistory schema.
type Occupationhistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WalletID holds the value of the "WalletID" field.
	WalletID string `json:"WalletID,omitempty"`
	// OccupationName holds the value of the "OccupationName" field.
	OccupationName string `json:"OccupationName,omitempty"`
	// RankOccupation holds the value of the "RankOccupation" field.
	RankOccupation int `json:"RankOccupation,omitempty"`
	// DateCalRank holds the value of the "DateCalRank" field.
	DateCalRank *time.Time `json:"DateCalRank,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Occupationhistory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case occupationhistory.FieldID, occupationhistory.FieldRankOccupation:
			values[i] = &sql.NullInt64{}
		case occupationhistory.FieldWalletID, occupationhistory.FieldOccupationName:
			values[i] = &sql.NullString{}
		case occupationhistory.FieldDateCalRank:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Occupationhistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Occupationhistory fields.
func (o *Occupationhistory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case occupationhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case occupationhistory.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletID", values[i])
			} else if value.Valid {
				o.WalletID = value.String
			}
		case occupationhistory.FieldOccupationName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field OccupationName", values[i])
			} else if value.Valid {
				o.OccupationName = value.String
			}
		case occupationhistory.FieldRankOccupation:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RankOccupation", values[i])
			} else if value.Valid {
				o.RankOccupation = int(value.Int64)
			}
		case occupationhistory.FieldDateCalRank:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DateCalRank", values[i])
			} else if value.Valid {
				o.DateCalRank = new(time.Time)
				*o.DateCalRank = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Occupationhistory.
// Note that you need to call Occupationhistory.Unwrap() before calling this method if this Occupationhistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Occupationhistory) Update() *OccupationhistoryUpdateOne {
	return (&OccupationhistoryClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Occupationhistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Occupationhistory) Unwrap() *Occupationhistory {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Occupationhistory is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Occupationhistory) String() string {
	var builder strings.Builder
	builder.WriteString("Occupationhistory(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", WalletID=")
	builder.WriteString(o.WalletID)
	builder.WriteString(", OccupationName=")
	builder.WriteString(o.OccupationName)
	builder.WriteString(", RankOccupation=")
	builder.WriteString(fmt.Sprintf("%v", o.RankOccupation))
	if v := o.DateCalRank; v != nil {
		builder.WriteString(", DateCalRank=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Occupationhistories is a parsable slice of Occupationhistory.
type Occupationhistories []*Occupationhistory

func (o Occupationhistories) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}