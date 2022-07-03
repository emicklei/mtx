package pg

import (
	"io"

	"github.com/emicklei/mtx"
)

type Database struct{}

func (d *Database) Table(name string) *mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab := new(mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = mtx.N(tab.Extensions.OwnerClass(), name)
	return tab
}

type TableExtensions struct {
}

func (t TableExtensions) OwnerClass() string { return "pg.Table" }

func (t TableExtensions) SQLOn(table any, w io.Writer) {}

type ColumnExtensions struct {
}

func (t ColumnExtensions) OwnerClass() string { return "pg.Column" }

type DatatypeExtensions struct {
}

func (d DatatypeExtensions) OwnerClass() string { return "pg.Datatype" }
