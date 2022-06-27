package spanner

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx/core"
)

type Database struct {
	*core.Named
	Tables []*core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions]
}

func (d *Database) Table(name string) *core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab, ok := core.FindByName(d.Tables, name)
	if ok {
		return tab
	}
	tab = new(core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = core.N(tab.Extensions.OwnerClass(), name)
	d.Tables = append(d.Tables, tab)
	return tab
}

func (d *Database) Doc(doc string) *Database {
	d.Documentation = doc
	return d
}

type TableExtensions struct {
	Key string
}

func (t TableExtensions) OwnerClass() string { return "spanner.Table" }

func (t TableExtensions) SQLOn(table any, w io.Writer) {
	// we know its actual type
	tab := table.(*core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	fmt.Fprintf(w, "CREATE TABLE %s (\n", tab.Name)
	prims := []string{}
	for _, each := range tab.Columns {
		if each.IsPrimary {
			prims = append(prims, each.Name)
		}
		each.SQLOn(w)
	}
	fmt.Fprint(w, ") PRIMARY KEY (\n")
	for _, each := range prims {
		fmt.Fprintf(w, "\t%s\n", each)
	}
	fmt.Fprintf(w, ")")
}

type ColumnExtensions struct {
	IsComplex bool
}

func (t ColumnExtensions) OwnerClass() string { return "spanner.Column" }

type DatatypeExtensions struct {
	Max int64 `json:"Max,omitempty"`
}

func (d DatatypeExtensions) OwnerClass() string { return "spanner.Datatype" }
