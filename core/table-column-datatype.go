package core

import (
	"bytes"
	"fmt"
	"io"
)

type ExtendsTable interface {
	OwnerClass() string
	SQLOn(table any, w io.Writer) // cannot use T,C,D here so we need to cast it back later
}

type ExtendsColumn interface {
	OwnerClass() string
}

type ExtendsDatatype interface {
	OwnerClass() string
}

type SQLWriter interface{ SQLOn(w io.Writer) }

type Table[T ExtendsTable, C ExtendsColumn, D ExtendsDatatype] struct {
	*Named
	Columns    []*Column[C, D]
	Extensions T `json:"ext"`
}

func (t *Table[T, C, D]) Doc(d string) *Table[T, C, D] {
	t.Documentation = d
	return t
}

func (t *Table[T, C, D]) SQL() string {
	buf := new(bytes.Buffer)
	t.Extensions.SQLOn(t, buf)
	return buf.String()
}

func (t *Table[T, C, D]) PrimaryKeyColumns() (list []*Column[C, D]) {
	for _, each := range t.Columns {
		if each.IsPrimary {
			list = append(list, each)
		}
	}
	return
}

// PrimaryKey is short for Column().IsPrimary()
func (t *Table[T, C, D]) PrimaryKey(name string) *Column[C, D] {
	return t.Column(name).Primary()
}

func (t *Table[T, C, D]) Column(name string) *Column[C, D] {
	c, ok := FindByName(t.Columns, name)
	if ok {
		return c
	}
	c = new(Column[C, D])
	c.Named = N(c.Extensions.OwnerClass(), name)
	t.Columns = append(t.Columns, c)
	return c
}

type Column[C ExtendsColumn, D ExtendsDatatype] struct {
	*Named
	ColumnType Datatype[D] `json:"type"`
	IsPrimary  bool        `json:"is_primary"`
	IsNotNull  bool        `json:"is_not_null"`
	Extensions C           `json:"ext"`
}

func (c *Column[C, D]) SQLOn(buf io.Writer) {
	fmt.Fprintf(buf, "\t%s %s", c.Name, c.ColumnType.Name)
	if c.IsNotNull {
		fmt.Fprint(buf, " NOT NULL")
	}
	fmt.Fprintf(buf, ", -- %s\n", c.Documentation)
}

func (c *Column[C, D]) Doc(d string) *Column[C, D] {
	c.Documentation = d
	return c
}

func (c *Column[C, D]) NotNull() *Column[C, D] {
	c.IsNotNull = true
	return c
}

func (c *Column[C, D]) Primary() *Column[C, D] {
	c.IsPrimary = true
	return c
}

func (c *Column[C, D]) Type(dt Datatype[D]) *Column[C, D] {
	c.ColumnType = dt
	return c
}

type Datatype[D ExtendsDatatype] struct {
	*Named
	Extensions D `json:"ext"`
}
