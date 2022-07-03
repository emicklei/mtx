package spanner

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

type Database struct {
	*mtx.Named
	Tables []*mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] `json:"tables"`
}

func (d *Database) Table(name string) *mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab, ok := mtx.FindByName(d.Tables, name)
	if ok {
		return tab
	}
	tab = new(mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = mtx.N(tab.Extensions.OwnerClass(), name)
	d.Tables = append(d.Tables, tab)
	return tab
}

func (d *Database) Doc(doc string) *Database {
	d.Documentation = doc
	return d
}

var _ mtx.ExtendsTable = TableExtensions{}

type TableExtensions struct {
	Interleave any
}

func (t TableExtensions) OwnerClass() string { return "spanner.Table" }

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

func (t ColumnExtensions) OwnerClass() string { return "spanner.Column" }

var _ mtx.ExtendsDatatype = DatatypeExtensions{}

type DatatypeExtensions struct {
	Max int64 `json:"max,omitempty"`
}

func (d DatatypeExtensions) OwnerClass() string { return "spanner.Datatype" }
