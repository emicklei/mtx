package spanner

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) Table() db.ExtendsTable { return new(TableExtensions) }

func (d DatabaseExtensions) TableClass() string { return "spanner.Table" }

var _ db.ExtendsTable = TableExtensions{}

type TableExtensions struct {
	Interleave any
}

func (t TableExtensions) OwnerClass() string { return "spanner.Table" }

func (t TableExtensions) Column() db.ExtendsColumn { return new(ColumnExtensions) }

func (t TableExtensions) SQLOn(table *db.Table, w io.Writer) {
	// Spanner DDL does not take commments
	// fmt.Fprintf(w, "-- %s\n", table.Documentation)
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
		// TODO how to handle extension driven comment writing
		// Cloud Spanner itself does not accept these as a part of a DDL-statement.
		fmt.Fprintf(w, "%s %s", each.Name, each.ColumnType.Name)
		if !each.IsNullable {
			fmt.Fprint(w, " NOT NULL")
		}
		fmt.Fprintln(w)
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

var _ db.ExtendsColumn = ColumnExtensions{}

type ColumnExtensions struct {
	IsComplex bool
}

func (t ColumnExtensions) Datatype() db.ExtendsDatatype { return new(DatatypeExtensions) }

func (t ColumnExtensions) OwnerClass() string { return "spanner.Column" }

func (t ColumnExtensions) PostBuildAttribute(c *db.Column, a *mtx.Attribute) {}

var _ db.ExtendsDatatype = DatatypeExtensions{}

type DatatypeExtensions struct {
	Max int64 `json:"max,omitempty"`
}

func (d DatatypeExtensions) OwnerClass() string { return "spanner.Datatype" }
