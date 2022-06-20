package bq

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestDataset(t *testing.T) {
	ns := NewNamespace("world")
	ds := ns.Dataset("mydataset").Doc("dataset comment")
	tab := ds.Table("mytable").Doc("my table")
	_ = tab.Column("id").Datatype(Bytes).Doc("my id")
	tab.Column("large").Datatype(BigNumeric(2, 10))
	core.JSONOut(ds)
}

func TestDataset2(t *testing.T) {
	// generic table with specific colum
	d := new(Dataset2)
	tab := d.Table("Test")
	tab.Column("id")
	core.JSONOut(tab)
}
