// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/transactionfactor"
	"go-api-report2/ent/transactionfactorhistory"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Transactionfactorhistory is the model entity for the Transactionfactorhistory schema.
type Transactionfactorhistory struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// WalletID holds the value of the "WalletID" field.
	WalletID string `json:"WalletID,omitempty"`
	// RankTransactionFactor holds the value of the "RankTransactionFactor" field.
	RankTransactionFactor int `json:"RankTransactionFactor,omitempty"`
	// DateCalRank holds the value of the "DateCalRank" field.
	DateCalRank *time.Time `json:"DateCalRank,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TransactionfactorhistoryQuery when eager-loading is set.
	Edges                TransactionfactorhistoryEdges `json:"edges"`
	_TransactionFactorID *int
}

// TransactionfactorhistoryEdges holds the relations/edges for other nodes in the graph.
type TransactionfactorhistoryEdges struct {
	// Transactionfactor holds the value of the Transactionfactor edge.
	Transactionfactor *Transactionfactor `json:"Transactionfactor,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TransactionfactorOrErr returns the Transactionfactor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TransactionfactorhistoryEdges) TransactionfactorOrErr() (*Transactionfactor, error) {
	if e.loadedTypes[0] {
		if e.Transactionfactor == nil {
			// The edge Transactionfactor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: transactionfactor.Label}
		}
		return e.Transactionfactor, nil
	}
	return nil, &NotLoadedError{edge: "Transactionfactor"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Transactionfactorhistory) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case transactionfactorhistory.FieldID, transactionfactorhistory.FieldRankTransactionFactor:
			values[i] = &sql.NullInt64{}
		case transactionfactorhistory.FieldWalletID:
			values[i] = &sql.NullString{}
		case transactionfactorhistory.FieldDateCalRank:
			values[i] = &sql.NullTime{}
		case transactionfactorhistory.ForeignKeys[0]: // _TransactionFactorID
			values[i] = &sql.NullInt64{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Transactionfactorhistory", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Transactionfactorhistory fields.
func (t *Transactionfactorhistory) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case transactionfactorhistory.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case transactionfactorhistory.FieldWalletID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field WalletID", values[i])
			} else if value.Valid {
				t.WalletID = value.String
			}
		case transactionfactorhistory.FieldRankTransactionFactor:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field RankTransactionFactor", values[i])
			} else if value.Valid {
				t.RankTransactionFactor = int(value.Int64)
			}
		case transactionfactorhistory.FieldDateCalRank:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field DateCalRank", values[i])
			} else if value.Valid {
				t.DateCalRank = new(time.Time)
				*t.DateCalRank = value.Time
			}
		case transactionfactorhistory.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field _TransactionFactorID", value)
			} else if value.Valid {
				t._TransactionFactorID = new(int)
				*t._TransactionFactorID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTransactionfactor queries the "Transactionfactor" edge of the Transactionfactorhistory entity.
func (t *Transactionfactorhistory) QueryTransactionfactor() *TransactionfactorQuery {
	return (&TransactionfactorhistoryClient{config: t.config}).QueryTransactionfactor(t)
}

// Update returns a builder for updating this Transactionfactorhistory.
// Note that you need to call Transactionfactorhistory.Unwrap() before calling this method if this Transactionfactorhistory
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Transactionfactorhistory) Update() *TransactionfactorhistoryUpdateOne {
	return (&TransactionfactorhistoryClient{config: t.config}).UpdateOne(t)
}

// Unwrap unwraps the Transactionfactorhistory entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Transactionfactorhistory) Unwrap() *Transactionfactorhistory {
	tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Transactionfactorhistory is not a transactional entity")
	}
	t.config.driver = tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Transactionfactorhistory) String() string {
	var builder strings.Builder
	builder.WriteString("Transactionfactorhistory(")
	builder.WriteString(fmt.Sprintf("id=%v", t.ID))
	builder.WriteString(", WalletID=")
	builder.WriteString(t.WalletID)
	builder.WriteString(", RankTransactionFactor=")
	builder.WriteString(fmt.Sprintf("%v", t.RankTransactionFactor))
	if v := t.DateCalRank; v != nil {
		builder.WriteString(", DateCalRank=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Transactionfactorhistories is a parsable slice of Transactionfactorhistory.
type Transactionfactorhistories []*Transactionfactorhistory

func (t Transactionfactorhistories) config(cfg config) {
	for _i := range t {
		t[_i].config = cfg
	}
}
