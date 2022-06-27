package pg

import (
	"io"

	"github.com/emicklei/mtx/core"
)

type Database struct{}

func (d *Database) Table(name string) *core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab := new(core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = core.N(tab.Extensions.OwnerClass(), name)
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
