package db

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

type Database struct {
	*mtx.Named
	Tables []*Table `json:"tables"`
	// Views TODO
	Extensions ExtendsDatabase `json:"ext"`
}

func (d *Database) Doc(doc string) *Database {
	d.Documentation = doc
	return d
}

// Table is a descriptor or schema of a database table.
// If not exits then create an empty.
func (d *Database) Table(name string) *Table {
	if t, ok := mtx.FindByName(d.Tables, name); ok {
		return t
	}
	ext := d.Extensions.TableExtensions()
	t := &Table{
		Named:      mtx.N(ext.OwnerClass(), name),
		Extensions: ext,
	}
	d.Tables = append(d.Tables, t)
	return t
}

type Table struct {
	*mtx.Named
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

func (t *Table) ToSQL() string {
	buf := new(bytes.Buffer)
	t.Extensions.SQLOn(t, buf)
	return buf.String()
}

func (t *Table) C(name string, d mtx.Datatype, doc string) *Column {
	return t.Column(name).Type(d).Doc(doc)
}

// PrimaryKey is short for Column().IsPrimary()
func (t *Table) PrimaryKey(name string) *Column {
	return t.Column(name).Primary()
}

// Column is a descriptor of database table schema column.
// If not exits then create an empty.
func (t *Table) Column(name string) *Column {
	if c, ok := mtx.FindByName(t.Columns, name); ok {
		return c
	}
	ext := t.Extensions.ColumnExtensions()
	c := &Column{Named: mtx.N(ext.OwnerClass(), name), Extensions: ext}
	t.Columns = append(t.Columns, c)
	return c
}

func (t *Table) PrimaryKeyColumns() (list []*Column) {
	for _, each := range t.Columns {
		if each.IsPrimary {
			list = append(list, each)
		}
	}
	return
}

func (t *Table) Validate(c *mtx.ErrorCollector) {
	t.Named.Validate(c)
	for _, each := range t.Columns {
		each.Validate(c)
	}
}

// ToEntity creates a new Entity that represents a Row in this table data.
func (t *Table) ToEntity() *mtx.Entity {
	m := mtx.NewEntity(t.Name)
	// see if property overrides this
	if n, ok := t.Get(mtx.EntityName); ok {
		m.Named.Name = n.(string)
	}
	m.Doc(t.Documentation)
	for _, each := range t.Columns {
		attr := m.Attribute(each.Name)
		// see if property overrides this
		if n, ok := each.Get(mtx.AttributeName); ok {
			attr.Named.Name = n.(string)
		}
		if at := each.GetDatatype().AttributeDatatype; at != nil {
			attr.AttributeType = *at
		}
		// could be nil=nil
		attr.AttributeType.NullableAttributeDatatype = each.GetDatatype().NullableAttributeDatatype
		attr.IsNullable = each.IsNullable
		attr.Doc(each.Documentation)
		// copy all
		attr.CopyPropertiesFrom(each.Named)
		each.Extensions.PostBuildAttribute(each, attr)
	}
	return m
}

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
	if c.ColumnType.AttributeDatatype == nil {
		e.Add(c.Named, errors.New("has unknown attribute type for "+c.ColumnType.String()))
	}
}
