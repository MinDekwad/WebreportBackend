// Code generated by entc, DO NOT EDIT.

package bankdetail

const (
	// Label holds the string label denoting the bankdetail type in the database.
	Label = "bankdetail"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldBankAccountNo holds the string denoting the bank_accountno field in the database.
	FieldBankAccountNo = "Bank_AccountNo"
	// FieldBankName holds the string denoting the bank_name field in the database.
	FieldBankName = "Bank_Name"
	// FieldBankAccountName holds the string denoting the bank_accountname field in the database.
	FieldBankAccountName = "Bank_AccountName"

	// EdgeStatements holds the string denoting the statements edge name in mutations.
	EdgeStatements = "statements"

	// Table holds the table name of the bankdetail in the database.
	Table = "bank_detail"
	// StatementsTable is the table the holds the statements relation/edge.
	StatementsTable = "statement_ending_balance"
	// StatementsInverseTable is the table name for the StatementEndingBalanc entity.
	// It exists in this package in order to avoid circular dependency with the "statementendingbalanc" package.
	StatementsInverseTable = "statement_ending_balance"
	// StatementsColumn is the table column denoting the statements relation/edge.
	StatementsColumn = "Bank_UID"
)

// Columns holds all SQL columns for bankdetail fields.
var Columns = []string{
	FieldID,
	FieldBankAccountNo,
	FieldBankName,
	FieldBankAccountName,
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
	// BankAccountNoValidator is a validator for the "Bank_AccountNo" field. It is called by the builders before save.
	BankAccountNoValidator func(string) error
	// BankNameValidator is a validator for the "Bank_Name" field. It is called by the builders before save.
	BankNameValidator func(string) error
	// BankAccountNameValidator is a validator for the "Bank_AccountName" field. It is called by the builders before save.
	BankAccountNameValidator func(string) error
)