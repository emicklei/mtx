package bq

import (
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) Table() db.ExtendsTable { return new(TableExtensions) }

func (d DatabaseExtensions) TableClass() string { return "bq.Table" }

type TableExtensions struct {
}

func (t TableExtensions) OwnerClass() string { return "bq.Table" }

func (t TableExtensions) SQLOn(table *db.Table, w io.Writer) {
	io.WriteString(w, "there will be BQ")
}

func (t TableExtensions) Column() db.ExtendsColumn { return new(ColumnExtensions) }

type ColumnExtensions struct {
}

func (t ColumnExtensions) Datatype() db.ExtendsDatatype { return new(DatatypeExtensions) }

func (t ColumnExtensions) OwnerClass() string { return "bq.Column" }

func (t ColumnExtensions) ExtendAttribute(c *db.Column, a *mtx.Attribute) {
	// TEMP TODO
	if c.IsNullable && a.AttributeType == mtx.STRING {
		a.Set(mtx.GoTypeName, "bigquery.NullString")
	}
	a.Tags = append(a.Tags, mtx.Tag{Name: "bigquery", Value: c.Name})
}

type DatatypeExtensions struct {
	Max       int64 `json:"max,omitempty"`
	Scale     int   `json:"scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision int   `json:"precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}

func (d DatatypeExtensions) OwnerClass() string { return "bq.Datatype" }
