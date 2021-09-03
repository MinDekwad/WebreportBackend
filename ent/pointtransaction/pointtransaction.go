// Code generated by entc, DO NOT EDIT.

package pointtransaction

const (
	// Label holds the string label denoting the pointtransaction type in the database.
	Label = "pointtransaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "Date"
	// FieldWalletID holds the string denoting the walletid field in the database.
	FieldWalletID = "WalletID"
	// FieldTransactionName holds the string denoting the transactionname field in the database.
	FieldTransactionName = "TransactionName"
	// FieldPoint holds the string denoting the point field in the database.
	FieldPoint = "Point"
	// FieldReference holds the string denoting the reference field in the database.
	FieldReference = "Reference"

	// Table holds the table name of the pointtransaction in the database.
	Table = "point_transaction"
)

// Columns holds all SQL columns for pointtransaction fields.
var Columns = []string{
	FieldID,
	FieldDate,
	FieldWalletID,
	FieldTransactionName,
	FieldPoint,
	FieldReference,
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
	// TransactionNameValidator is a validator for the "TransactionName" field. It is called by the builders before save.
	TransactionNameValidator func(string) error
)
