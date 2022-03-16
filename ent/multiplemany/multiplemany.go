// Code generated by entc, DO NOT EDIT.

package multiplemany

const (
	// Label holds the string label denoting the multiplemany type in the database.
	Label = "multiple_many"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUser1 holds the string denoting the user_1 field in the database.
	FieldUser1 = "user_1"
	// FieldUser2 holds the string denoting the user_2 field in the database.
	FieldUser2 = "user_2"
	// EdgeUser1 holds the string denoting the user1 edge name in mutations.
	EdgeUser1 = "user1"
	// EdgeUser2 holds the string denoting the user2 edge name in mutations.
	EdgeUser2 = "user2"
	// Table holds the table name of the multiplemany in the database.
	Table = "multiple_manies"
	// User1Table is the table that holds the user1 relation/edge.
	User1Table = "multiple_manies"
	// User1InverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	User1InverseTable = "users"
	// User1Column is the table column denoting the user1 relation/edge.
	User1Column = "user_1"
	// User2Table is the table that holds the user2 relation/edge.
	User2Table = "multiple_manies"
	// User2InverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	User2InverseTable = "users"
	// User2Column is the table column denoting the user2 relation/edge.
	User2Column = "user_2"
)

// Columns holds all SQL columns for multiplemany fields.
var Columns = []string{
	FieldID,
	FieldUser1,
	FieldUser2,
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