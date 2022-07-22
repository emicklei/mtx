package bq

import (
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
	"github.com/emicklei/mtx/golang"
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

// TODO this is very Go specific ; should not be here
func (t ColumnExtensions) PostBuildAttribute(c *db.Column, a *mtx.Attribute) {
	// TEMP TODO
	if c.IsNullable && a.AttributeType == mtx.BYTES {
		a.Set(golang.GoTypeName, "[]byte")
	}
	if c.IsNullable && a.AttributeType == mtx.STRING {
		a.Set(golang.GoTypeName, "bigquery.NullString")
	}
	if a.AttributeType == mtx.JSON {
		if c.IsNullable {
			a.Set(golang.GoTypeName, "bigquery.NullString")
		} else {
			a.Set(golang.GoTypeName, "string")
		}
	}
	// END TEMP
	a.Tags = append(a.Tags, mtx.Tag{Name: "bigquery", Value: c.Name})
}

type DatatypeExtensions struct {
	Max       int64 `json:"max,omitempty"`
	Scale     int   `json:"scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision int   `json:"precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}

func (d DatatypeExtensions) OwnerClass() string { return "bq.Datatype" }
