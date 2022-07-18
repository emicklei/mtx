package bq

import (
	"io"

	"github.com/emicklei/mtx"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) Table() mtx.ExtendsTable { return new(TableExtensions) }

func (d DatabaseExtensions) TableClass() string { return "bq.Table" }

type TableExtensions struct {
}

func (t TableExtensions) OwnerClass() string { return "bq.Table" }

func (t TableExtensions) SQLOn(table *mtx.Table, w io.Writer) {}

func (t TableExtensions) Column() mtx.ExtendsColumn { return new(ColumnExtensions) }

type ColumnExtensions struct {
}

func (t ColumnExtensions) Datatype() mtx.ExtendsDatatype { return new(DatatypeExtensions) }

func (t ColumnExtensions) OwnerClass() string { return "bq.Column" }

type DatatypeExtensions struct {
	Max       int64 `json:"max,omitempty"`
	Scale     int   `json:"scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision int   `json:"precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}

func (d DatatypeExtensions) OwnerClass() string { return "bq.Datatype" }
