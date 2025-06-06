// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv2/blog"
)

// Blog is the model entity for the Blog schema.
type Blog struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Oid holds the value of the "oid" field.
	Oid int `json:"oid,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlogQuery when eager-loading is set.
	Edges        BlogEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BlogEdges holds the relations/edges for other nodes in the graph.
type BlogEdges struct {
	// Admins holds the value of the admins edge.
	Admins []*User `json:"admins,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// AdminsOrErr returns the Admins value or an error if the edge
// was not loaded in eager-loading.
func (e BlogEdges) AdminsOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.Admins, nil
	}
	return nil, &NotLoadedError{edge: "admins"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Blog) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blog.FieldID, blog.FieldOid:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Blog fields.
func (_m *Blog) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blog.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case blog.FieldOid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				_m.Oid = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Blog.
// This includes values selected through modifiers, order, etc.
func (_m *Blog) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryAdmins queries the "admins" edge of the Blog entity.
func (_m *Blog) QueryAdmins() *UserQuery {
	return NewBlogClient(_m.config).QueryAdmins(_m)
}

// Update returns a builder for updating this Blog.
// Note that you need to call Blog.Unwrap() before calling this method if this Blog
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Blog) Update() *BlogUpdateOne {
	return NewBlogClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Blog entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Blog) Unwrap() *Blog {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Blog is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Blog) String() string {
	var builder strings.Builder
	builder.WriteString("Blog(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("oid=")
	builder.WriteString(fmt.Sprintf("%v", _m.Oid))
	builder.WriteByte(')')
	return builder.String()
}

// Blogs is a parsable slice of Blog.
type Blogs []*Blog
