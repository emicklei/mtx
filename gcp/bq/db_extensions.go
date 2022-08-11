package bq

import (
	"io"
	"strings"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
	"github.com/emicklei/mtx/golang"
	"github.com/iancoleman/strcase"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) TableExtensions() db.ExtendsTable { return new(TableExtensions) }

func (d DatabaseExtensions) TableClass() string { return "bq.Table" }

type TableExtensions struct {
}

func (t TableExtensions) OwnerClass() string { return "bq.Table" }

func (t TableExtensions) SQLOn(table *db.Table, w io.Writer) {
	io.WriteString(w, "there will be BQ")
}

func (t TableExtensions) ColumnExtensions() db.ExtendsColumn { return new(ColumnExtensions) }

type ColumnExtensions struct {
	NestedColumns []*db.Column
}

func Extensions(c *db.Column) *ColumnExtensions {
	return c.Extensions.(*ColumnExtensions)
}

// e.g identity/TimeInterval
func NestedColumn(t *db.Table, path string, separator string) (*db.Column, bool) {
	tokens := strings.Split(path, separator)
	here := t.Columns
	var found *db.Column
	for _, each := range tokens {
		if c, ok := mtx.FindByName(here, each); ok {
			here = c.Extensions.(*ColumnExtensions).NestedColumns
			found = c
		} else {
			return nil, false
		}
	}
	return found, true
}

func (e *ColumnExtensions) Column(name string) *db.Column {
	if c, ok := mtx.FindByName(e.NestedColumns, name); ok {
		return c
	}
	c := &db.Column{
		Named:      mtx.N("bq.Column", name),
		ColumnType: Unknown,
		Extensions: new(ColumnExtensions),
	}
	e.NestedColumns = append(e.NestedColumns, c)
	return c
}

func (t ColumnExtensions) Datatype() db.ExtendsDatatype { return new(DatatypeExtensions) }

func (t ColumnExtensions) OwnerClass() string { return "bq.Column" }

// TODO this is very Go specific ; should not be here
func (t ColumnExtensions) PostBuildAttribute(c *db.Column, a *mtx.Attribute) {
	// TODO
	if c.ColumnType == Record {
		a.Set(golang.GoTypeName, strcase.ToCamel(c.Name))
		a.Set("IsEntity", true)                       // TODO constant?
		a.Set("AttributeCount", len(t.NestedColumns)) // TODO constant?
	}
}

type DatatypeExtensions struct {
	Max       int64 `json:"max,omitempty"`
	Scale     int   `json:"scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision int   `json:"precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}

func (d DatatypeExtensions) OwnerClass() string { return "bq.Datatype" }
