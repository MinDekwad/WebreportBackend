// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/merchanttransaction"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// MerchantTransaction is the model entity for the MerchantTransaction schema.
type MerchantTransaction struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID *string `json:"transaction_id,omitempty"`
	// DateTime holds the value of the "dateTime" field.
	DateTime *time.Time `json:"dateTime,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount *float64 `json:"amount,omitempty"`
	// PaymentType holds the value of the "PaymentType" field.
	PaymentType *string `json:"PaymentType,omitempty"`
	// PaymentChannel holds the value of the "PaymentChannel" field.
	PaymentChannel *string `json:"PaymentChannel,omitempty"`
	// Status holds the value of the "Status" field.
	Status *string `json:"Status,omitempty"`
	// MerchantID holds the value of the "MerchantID" field.
	MerchantID *string `json:"MerchantID,omitempty"`
	// TerminalID holds the value of the "TerminalID" field.
	TerminalID *string `json:"TerminalID,omitempty"`
	// MerchantFullName holds the value of the "MerchantFullName" field.
	MerchantFullName *string `json:"MerchantFullName,omitempty"`
	// FromAccount holds the value of the "FromAccount" field.
	FromAccount *string `json:"FromAccount,omitempty"`
	// SettlementAccount holds the value of the "SettlementAccount" field.
	SettlementAccount *string `json:"SettlementAccount,omitempty"`
	// MDRFEETHB holds the value of the "MDR_FEETHB" field.
	MDRFEETHB *float64 `json:"MDR_FEETHB,omitempty"`
	// TransactionFEETHB holds the value of the "TransactionFEETHB" field.
	TransactionFEETHB *float64 `json:"TransactionFEETHB,omitempty"`
	// TotalFeeincVAT holds the value of the "TotalFeeincVAT" field.
	TotalFeeincVAT *float64 `json:"TotalFeeincVAT,omitempty"`
	// VATTHB holds the value of the "VATTHB" field.
	VATTHB *float64 `json:"VATTHB,omitempty"`
	// TotalFeeExcVAT holds the value of the "TotalFeeExcVAT" field.
	TotalFeeExcVAT *float64 `json:"TotalFeeExcVAT,omitempty"`
	// WHTTHB holds the value of the "WHTTHB" field.
	WHTTHB *float64 `json:"WHTTHB,omitempty"`
	// NetPayTHB holds the value of the "NetPayTHB" field.
	NetPayTHB *float64 `json:"NetPayTHB,omitempty"`
	// TransactionType holds the value of the "TransactionType" field.
	TransactionType *string `json:"TransactionType,omitempty"`
	// BankCode holds the value of the "BankCode" field.
	BankCode *string `json:"BankCode,omitempty"`
	// FileimportID holds the value of the "FileimportID" field.
	FileimportID *int `json:"FileimportID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MerchantTransaction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case merchanttransaction.FieldAmount, merchanttransaction.FieldMDRFEETHB, merchanttransaction.FieldTransactionFEETHB, merchanttransaction.FieldTotalFeeincVAT, merchanttransaction.FieldVATTHB, merchanttransaction.FieldTotalFeeExcVAT, merchanttransaction.FieldWHTTHB, merchanttransaction.FieldNetPayTHB:
			values[i] = &sql.NullFloat64{}
		case merchanttransaction.FieldID, merchanttransaction.FieldFileimportID:
			values[i] = &sql.NullInt64{}
		case merchanttransaction.FieldTransactionID, merchanttransaction.FieldPaymentType, merchanttransaction.FieldPaymentChannel, merchanttransaction.FieldStatus, merchanttransaction.FieldMerchantID, merchanttransaction.FieldTerminalID, merchanttransaction.FieldMerchantFullName, merchanttransaction.FieldFromAccount, merchanttransaction.FieldSettlementAccount, merchanttransaction.FieldTransactionType, merchanttransaction.FieldBankCode:
			values[i] = &sql.NullString{}
		case merchanttransaction.FieldDateTime:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type MerchantTransaction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MerchantTransaction fields.
