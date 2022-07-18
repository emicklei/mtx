package spanner

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) Table() mtx.ExtendsTable { return new(TableExtensions) }

func (d DatabaseExtensions) TableClass() string { return "spanner.Table" }

var _ mtx.ExtendsTable = TableExtensions{}

type TableExtensions struct {
	Interleave any
}

func (t TableExtensions) OwnerClass() string { return "spanner.Table" }

func (t TableExtensions) Column() mtx.ExtendsColumn { return new(ColumnExtensions) }

func (t TableExtensions) SQLOn(table *mtx.Table, w io.Writer) {
	fmt.Fprintf(w, "CREATE TABLE %s (\n", table.Name)
	prims := []string{}
	for i, each := range table.Columns {
		if each.IsPrimary {
			prims = append(prims, each.Name)
		}
		if i > 0 {
			fmt.Fprintf(w, ",\t")
		} else {
			fmt.Fprintf(w, " \t")
		}
		each.SQLOn(w)
	}
	fmt.Fprint(w, ") PRIMARY KEY (\n")
	for i, each := range prims {
		if i > 0 {
			fmt.Fprintf(w, ",\t")
		} else {
			fmt.Fprintf(w, " \t")
		}
		fmt.Fprintf(w, "%s\n", each)
	}
	fmt.Fprintf(w, ")")
	// TODO check for Interleave
}

var _ mtx.ExtendsColumn = ColumnExtensions{}

type ColumnExtensions struct {
	IsComplex bool
}

func (t ColumnExtensions) Datatype() mtx.ExtendsDatatype { return new(DatatypeExtensions) }

func (t ColumnExtensions) OwnerClass() string { return "spanner.Column" }

var _ mtx.ExtendsDatatype = DatatypeExtensions{}

type DatatypeExtensions struct {
	Max int64 `json:"max,omitempty"`
}

func (d DatatypeExtensions) OwnerClass() string { return "spanner.Datatype" }
