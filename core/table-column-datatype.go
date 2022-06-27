package core

import (
	"bytes"
	"fmt"
	"io"
)

type ExtendsTable interface {
	OwnerClass() string
}

type ExtendsColumn interface {
	OwnerClass() string
}

type ExtendsDatatype interface {
	OwnerClass() string
}

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
	// TODO this is spanner specific, need to pass control to T
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "CREATE TABLE %s (\n", t.Name)
	for _, each := range t.Columns {
		each.SQLOn(buf)
	}
	fmt.Fprint(buf, ") PRIMARY KEY (\n", t.Name)
	fmt.Fprintf(buf, ")")
	return buf.String()
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
	Primary    bool        `json:"is_primary"`
	NotNull    bool        `json:"is_not_null"`
	Extensions C           `json:"ext"`
}

func (c *Column[C, D]) SQLOn(buf io.Writer) {
	fmt.Fprintf(buf, "\t%s %s, -- %s\n", c.Name, c.ColumnType.Name, c.Documentation)
}

func (c *Column[C, D]) Doc(d string) *Column[C, D] {
	c.Documentation = d
	return c
}

func (c *Column[C, D]) IsNotNull() *Column[C, D] {
	c.NotNull = true
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