func (mt *MerchantTransaction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case merchanttransaction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mt.ID = int(value.Int64)
		case merchanttransaction.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				mt.TransactionID = new(string)
				*mt.TransactionID = value.String
			}
		case merchanttransaction.FieldDateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateTime", values[i])
			} else if value.Valid {
				mt.DateTime = new(time.Time)
				*mt.DateTime = value.Time
			}
		case merchanttransaction.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				mt.Amount = new(float64)
				*mt.Amount = value.Float64
			}
		case merchanttransaction.FieldPaymentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaymentType", values[i])
			} else if value.Valid {
				mt.PaymentType = new(string)
				*mt.PaymentType = value.String
			}
		case merchanttransaction.FieldPaymentChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaymentChannel", values[i])
			} else if value.Valid {
				mt.PaymentChannel = new(string)
				*mt.PaymentChannel = value.String
			}
		case merchanttransaction.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Status", values[i])
			} else if value.Valid {
				mt.Status = new(string)
				*mt.Status = value.String
			}
		case merchanttransaction.FieldMerchantID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MerchantID", values[i])
			} else if value.Valid {
				mt.MerchantID = new(string)
				*mt.MerchantID = value.String
			}
		case merchanttransaction.FieldTerminalID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TerminalID", values[i])
			} else if value.Valid {
				mt.TerminalID = new(string)
				*mt.TerminalID = value.String
			}
		case merchanttransaction.FieldMerchantFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field MerchantFullName", values[i])
			} else if value.Valid {
				mt.MerchantFullName = new(string)
				*mt.MerchantFullName = value.String
			}
		case merchanttransaction.FieldFromAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FromAccount", values[i])
			} else if value.Valid {
				mt.FromAccount = new(string)
				*mt.FromAccount = value.String
			}
		case merchanttransaction.FieldSettlementAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field SettlementAccount", values[i])
			} else if value.Valid {
				mt.SettlementAccount = new(string)
				*mt.SettlementAccount = value.String
			}
		case merchanttransaction.FieldMDRFEETHB:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field MDR_FEETHB", values[i])
			} else if value.Valid {
				mt.MDRFEETHB = new(float64)
				*mt.MDRFEETHB = value.Float64
			}
		case merchanttransaction.FieldTransactionFEETHB:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionFEETHB", values[i])
			} else if value.Valid {
				mt.TransactionFEETHB = new(float64)
				*mt.TransactionFEETHB = value.Float64
			}
		case merchanttransaction.FieldTotalFeeincVAT:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field TotalFeeincVAT", values[i])
			} else if value.Valid {
				mt.TotalFeeincVAT = new(float64)
				*mt.TotalFeeincVAT = value.Float64
			}
		case merchanttransaction.FieldVATTHB:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field VATTHB", values[i])
			} else if value.Valid {
				mt.VATTHB = new(float64)
				*mt.VATTHB = value.Float64
			}
		case merchanttransaction.FieldTotalFeeExcVAT:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field TotalFeeExcVAT", values[i])
			} else if value.Valid {
				mt.TotalFeeExcVAT = new(float64)
				*mt.TotalFeeExcVAT = value.Float64
			}
		case merchanttransaction.FieldWHTTHB:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field WHTTHB", values[i])
			} else if value.Valid {
				mt.WHTTHB = new(float64)
				*mt.WHTTHB = value.Float64
			}
		case merchanttransaction.FieldNetPayTHB:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field NetPayTHB", values[i])
			} else if value.Valid {
				mt.NetPayTHB = new(float64)
				*mt.NetPayTHB = value.Float64
			}
		case merchanttransaction.FieldTransactionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionType", values[i])
			} else if value.Valid {
				mt.TransactionType = new(string)
				*mt.TransactionType = value.String
			}
		case merchanttransaction.FieldBankCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BankCode", values[i])
			} else if value.Valid {
				mt.BankCode = new(string)
				*mt.BankCode = value.String
			}
		case merchanttransaction.FieldFileimportID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field FileimportID", values[i])
			} else if value.Valid {
				mt.FileimportID = new(int)
				*mt.FileimportID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MerchantTransaction.
// Note that you need to call MerchantTransaction.Unwrap() before calling this method if this MerchantTransaction
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MerchantTransaction) Update() *MerchantTransactionUpdateOne {
	return (&MerchantTransactionClient{config: mt.config}).UpdateOne(mt)
}

// Unwrap unwraps the MerchantTransaction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mt *MerchantTransaction) Unwrap() *MerchantTransaction {
	tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MerchantTransaction is not a transactional entity")
	}
	mt.config.driver = tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MerchantTransaction) String() string {
	var builder strings.Builder
	builder.WriteString("MerchantTransaction(")
	builder.WriteString(fmt.Sprintf("id=%v", mt.ID))
	if v := mt.TransactionID; v != nil {
		builder.WriteString(", transaction_id=")
		builder.WriteString(*v)
	}
	if v := mt.DateTime; v != nil {
		builder.WriteString(", dateTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := mt.Amount; v != nil {
		builder.WriteString(", amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.PaymentType; v != nil {
		builder.WriteString(", PaymentType=")
		builder.WriteString(*v)
	}
	if v := mt.PaymentChannel; v != nil {
		builder.WriteString(", PaymentChannel=")
		builder.WriteString(*v)
	}
	if v := mt.Status; v != nil {
		builder.WriteString(", Status=")
		builder.WriteString(*v)
	}
	if v := mt.MerchantID; v != nil {
		builder.WriteString(", MerchantID=")
		builder.WriteString(*v)
	}
	if v := mt.TerminalID; v != nil {
		builder.WriteString(", TerminalID=")
		builder.WriteString(*v)
	}
	if v := mt.MerchantFullName; v != nil {
		builder.WriteString(", MerchantFullName=")
		builder.WriteString(*v)
	}
	if v := mt.FromAccount; v != nil {
		builder.WriteString(", FromAccount=")
		builder.WriteString(*v)
	}
	if v := mt.SettlementAccount; v != nil {
		builder.WriteString(", SettlementAccount=")
		builder.WriteString(*v)
	}
	if v := mt.MDRFEETHB; v != nil {
		builder.WriteString(", MDR_FEETHB=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.TransactionFEETHB; v != nil {
		builder.WriteString(", TransactionFEETHB=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.TotalFeeincVAT; v != nil {
		builder.WriteString(", TotalFeeincVAT=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.VATTHB; v != nil {
		builder.WriteString(", VATTHB=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.TotalFeeExcVAT; v != nil {
		builder.WriteString(", TotalFeeExcVAT=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.WHTTHB; v != nil {
		builder.WriteString(", WHTTHB=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.NetPayTHB; v != nil {
		builder.WriteString(", NetPayTHB=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := mt.TransactionType; v != nil {
		builder.WriteString(", TransactionType=")
		builder.WriteString(*v)
	}
	if v := mt.BankCode; v != nil {
		builder.WriteString(", BankCode=")
		builder.WriteString(*v)
	}
	if v := mt.FileimportID; v != nil {
		builder.WriteString(", FileimportID=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// MerchantTransactions is a parsable slice of MerchantTransaction.
type MerchantTransactions []*MerchantTransaction

func (mt MerchantTransactions) config(cfg config) {
	for _i := range mt {
		mt[_i].config = cfg
	}
}
