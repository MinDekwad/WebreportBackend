// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/writelog"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Writelog is the model entity for the Writelog schema.
type Writelog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// AdminID holds the value of the "admin_id" field.
	AdminID int `json:"admin_id,omitempty"`
	// Resource holds the value of the "resource" field.
	Resource *string `json:"resource,omitempty"`
	// Actions holds the value of the "actions" field.
	Actions *string `json:"actions,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Writelog) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case writelog.FieldID, writelog.FieldAdminID:
			values[i] = &sql.NullInt64{}
		case writelog.FieldResource, writelog.FieldActions:
			values[i] = &sql.NullString{}
		case writelog.FieldCreatedAt:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Writelog", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Writelog fields.
func (w *Writelog) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case writelog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case writelog.FieldAdminID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field admin_id", values[i])
			} else if value.Valid {
				w.AdminID = int(value.Int64)
			}
		case writelog.FieldResource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resource", values[i])
			} else if value.Valid {
				w.Resource = new(string)
				*w.Resource = value.String
			}
		case writelog.FieldActions:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field actions", values[i])
			} else if value.Valid {
				w.Actions = new(string)
				*w.Actions = value.String
			}
		case writelog.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = new(time.Time)
				*w.CreatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Writelog.
// Note that you need to call Writelog.Unwrap() before calling this method if this Writelog
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Writelog) Update() *WritelogUpdateOne {
	return (&WritelogClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Writelog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Writelog) Unwrap() *Writelog {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Writelog is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Writelog) String() string {
	var builder strings.Builder
	builder.WriteString("Writelog(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", admin_id=")
	builder.WriteString(fmt.Sprintf("%v", w.AdminID))
	if v := w.Resource; v != nil {
		builder.WriteString(", resource=")
		builder.WriteString(*v)
	}
	if v := w.Actions; v != nil {
		builder.WriteString(", actions=")
		builder.WriteString(*v)
	}
	if v := w.CreatedAt; v != nil {
		builder.WriteString(", created_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Writelogs is a parsable slice of Writelog.
type Writelogs []*Writelog

func (w Writelogs) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
