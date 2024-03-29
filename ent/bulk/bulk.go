// Code generated by entc, DO NOT EDIT.

package bulk

const (
	// Label holds the string label denoting the bulk type in the database.
	Label = "bulk"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldBulkCreditSameday holds the string denoting the bulkcreditsameday field in the database.
	FieldBulkCreditSameday = "BulkCreditSameday"
	// FieldBulkCreditSamedayFee holds the string denoting the bulkcreditsamedayfee field in the database.
	FieldBulkCreditSamedayFee = "BulkCreditSamedayFee"
	// FieldTransfertobankaccount holds the string denoting the transfertobankaccount field in the database.
	FieldTransfertobankaccount = "Transfertobankaccount"
	// FieldDateTime holds the string denoting the datetime field in the database.
	FieldDateTime = "DateTime"

	// Table holds the table name of the bulk in the database.
	Table = "bulktransaction"
)

// Columns holds all SQL columns for bulk fields.
var Columns = []string{
	FieldID,
	FieldBulkCreditSameday,
	FieldBulkCreditSamedayFee,
	FieldTransfertobankaccount,
	FieldDateTime,
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
