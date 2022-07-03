package bq

import (
	"io"

	"github.com/emicklei/mtx"
)

type bqSpace mtx.Namespace

func NewNamespace(name string) *bqSpace {
	return (*bqSpace)(mtx.NewNamespace(name))
}

func (s bqSpace) Dataset(n string) *Dataset {
	return &Dataset{Named: mtx.N("bq.Dataset", n)}
}

type Dataset struct {
	*mtx.Named
	Tables []*mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] `json:"tables"`
}

func (d *Dataset) Table(name string) *mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions] {
	tab, ok := mtx.FindByName(d.Tables, name)
	if ok {
		return tab
	}
	tab = new(mtx.Table[TableExtensions, ColumnExtensions, DatatypeExtensions])
	tab.Named = mtx.N(tab.Extensions.OwnerClass(), name)
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
	Max       int64 `json:"max,omitempty"`
	Scale     int   `json:"scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision int   `json:"precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}

func (d DatatypeExtensions) OwnerClass() string { return "bq.Datatype" }
