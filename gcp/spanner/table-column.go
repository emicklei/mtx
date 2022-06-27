package spanner

import "github.com/emicklei/mtx/core"

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

type ColumnExtensions struct {
	IsComplex bool
}

func (t ColumnExtensions) OwnerClass() string { return "spanner.Column" }

type DatatypeExtensions struct {
	Max int64 `json:"Max,omitempty"`
}

func (d DatatypeExtensions) OwnerClass() string { return "spanner.Datatype" }
