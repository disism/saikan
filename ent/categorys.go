// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/disism/saikan/ent/categorys"
)

// Categorys is the model entity for the Categorys schema.
type Categorys struct {
	config
	// ID of the ent.
	ID           int `json:"id,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Categorys) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case categorys.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Categorys fields.
func (c *Categorys) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case categorys.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Categorys.
// This includes values selected through modifiers, order, etc.
func (c *Categorys) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Categorys.
// Note that you need to call Categorys.Unwrap() before calling this method if this Categorys
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Categorys) Update() *CategorysUpdateOne {
	return NewCategorysClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Categorys entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Categorys) Unwrap() *Categorys {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Categorys is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Categorys) String() string {
	var builder strings.Builder
	builder.WriteString("Categorys(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteByte(')')
	return builder.String()
}

// CategorysSlice is a parsable slice of Categorys.
type CategorysSlice []*Categorys
