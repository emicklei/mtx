package db

import (
	"bytes"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

type Database struct {
	mtx.Named
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
	mtx.Named
	Columns    []*Column    `json:"columns,omitempty"`
	Extensions ExtendsTable `json:"ext"`
}

func (t *Table) Doc(d string) *Table {
	t.Documentation = d
	return t
}

func (t *Table) Set(key string, value any) *Table {
	t.Named = t.Named.Set(key, value)
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
func (t *Table) ToEntity() *basic.Entity {
	m := basic.NewEntity(t.Name)
	m.Named = m.Named.WithPropertiesCopiedFrom(t.Named) // TODO
	// see if property overrides this
	if n, ok := t.Get(basic.EntityName); ok {
		m.Named.Name = n.(string)
	}
	m.Doc(t.Documentation)
	for _, each := range t.Columns {
		attr := m.Attribute(each.Name)
		// see if property overrides this
		if n, ok := each.Get(basic.AttributeName); ok {
			attr.Named.Name = n.(string)
		}
		dt := each.GetDatatype()
		// attr.IsNullable = each.IsNullable
		// for conversion bring the nullable info into the datatype.
		if each.IsNullable {
			dt = dt.WithNullable()
		}
		// convert
		attr.AttributeType = each.Extensions.ToBasicType(dt)
		attr.Doc(each.Documentation)
		// copy all
		attr.Named = attr.Named.WithPropertiesCopiedFrom(each.Named)
		// needed? TODO
		//each.Extensions.PostBuildAttribute(each, attr)
	}
	return m
}
