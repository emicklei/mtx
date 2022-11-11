package db

import (
	"errors"
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

var _ mtx.TypedLabel = new(Column)

// Column is a descriptor of database table schema column.
type Column struct {
	*mtx.Named
	ColumnType mtx.Datatype `json:"type"`
	IsPrimary  bool         `json:"is_primary"`
	// IsNotNull = true means the value is never NULL
	IsNullable bool          `json:"is_nullable"`
	Extensions ExtendsColumn `json:"ext"`
}

func (c *Column) GetDatatype() mtx.Datatype { return c.ColumnType }

func (c *Column) Set(key string, value any) *Column {
	c.Named.Set(key, value)
	return c
}

func (c *Column) Type(d mtx.Datatype) *Column {
	if c.Namespace() != d.Namespace() {
		panic(fmt.Sprintf("cannot set datatype of namespace %s, expected %s", d.Namespace(), c.Namespace()))
	}
	c.ColumnType = d
	return c
}

func (c *Column) Primary() *Column {
	c.IsPrimary = true
	return c
}

// Nullable means the Column value can be NULL.
func (c *Column) Nullable() *Column {
	c.IsNullable = true
	return c
}

func (c *Column) Doc(d string) *Column {
	c.Documentation = d
	return c
}

func (c *Column) SQLOn(buf io.Writer) {
	fmt.Fprintf(buf, "%s %s", c.Name, c.ColumnType.Name)
	if !c.IsNullable {
		fmt.Fprint(buf, " NOT NULL")
	}
	if len(c.Documentation) > 0 {
		fmt.Fprintf(buf, " -- %s\n", c.Documentation)
	}
}

func (c *Column) Validate(e *mtx.ErrorCollector) {
	c.Named.Validate(e)
	if c.ColumnType.Named == nil {
		e.Add(c.Named, errors.New("has no type"))
		return
	}
	if c.ColumnType.Name == mtx.Unknown.Name { // we don't know the class so check against name
		e.Add(c.Named, errors.New("has unknown type"))
		return
	}
	if c.ColumnType.BasicDatatype == nil {
		e.Add(c.Named, errors.New("has unknown attribute type for "+c.ColumnType.String()))
	}
}
