// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/pendingloanbinding"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Pendingloanbinding is the model entity for the Pendingloanbinding schema.
type Pendingloanbinding struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WalletID holds the value of the "WalletID" field.
	WalletID string `json:"WalletID,omitempty"`
	// NameLB holds the value of the "NameLB" field.
	NameLB string `json:"NameLB,omitempty"`
	// StatusGenLB holds the value of the "StatusGenLB" field.
	StatusGenLB bool `json:"StatusGenLB,omitempty"`
	// PointLB holds the value of the "PointLB" field.
	PointLB int `json:"PointLB,omitempty"`
	// DateTime holds the value of the "DateTime" field.
	DateTime *time.Time `json:"DateTime,omitempty"`
	// DateGenLB holds the value of the "DateGenLB" field.
	DateGenLB *time.Time `json:"DateGenLB,omitempty"`
	// FileimportID holds the value of the "FileimportID" field.
	FileimportID *int `json:"FileimportID,omitempty"`
	// CAWalletID holds the value of the "CAWalletID" field.
	CAWalletID *string `json:"CAWalletID,omitempty"`
	// CAPort holds the value of the "CAPort" field.
	CAPort *string `json:"CAPort,omitempty"`
	// MainBranch holds the value of the "MainBranch" field.
	MainBranch *string `json:"MainBranch,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pendingloanbinding) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pendingloanbinding.FieldStatusGenLB:
			values[i] = &sql.NullBool{}
		case pendingloanbinding.FieldID, pendingloanbinding.FieldPointLB, pendingloanbinding.FieldFileimportID:
			values[i] = &sql.NullInt64{}
		case pendingloanbinding.FieldWalletID, pendingloanbinding.FieldNameLB, pendingloanbinding.FieldCAWalletID, pendingloanbinding.FieldCAPort, pendingloanbinding.FieldMainBranch:
			values[i] = &sql.NullString{}
		case pendingloanbinding.FieldDateTime, pendingloanbinding.FieldDateGenLB:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pendingloanbinding", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pendingloanbinding fields.
func (pe *Pendingloanbinding) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pendingloanbinding.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case pendingloanbinding.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletID", values[i])
			} else if value.Valid {
				pe.WalletID = value.String
			}
		case pendingloanbinding.FieldNameLB:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field NameLB", values[i])
			} else if value.Valid {
				pe.NameLB = value.String
			}
		case pendingloanbinding.FieldStatusGenLB:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field StatusGenLB", values[i])
			} else if value.Valid {
				pe.StatusGenLB = value.Bool
			}
		case pendingloanbinding.FieldPointLB:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field PointLB", values[i])
			} else if value.Valid {
				pe.PointLB = int(value.Int64)
			}
		case pendingloanbinding.FieldDateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DateTime", values[i])
			} else if value.Valid {
				pe.DateTime = new(time.Time)
				*pe.DateTime = value.Time
			}
		case pendingloanbinding.FieldDateGenLB:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DateGenLB", values[i])
			} else if value.Valid {
				pe.DateGenLB = new(time.Time)
				*pe.DateGenLB = value.Time
			}
		case pendingloanbinding.FieldFileimportID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field FileimportID", values[i])
			} else if value.Valid {
				pe.FileimportID = new(int)
				*pe.FileimportID = int(value.Int64)
			}
		case pendingloanbinding.FieldCAWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CAWalletID", values[i])
			} else if value.Valid {
				pe.CAWalletID = new(string)
				*pe.CAWalletID = value.String
			}
		case pendingloanbinding.FieldCAPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field CAPort", values[i])
			} else if value.Valid {
				pe.CAPort = new(string)
				*pe.CAPort = value.String
			}
		case pendingloanbinding.FieldMainBranch:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MainBranch", values[i])
			} else if value.Valid {
				pe.MainBranch = new(string)
				*pe.MainBranch = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Pendingloanbinding.
// Note that you need to call Pendingloanbinding.Unwrap() before calling this method if this Pendingloanbinding
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pendingloanbinding) Update() *PendingloanbindingUpdateOne {
	return (&PendingloanbindingClient{config: pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the Pendingloanbinding entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Pendingloanbinding) Unwrap() *Pendingloanbinding {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pendingloanbinding is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pendingloanbinding) String() string {
	var builder strings.Builder
	builder.WriteString("Pendingloanbinding(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", WalletID=")
	builder.WriteString(pe.WalletID)
	builder.WriteString(", NameLB=")
	builder.WriteString(pe.NameLB)
	builder.WriteString(", StatusGenLB=")
	builder.WriteString(fmt.Sprintf("%v", pe.StatusGenLB))
	builder.WriteString(", PointLB=")
	builder.WriteString(fmt.Sprintf("%v", pe.PointLB))
	if v := pe.DateTime; v != nil {
		builder.WriteString(", DateTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := pe.DateGenLB; v != nil {
		builder.WriteString(", DateGenLB=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := pe.FileimportID; v != nil {
		builder.WriteString(", FileimportID=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := pe.CAWalletID; v != nil {
		builder.WriteString(", CAWalletID=")
		builder.WriteString(*v)
	}
	if v := pe.CAPort; v != nil {
		builder.WriteString(", CAPort=")
		builder.WriteString(*v)
	}
	if v := pe.MainBranch; v != nil {
		builder.WriteString(", MainBranch=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Pendingloanbindings is a parsable slice of Pendingloanbinding.
type Pendingloanbindings []*Pendingloanbinding

func (pe Pendingloanbindings) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
