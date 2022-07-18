package mtx

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type Database struct {
	*Named
	Tables []*Table
	// Views TODO
	Extensions ExtendsDatabase
}

func (d *Database) Doc(doc string) *Database {
	d.Documentation = doc
	return d
}

func (d *Database) Table(name string) *Table {
	ext := d.Extensions.Table()
	return &Table{
		Named:      N(ext.OwnerClass(), name),
		Extensions: ext,
	}
}

type Table struct {
	*Named
	Columns    []*Column    `json:"columns,omitempty"`
	Extensions ExtendsTable `json:"ext"`
}

func (t *Table) Doc(d string) *Table {
	t.Documentation = d
	return t
}

func (t *Table) Set(key string, value any) *Table {
	t.Named.Set(key, value)
	return t
}

func (t *Table) SQL() string {
	buf := new(bytes.Buffer)
	t.Extensions.SQLOn(t, buf)
	return buf.String()
}

func (t *Table) C(name string, d Datatype, doc string) *Column {
	return t.Column(name).Type(d).Doc(doc)
}

// PrimaryKey is short for Column().IsPrimary()
func (t *Table) PrimaryKey(name string) *Column {
	return t.Column(name).Primary()
}

func (t *Table) Column(name string) *Column {
	ext := t.Extensions.Column()
	return &Column{Named: N(ext.OwnerClass(), name)}
}

func (t *Table) PrimaryKeyColumns() (list []*Column) {
	for _, each := range t.Columns {
		if each.IsPrimary {
			list = append(list, each)
		}
	}
	return
}

// ToEntity creates a new Entity that represents a Row in this table data.
// TODO how to handle name mapping?  type mapping?
func (t *Table) ToEntity() *Entity {
	m := NewEntity(t.Name)
	// see if property overrides this
	if n, ok := t.Get(EntityName); ok {
		m.Named.Name = n.(string)
	}
	m.Doc(t.Documentation)
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

type Column struct {
	*Named
	ColumnType Datatype `json:"type"`
	IsPrimary  bool     `json:"is_primary"`
	IsNotNull  bool     `json:"is_not_null"`
	Extensions ExtendsColumn
}

func (c *Column) Set(key string, value any) *Column {
	c.Named.Set(key, value)
	return c
}

func (c *Column) Type(d Datatype) *Column {
	c.ColumnType = d
	return c
}

func (c *Column) Primary() *Column {
	c.IsPrimary = true
	return c
}

func (c *Column) NotNull() *Column {
	c.IsNotNull = true
	return c
}

func (c *Column) Doc(d string) *Column {
	c.Documentation = d
	return c
}

func (c *Column) SQLOn(buf io.Writer) {
	fmt.Fprintf(buf, "%s %s", c.Name, c.ColumnType.Name)
	if c.IsNotNull {
		fmt.Fprint(buf, " NOT NULL")
	}
	fmt.Fprintf(buf, " -- %s\n", c.Documentation)
}

type Datatype struct {
	*Named
	AttributeType AttributeType `json:"-"`
	IsUserDefined bool          `json:"is_user_defined,omitempty"`
	Extensions    ExtendsDatatype
}

func (d Datatype) EncodedFrom(at AttributeType) Datatype {
	d.AttributeType = at
	return d
}

func (d Datatype) String() string {
	return fmt.Sprintf("%s (%s) : %T", d.Name, d.Class, d)
}

func (d Datatype) SourceOn(w io.Writer) {
	pkg := d.Class[0:strings.Index(d.Class, ".")]
	if d.IsUserDefined {
		fmt.Fprintf(w, "%s.Type(\"%s\")", pkg, d.Name)
		return
	}
	fmt.Fprintf(w, "%s.%s", pkg, strings.ToUpper(d.Name))
}

func (d Datatype) AttrType() AttributeType { return d.AttributeType }

func (d Datatype) WithAttributeType(at AttributeType) Datatype {
	d.AttributeType = at
	return d
}

// Set overrides Named.Set to preserve return type
func (d Datatype) Set(key string, value any) Datatype {
	d.Named.Set(key, value)
	return d
}
