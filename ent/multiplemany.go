// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/bug/ent/multiplemany"
	"entgo.io/bug/ent/user"
	"entgo.io/ent/dialect/sql"
)

// MultipleMany is the model entity for the MultipleMany schema.
type MultipleMany struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// User1 holds the value of the "user_1" field.
	User1 int `json:"user_1,omitempty"`
	// User2 holds the value of the "user_2" field.
	User2 int `json:"user_2,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MultipleManyQuery when eager-loading is set.
	Edges MultipleManyEdges `json:"edges"`
}

// MultipleManyEdges holds the relations/edges for other nodes in the graph.
type MultipleManyEdges struct {
	// User1 holds the value of the user1 edge.
	User1 *User `json:"user1,omitempty"`
	// User2 holds the value of the user2 edge.
	User2 *User `json:"user2,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// User1OrErr returns the User1 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MultipleManyEdges) User1OrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User1 == nil {
			// The edge user1 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User1, nil
	}
	return nil, &NotLoadedError{edge: "user1"}
}

// User2OrErr returns the User2 value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MultipleManyEdges) User2OrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User2 == nil {
			// The edge user2 was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User2, nil
	}
	return nil, &NotLoadedError{edge: "user2"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MultipleMany) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case multiplemany.FieldID, multiplemany.FieldUser1, multiplemany.FieldUser2:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MultipleMany", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MultipleMany fields.
func (mm *MultipleMany) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case multiplemany.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mm.ID = int(value.Int64)
		case multiplemany.FieldUser1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_1", values[i])
			} else if value.Valid {
				mm.User1 = int(value.Int64)
			}
		case multiplemany.FieldUser2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_2", values[i])
			} else if value.Valid {
				mm.User2 = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser1 queries the "user1" edge of the MultipleMany entity.
func (mm *MultipleMany) QueryUser1() *UserQuery {
	return (&MultipleManyClient{config: mm.config}).QueryUser1(mm)
}

// QueryUser2 queries the "user2" edge of the MultipleMany entity.
func (mm *MultipleMany) QueryUser2() *UserQuery {
	return (&MultipleManyClient{config: mm.config}).QueryUser2(mm)
}

// Update returns a builder for updating this MultipleMany.
// Note that you need to call MultipleMany.Unwrap() before calling this method if this MultipleMany
// was returned from a transaction, and the transaction was committed or rolled back.
func (mm *MultipleMany) Update() *MultipleManyUpdateOne {
	return (&MultipleManyClient{config: mm.config}).UpdateOne(mm)
}

// Unwrap unwraps the MultipleMany entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mm *MultipleMany) Unwrap() *MultipleMany {
	tx, ok := mm.config.driver.(*txDriver)
	if !ok {
		panic("ent: MultipleMany is not a transactional entity")
	}
	mm.config.driver = tx.drv
	return mm
}

// String implements the fmt.Stringer.
func (mm *MultipleMany) String() string {
	var builder strings.Builder
	builder.WriteString("MultipleMany(")
	builder.WriteString(fmt.Sprintf("id=%v", mm.ID))
	builder.WriteString(", user_1=")
	builder.WriteString(fmt.Sprintf("%v", mm.User1))
	builder.WriteString(", user_2=")
	builder.WriteString(fmt.Sprintf("%v", mm.User2))
	builder.WriteByte(')')
	return builder.String()
}

// MultipleManies is a parsable slice of MultipleMany.
type MultipleManies []*MultipleMany

func (mm MultipleManies) config(cfg config) {
	for _i := range mm {
		mm[_i].config = cfg
	}
}
