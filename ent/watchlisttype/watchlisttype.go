// Code generated by entc, DO NOT EDIT.

package watchlisttype

const (
	// Label holds the string label denoting the watchlisttype type in the database.
	Label = "watchlisttype"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTypeID holds the string denoting the typeid field in the database.
	FieldTypeID = "TypeID"
	// FieldTypeName holds the string denoting the typename field in the database.
	FieldTypeName = "TypeName"
	// FieldTypeDescription holds the string denoting the typedescription field in the database.
	FieldTypeDescription = "TypeDescription"

	// EdgeWatchlist holds the string denoting the watchlist edge name in mutations.
	EdgeWatchlist = "watchlist"

	// Table holds the table name of the watchlisttype in the database.
	Table = "watchlist_type"
	// WatchlistTable is the table the holds the watchlist relation/edge.
	WatchlistTable = "watchlist"
	// WatchlistInverseTable is the table name for the Watchlist entity.
	// It exists in this package in order to avoid circular dependency with the "watchlist" package.
	WatchlistInverseTable = "watchlist"
	// WatchlistColumn is the table column denoting the watchlist relation/edge.
	WatchlistColumn = "TypeID"
)

// Columns holds all SQL columns for watchlisttype fields.
var Columns = []string{
	FieldID,
	FieldTypeID,
	FieldTypeName,
	FieldTypeDescription,
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
	// TypeNameValidator is a validator for the "TypeName" field. It is called by the builders before save.
	TypeNameValidator func(string) error
	// TypeDescriptionValidator is a validator for the "TypeDescription" field. It is called by the builders before save.
	TypeDescriptionValidator func(string) error
)
