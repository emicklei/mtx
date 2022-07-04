package mtx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type ExtendsTable interface {
	OwnerClass() string
	SQLOn(table any, w io.Writer) // cannot use T,C,D here so we need to type-assert it back later
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
	Columns    []*Column[C, D] `json:"columns"`
	Extensions T               `json:"ext"`
}

// Doc overrides Named.Doc to preserve return type
func (t *Table[T, C, D]) Doc(d string) *Table[T, C, D] {
	t.Documentation = d
	return t
}

// Set overrides Named.Set to preserve return type
func (t *Table[T, C, D]) Set(key string, value any) *Table[T, C, D] {
	t.Named.Set(key, value)
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

// C is a shortcut for Column.Type.Doc
func (t *Table[T, C, D]) C(name string, dt Datatype[D], doc string) *Column[C, D] {
	return t.Column(name).Type(dt).Doc(doc)
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

// ToEntity creates a new Entity that represents a Row in this table data.
// TODO how to handle name mapping?  type mapping?
func (t *Table[T, C, D]) ToEntity() *Entity {
	m := NewEntity(t.Name)
	// see if property overrides this
	if n, ok := t.Get(EntityName); ok {
		m.Named.Name = n.(string)
	}
	for _, each := range t.Columns {
		attr := m.Attribute(each.Name)
		// see if property overrides this
		if n, ok := each.Get(EntityName); ok {
			attr.Named.Name = n.(string)
		}
		attr.Doc(each.Documentation)
		attr.AttributeType = each.ColumnType.AttributeType
	}
	return m
}

type Column[C ExtendsColumn, D ExtendsDatatype] struct {
	*Named
	ColumnType Datatype[D] `json:"type"`
	IsPrimary  bool        `json:"is_primary"`
	IsNotNull  bool        `json:"is_not_null"`
	Extensions C           `json:"ext"`
}

func (c *Column[C, D]) SQLOn(buf io.Writer) {
	fmt.Fprintf(buf, "%s %s", c.Name, c.ColumnType.Name)
	if c.IsNotNull {
		fmt.Fprint(buf, " NOT NULL")
	}
	fmt.Fprintf(buf, " -- %s\n", c.Documentation)
}

// Doc overrides Named.Doc to preserve return type
func (c *Column[C, D]) Doc(d string) *Column[C, D] {
	c.Documentation = d
	return c
}

// Set overrides Named.Set to preserve return type
func (c *Column[C, D]) Set(key string, value any) *Column[C, D] {
	c.Named.Set(key, value)
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
	AttributeType AttributeType `json:"-"`
	IsUserDefined bool          `json:"is_user_defined,omitempty"`
	Extensions    D             `json:"ext,omitempty"`
}

// Set overrides Named.Set to preserve return type
func (d Datatype[D]) Set(key string, value any) Datatype[D] {
	d.Named.Set(key, value)
	return d
}

func (d Datatype[D]) WithAttributeType(at AttributeType) Datatype[D] {
	d.AttributeType = at
	return d
}

func (d Datatype[D]) AttrType() AttributeType { return d.AttributeType }

func (d Datatype[D]) SourceOn(w io.Writer) {
	pkg := d.Class[0:strings.Index(d.Class, ".")]
	if d.IsUserDefined {
		fmt.Fprintf(w, "%s.Type(\"%s\")", pkg, d.Name)
		return
	}
	fmt.Fprintf(w, "%s.%s", pkg, strings.ToUpper(d.Name))
}
