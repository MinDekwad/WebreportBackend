// Code generated by entc, DO NOT EDIT.

package pointpendinglbtransaction

const (
	// Label holds the string label denoting the pointpendinglbtransaction type in the database.
	Label = "pointpendinglbtransaction"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldWalletID holds the string denoting the walletid field in the database.
	FieldWalletID = "WalletID"
	// FieldPointLB holds the string denoting the pointlb field in the database.
	FieldPointLB = "PointLB"
	// FieldDateExportLB holds the string denoting the dateexportlb field in the database.
	FieldDateExportLB = "DateExportLB"
	// FieldDateGenLB holds the string denoting the dategenlb field in the database.
	FieldDateGenLB = "DateGenLB"
	// FieldNoteLB holds the string denoting the notelb field in the database.
	FieldNoteLB = "NoteLB"
	// FieldStatusExportLB holds the string denoting the statusexportlb field in the database.
	FieldStatusExportLB = "StatusExportLB"
	// FieldLBDate holds the string denoting the lbdate field in the database.
	FieldLBDate = "LBDate"

	// Table holds the table name of the pointpendinglbtransaction in the database.
	Table = "point_pending_lb_transaction"
)

// Columns holds all SQL columns for pointpendinglbtransaction fields.
var Columns = []string{
	FieldID,
	FieldWalletID,
	FieldPointLB,
	FieldDateExportLB,
	FieldDateGenLB,
	FieldNoteLB,
	FieldStatusExportLB,
	FieldLBDate,
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
	// DefaultPointLB holds the default value on creation for the "PointLB" field.
	DefaultPointLB int
	// NoteLBValidator is a validator for the "NoteLB" field. It is called by the builders before save.
	NoteLBValidator func(string) error
	// DefaultStatusExportLB holds the default value on creation for the "StatusExportLB" field.
	DefaultStatusExportLB bool
)