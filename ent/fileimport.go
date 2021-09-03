// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/fileimport"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Fileimport is the model entity for the Fileimport schema.
type Fileimport struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename *string `json:"filename,omitempty"`
	// Filetype holds the value of the "filetype" field.
	Filetype *string `json:"filetype,omitempty"`
	// Importdate holds the value of the "importdate" field.
	Importdate *time.Time `json:"importdate,omitempty"`
	// Status holds the value of the "status" field.
	Status *string `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Fileimport) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fileimport.FieldID:
			values[i] = &sql.NullInt64{}
		case fileimport.FieldFilename, fileimport.FieldFiletype, fileimport.FieldStatus:
			values[i] = &sql.NullString{}
		case fileimport.FieldImportdate:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Fileimport", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Fileimport fields.
func (f *Fileimport) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fileimport.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case fileimport.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				f.Filename = new(string)
				*f.Filename = value.String
			}
		case fileimport.FieldFiletype:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filetype", values[i])
			} else if value.Valid {
				f.Filetype = new(string)
				*f.Filetype = value.String
			}
		case fileimport.FieldImportdate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field importdate", values[i])
			} else if value.Valid {
				f.Importdate = new(time.Time)
				*f.Importdate = value.Time
			}
		case fileimport.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = new(string)
				*f.Status = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Fileimport.
// Note that you need to call Fileimport.Unwrap() before calling this method if this Fileimport
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Fileimport) Update() *FileimportUpdateOne {
	return (&FileimportClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Fileimport entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Fileimport) Unwrap() *Fileimport {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Fileimport is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Fileimport) String() string {
	var builder strings.Builder
	builder.WriteString("Fileimport(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	if v := f.Filename; v != nil {
		builder.WriteString(", filename=")
		builder.WriteString(*v)
	}
	if v := f.Filetype; v != nil {
		builder.WriteString(", filetype=")
		builder.WriteString(*v)
	}
	if v := f.Importdate; v != nil {
		builder.WriteString(", importdate=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := f.Status; v != nil {
		builder.WriteString(", status=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Fileimports is a parsable slice of Fileimport.
type Fileimports []*Fileimport

func (f Fileimports) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
