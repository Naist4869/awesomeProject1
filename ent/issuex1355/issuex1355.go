// Code generated by entc, DO NOT EDIT.

package issuex1355

const (
	// Label holds the string label denoting the issuex1355 type in the database.
	Label = "issue_x1355"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldInt32 holds the string denoting the int32 field in the database.
	FieldInt32 = "int32"
	// FieldStr holds the string denoting the str field in the database.
	FieldStr = "str"
	// Table holds the table name of the issuex1355 in the database.
	Table = "issue_x1355s"
)

// Columns holds all SQL columns for issuex1355 fields.
var Columns = []string{
	FieldID,
	FieldTime,
	FieldInt32,
	FieldStr,
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
	// DefaultStr holds the default value on creation for the "str" field.
	DefaultStr string
)