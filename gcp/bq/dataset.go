package bq

import (
	"io"

	"github.com/emicklei/mtx/core"
)

type bqSpace core.Namespace

func NewNamespace(name string) *bqSpace {
	return (*bqSpace)(core.NewNamespace(name))
}

func (s bqSpace) Dataset(n string) *Dataset {
	return &Dataset{Named: core.N("bq.Dataset", n)}
}

type Dataset struct {
	*core.Named
	Tables []*core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions]
}

func (d *Dataset) Table(name string) *core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab, ok := core.FindByName(d.Tables, name)
	if ok {
		return tab
	}
	tab = new(core.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = core.N(tab.Extensions.OwnerClass(), name)
	d.Tables = append(d.Tables, tab)
	return tab
}

func (d *Dataset) Doc(doc string) *Dataset {
	d.Documentation = doc
	return d
}

type TableExtensions struct {
}

func (t TableExtensions) OwnerClass() string { return "bq.Table" }

func (t TableExtensions) SQLOn(table any, w io.Writer) {}

type ColumnExtensions struct {
}

func (t ColumnExtensions) OwnerClass() string { return "bq.Column" }

type DatatypeExtensions struct {
	Max       int64
	Scale     int `json:"Scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision int `json:"Precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}

func (d DatatypeExtensions) OwnerClass() string { return "bq.Datatype" }
