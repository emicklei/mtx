package pg

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

type Table = mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions]

type Database struct{}

func (d *Database) Table(name string) *mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab := new(mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = mtx.N(tab.Extensions.OwnerClass(), name)
	return tab
}

type TableExtensions struct {
}

func (t TableExtensions) OwnerClass() string { return "pg.Table" }

func (t TableExtensions) SQLOn(table any, w io.Writer) {
	// we know its actual type
	tab := table.(*mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
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

type ColumnExtensions struct {
}

func (t ColumnExtensions) OwnerClass() string { return "pg.Column" }

type DatatypeExtensions struct {
}

func (d DatatypeExtensions) OwnerClass() string { return "pg.Datatype" }
