// Code generated by entc, DO NOT EDIT.

package transactionfactoritem

const (
	// Label holds the string label denoting the transactionfactoritem type in the database.
	Label = "transactionfactoritem"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTransactionFactorID holds the string denoting the transactionfactorid field in the database.
	FieldTransactionFactorID = "TransactionFactorID"
	// FieldMin holds the string denoting the min field in the database.
	FieldMin = "Min"
	// FieldMax holds the string denoting the max field in the database.
	FieldMax = "Max"
	// FieldRank holds the string denoting the rank field in the database.
	FieldRank = "Rank"

	// Table holds the table name of the transactionfactoritem in the database.
	Table = "transaction_factor_item"
)

// Columns holds all SQL columns for transactionfactoritem fields.
var Columns = []string{
	FieldID,
	FieldTransactionFactorID,
	FieldMin,
	FieldMax,
	FieldRank,
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
