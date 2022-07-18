package pg

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) Table() mtx.ExtendsTable { return new(TableExtensions) }

func (d DatabaseExtensions) TableClass() string { return "pg.Table" }

type TableExtensions struct{}

func (t TableExtensions) OwnerClass() string { return "pg.Table" }

func (t TableExtensions) Column() mtx.ExtendsColumn { return new(ColumnExtensions) }

func (t TableExtensions) SQLOn(tab *mtx.Table, w io.Writer) {
	// we know its actual type
	fmt.Fprintf(w, "CREATE TABLE %s (\n", tab.Name)
	prims := []string{}
	for i, each := range tab.Columns {
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
	fmt.Fprint(w, ")\n")
}

type ColumnExtensions struct{}

func (t ColumnExtensions) OwnerClass() string { return "pg.Column" }

func (t ColumnExtensions) Datatype() mtx.ExtendsDatatype { return new(DatatypeExtensions) }

type DatatypeExtensions struct{}

func (d DatatypeExtensions) OwnerClass() string { return "pg.Datatype" }
