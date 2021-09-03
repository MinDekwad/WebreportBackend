// Code generated by entc, DO NOT EDIT.

package pendingkyc

const (
	// Label holds the string label denoting the pendingkyc type in the database.
	Label = "pendingkyc"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldWalletID holds the string denoting the walletid field in the database.
	FieldWalletID = "WalletID"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "Name"
	// FieldAgentID holds the string denoting the agentid field in the database.
	FieldAgentID = "AgentID"
	// FieldAgentNameLastname holds the string denoting the agentnamelastname field in the database.
	FieldAgentNameLastname = "AgentNameLastname"
	// FieldKYCDate holds the string denoting the kycdate field in the database.
	FieldKYCDate = "KYCDate"
	// FieldDateGen holds the string denoting the dategen field in the database.
	FieldDateGen = "DateGen"
	// FieldStatusGen holds the string denoting the statusgen field in the database.
	FieldStatusGen = "StatusGen"
	// FieldPoint holds the string denoting the point field in the database.
	FieldPoint = "Point"
	// FieldFileimportID holds the string denoting the fileimportid field in the database.
	FieldFileimportID = "FileimportID"

	// Table holds the table name of the pendingkyc in the database.
	Table = "pending_kyc"
)

// Columns holds all SQL columns for pendingkyc fields.
var Columns = []string{
	FieldID,
	FieldWalletID,
	FieldName,
	FieldAgentID,
	FieldAgentNameLastname,
	FieldKYCDate,
	FieldDateGen,
	FieldStatusGen,
	FieldPoint,
	FieldFileimportID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// WalletIDValidator is a validator for the "WalletID" field. It is called by the builders before save.
	WalletIDValidator func(string) error
	// NameValidator is a validator for the "Name" field. It is called by the builders before save.
	NameValidator func(string) error
	// AgentIDValidator is a validator for the "AgentID" field. It is called by the builders before save.
	AgentIDValidator func(string) error
	// AgentNameLastnameValidator is a validator for the "AgentNameLastname" field. It is called by the builders before save.
	AgentNameLastnameValidator func(string) error
	// KYCDateValidator is a validator for the "KYCDate" field. It is called by the builders before save.
	KYCDateValidator func(string) error
	// DefaultStatusGen holds the default value on creation for the "StatusGen" field.
	DefaultStatusGen bool
	// DefaultPoint holds the default value on creation for the "Point" field.
	DefaultPoint int
)