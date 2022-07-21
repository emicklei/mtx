package pg

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

type DatabaseExtensions struct{}

func (d *DatabaseExtensions) Table() db.ExtendsTable { return new(TableExtensions) }

type TableExtensions struct{}

func (t TableExtensions) OwnerClass() string { return "pg.Table" }

func (t TableExtensions) Column() db.ExtendsColumn { return new(ColumnExtensions) }

func (t TableExtensions) SQLOn(tab *db.Table, w io.Writer) {
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

func (t ColumnExtensions) Datatype() db.ExtendsDatatype { return new(DatatypeExtensions) }

func (t ColumnExtensions) PostBuildAttribute(c *db.Column, a *mtx.Attribute) {}

type DatatypeExtensions struct{}

func (d DatatypeExtensions) OwnerClass() string { return "pg.Datatype" }
