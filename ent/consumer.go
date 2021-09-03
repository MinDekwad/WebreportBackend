// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"go-api-report2/ent/consumer"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Consumer is the model entity for the Consumer schema.
type Consumer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TransactionID holds the value of the "transaction_id" field.
	TransactionID *string `json:"transaction_id,omitempty"`
	// TransactionStatus holds the value of the "TransactionStatus" field.
	TransactionStatus *string `json:"TransactionStatus,omitempty"`
	// TransactionType holds the value of the "TransactionType" field.
	TransactionType *string `json:"TransactionType,omitempty"`
	// PaymentChannel holds the value of the "PaymentChannel" field.
	PaymentChannel *string `json:"PaymentChannel,omitempty"`
	// PaymentType holds the value of the "PaymentType" field.
	PaymentType *string `json:"PaymentType,omitempty"`
	// TypeCode holds the value of the "TypeCode" field.
	TypeCode *string `json:"TypeCode,omitempty"`
	// ApprovalCode holds the value of the "ApprovalCode" field.
	ApprovalCode *string `json:"ApprovalCode,omitempty"`
	// BillerID holds the value of the "BillerID" field.
	BillerID *string `json:"BillerID,omitempty"`
	// Ref1 holds the value of the "ref1" field.
	Ref1 *string `json:"ref1,omitempty"`
	// Ref2 holds the value of the "ref2" field.
	Ref2 *string `json:"ref2,omitempty"`
	// Ref3 holds the value of the "ref3" field.
	Ref3 *string `json:"ref3,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount *float64 `json:"amount,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee *float64 `json:"fee,omitempty"`
	// Total holds the value of the "total" field.
	Total *float64 `json:"total,omitempty"`
	// FromReference holds the value of the "FromReference" field.
	FromReference *string `json:"FromReference,omitempty"`
	// FromPhoneNo holds the value of the "FromPhoneNo" field.
	FromPhoneNo *string `json:"FromPhoneNo,omitempty"`
	// FromName holds the value of the "FromName" field.
	FromName *string `json:"FromName,omitempty"`
	// ToAccount holds the value of the "ToAccount" field.
	ToAccount *string `json:"ToAccount,omitempty"`
	// ToAccountPhoneNo holds the value of the "ToAccountPhoneNo" field.
	ToAccountPhoneNo *string `json:"ToAccountPhoneNo,omitempty"`
	// ToAccountName holds the value of the "ToAccountName" field.
	ToAccountName *string `json:"ToAccountName,omitempty"`
	// BankCode holds the value of the "BankCode" field.
	BankCode *string `json:"BankCode,omitempty"`
	// TerminalId holds the value of the "TerminalId" field.
	TerminalId *string `json:"TerminalId,omitempty"`
	// TerminalType holds the value of the "TerminalType" field.
	TerminalType *string `json:"TerminalType,omitempty"`
	// ToAccount105 holds the value of the "ToAccount105" field.
	ToAccount105 *string `json:"ToAccount105,omitempty"`
	// FromReference105 holds the value of the "FromReference105" field.
	FromReference105 *string `json:"FromReference105,omitempty"`
	// DateTime holds the value of the "dateTime" field.
	DateTime *time.Time `json:"dateTime,omitempty"`
	// PartnerRef holds the value of the "PartnerRef" field.
	PartnerRef *string `json:"PartnerRef,omitempty"`
	// ResponseCode holds the value of the "ResponseCode" field.
	ResponseCode *string `json:"ResponseCode,omitempty"`
	// ResponseDescription holds the value of the "ResponseDescription" field.
	ResponseDescription *string `json:"ResponseDescription,omitempty"`
	// FileimportID holds the value of the "FileimportID" field.
	FileimportID *int `json:"FileimportID,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Consumer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case consumer.FieldAmount, consumer.FieldFee, consumer.FieldTotal:
			values[i] = &sql.NullFloat64{}
		case consumer.FieldID, consumer.FieldFileimportID:
			values[i] = &sql.NullInt64{}
		case consumer.FieldTransactionID, consumer.FieldTransactionStatus, consumer.FieldTransactionType, consumer.FieldPaymentChannel, consumer.FieldPaymentType, consumer.FieldTypeCode, consumer.FieldApprovalCode, consumer.FieldBillerID, consumer.FieldRef1, consumer.FieldRef2, consumer.FieldRef3, consumer.FieldFromReference, consumer.FieldFromPhoneNo, consumer.FieldFromName, consumer.FieldToAccount, consumer.FieldToAccountPhoneNo, consumer.FieldToAccountName, consumer.FieldBankCode, consumer.FieldTerminalId, consumer.FieldTerminalType, consumer.FieldToAccount105, consumer.FieldFromReference105, consumer.FieldPartnerRef, consumer.FieldResponseCode, consumer.FieldResponseDescription:
			values[i] = &sql.NullString{}
		case consumer.FieldDateTime:
			values[i] = &sql.NullTime{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Consumer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Consumer fields.
func (c *Consumer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case consumer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case consumer.FieldTransactionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field transaction_id", values[i])
			} else if value.Valid {
				c.TransactionID = new(string)
				*c.TransactionID = value.String
			}
		case consumer.FieldTransactionStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionStatus", values[i])
			} else if value.Valid {
				c.TransactionStatus = new(string)
				*c.TransactionStatus = value.String
			}
		case consumer.FieldTransactionType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TransactionType", values[i])
			} else if value.Valid {
				c.TransactionType = new(string)
				*c.TransactionType = value.String
			}
		case consumer.FieldPaymentChannel:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaymentChannel", values[i])
			} else if value.Valid {
				c.PaymentChannel = new(string)
				*c.PaymentChannel = value.String
			}
		case consumer.FieldPaymentType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PaymentType", values[i])
			} else if value.Valid {
				c.PaymentType = new(string)
				*c.PaymentType = value.String
			}
		case consumer.FieldTypeCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TypeCode", values[i])
			} else if value.Valid {
				c.TypeCode = new(string)
				*c.TypeCode = value.String
			}
		case consumer.FieldApprovalCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ApprovalCode", values[i])
			} else if value.Valid {
				c.ApprovalCode = new(string)
				*c.ApprovalCode = value.String
			}
		case consumer.FieldBillerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BillerID", values[i])
			} else if value.Valid {
				c.BillerID = new(string)
				*c.BillerID = value.String
			}
		case consumer.FieldRef1:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref1", values[i])
			} else if value.Valid {
				c.Ref1 = new(string)
				*c.Ref1 = value.String
			}
		case consumer.FieldRef2:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref2", values[i])
			} else if value.Valid {
				c.Ref2 = new(string)
				*c.Ref2 = value.String
			}
		case consumer.FieldRef3:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ref3", values[i])
			} else if value.Valid {
				c.Ref3 = new(string)
				*c.Ref3 = value.String
			}
		case consumer.FieldAmount:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				c.Amount = new(float64)
				*c.Amount = value.Float64
			}
		case consumer.FieldFee:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				c.Fee = new(float64)
				*c.Fee = value.Float64
			}
		case consumer.FieldTotal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value.Valid {
				c.Total = new(float64)
				*c.Total = value.Float64
			}
		case consumer.FieldFromReference:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FromReference", values[i])
			} else if value.Valid {
				c.FromReference = new(string)
				*c.FromReference = value.String
			}
		case consumer.FieldFromPhoneNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FromPhoneNo", values[i])
			} else if value.Valid {
				c.FromPhoneNo = new(string)
				*c.FromPhoneNo = value.String
			}
		case consumer.FieldFromName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FromName", values[i])
			} else if value.Valid {
				c.FromName = new(string)
				*c.FromName = value.String
			}
		case consumer.FieldToAccount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ToAccount", values[i])
			} else if value.Valid {
				c.ToAccount = new(string)
				*c.ToAccount = value.String
			}
		case consumer.FieldToAccountPhoneNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ToAccountPhoneNo", values[i])
			} else if value.Valid {
				c.ToAccountPhoneNo = new(string)
				*c.ToAccountPhoneNo = value.String
			}
		case consumer.FieldToAccountName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ToAccountName", values[i])
			} else if value.Valid {
				c.ToAccountName = new(string)
				*c.ToAccountName = value.String
			}
		case consumer.FieldBankCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field BankCode", values[i])
			} else if value.Valid {
				c.BankCode = new(string)
				*c.BankCode = value.String
			}
		case consumer.FieldTerminalId:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TerminalId", values[i])
			} else if value.Valid {
				c.TerminalId = new(string)
				*c.TerminalId = value.String
			}
		case consumer.FieldTerminalType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field TerminalType", values[i])
			} else if value.Valid {
				c.TerminalType = new(string)
				*c.TerminalType = value.String
			}
		case consumer.FieldToAccount105:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ToAccount105", values[i])
			} else if value.Valid {
				c.ToAccount105 = new(string)
				*c.ToAccount105 = value.String
			}
		case consumer.FieldFromReference105:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FromReference105", values[i])
			} else if value.Valid {
				c.FromReference105 = new(string)
				*c.FromReference105 = value.String
			}
		case consumer.FieldDateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field dateTime", values[i])
			} else if value.Valid {
				c.DateTime = new(time.Time)
				*c.DateTime = value.Time
			}
		case consumer.FieldPartnerRef:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field PartnerRef", values[i])
			} else if value.Valid {
				c.PartnerRef = new(string)
				*c.PartnerRef = value.String
			}
		case consumer.FieldResponseCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ResponseCode", values[i])
			} else if value.Valid {
				c.ResponseCode = new(string)
				*c.ResponseCode = value.String
			}
		case consumer.FieldResponseDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ResponseDescription", values[i])
			} else if value.Valid {
				c.ResponseDescription = new(string)
				*c.ResponseDescription = value.String
			}
		case consumer.FieldFileimportID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field FileimportID", values[i])
			} else if value.Valid {
				c.FileimportID = new(int)
				*c.FileimportID = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Consumer.
// Note that you need to call Consumer.Unwrap() before calling this method if this Consumer
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Consumer) Update() *ConsumerUpdateOne {
	return (&ConsumerClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Consumer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Consumer) Unwrap() *Consumer {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Consumer is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Consumer) String() string {
	var builder strings.Builder
	builder.WriteString("Consumer(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	if v := c.TransactionID; v != nil {
		builder.WriteString(", transaction_id=")
		builder.WriteString(*v)
	}
	if v := c.TransactionStatus; v != nil {
		builder.WriteString(", TransactionStatus=")
		builder.WriteString(*v)
	}
	if v := c.TransactionType; v != nil {
		builder.WriteString(", TransactionType=")
		builder.WriteString(*v)
	}
	if v := c.PaymentChannel; v != nil {
		builder.WriteString(", PaymentChannel=")
		builder.WriteString(*v)
	}
	if v := c.PaymentType; v != nil {
		builder.WriteString(", PaymentType=")
		builder.WriteString(*v)
	}
	if v := c.TypeCode; v != nil {
		builder.WriteString(", TypeCode=")
		builder.WriteString(*v)
	}
	if v := c.ApprovalCode; v != nil {
		builder.WriteString(", ApprovalCode=")
		builder.WriteString(*v)
	}
	if v := c.BillerID; v != nil {
		builder.WriteString(", BillerID=")
		builder.WriteString(*v)
	}
	if v := c.Ref1; v != nil {
		builder.WriteString(", ref1=")
		builder.WriteString(*v)
	}
	if v := c.Ref2; v != nil {
		builder.WriteString(", ref2=")
		builder.WriteString(*v)
	}
	if v := c.Ref3; v != nil {
		builder.WriteString(", ref3=")
		builder.WriteString(*v)
	}
	if v := c.Amount; v != nil {
		builder.WriteString(", amount=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := c.Fee; v != nil {
		builder.WriteString(", fee=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := c.Total; v != nil {
		builder.WriteString(", total=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := c.FromReference; v != nil {
		builder.WriteString(", FromReference=")
		builder.WriteString(*v)
	}
	if v := c.FromPhoneNo; v != nil {
		builder.WriteString(", FromPhoneNo=")
		builder.WriteString(*v)
	}
	if v := c.FromName; v != nil {
		builder.WriteString(", FromName=")
		builder.WriteString(*v)
	}
	if v := c.ToAccount; v != nil {
		builder.WriteString(", ToAccount=")
		builder.WriteString(*v)
	}
	if v := c.ToAccountPhoneNo; v != nil {
		builder.WriteString(", ToAccountPhoneNo=")
		builder.WriteString(*v)
	}
	if v := c.ToAccountName; v != nil {
		builder.WriteString(", ToAccountName=")
		builder.WriteString(*v)
	}
	if v := c.BankCode; v != nil {
		builder.WriteString(", BankCode=")
		builder.WriteString(*v)
	}
	if v := c.TerminalId; v != nil {
		builder.WriteString(", TerminalId=")
		builder.WriteString(*v)
	}
	if v := c.TerminalType; v != nil {
		builder.WriteString(", TerminalType=")
		builder.WriteString(*v)
	}
	if v := c.ToAccount105; v != nil {
		builder.WriteString(", ToAccount105=")
		builder.WriteString(*v)
	}
	if v := c.FromReference105; v != nil {
		builder.WriteString(", FromReference105=")
		builder.WriteString(*v)
	}
	if v := c.DateTime; v != nil {
		builder.WriteString(", dateTime=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	if v := c.PartnerRef; v != nil {
		builder.WriteString(", PartnerRef=")
		builder.WriteString(*v)
	}
	if v := c.ResponseCode; v != nil {
		builder.WriteString(", ResponseCode=")
		builder.WriteString(*v)
	}
	if v := c.ResponseDescription; v != nil {
		builder.WriteString(", ResponseDescription=")
		builder.WriteString(*v)
	}
	if v := c.FileimportID; v != nil {
		builder.WriteString(", FileimportID=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Consumers is a parsable slice of Consumer.
type Consumers []*Consumer

func (c Consumers) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
