package bq

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestDataset(t *testing.T) {
	ns := NewNamespace("world")
	ds := ns.Dataset("mydataset")
	tab := ds.Table("mytable")
	_ = tab.Column("id").Datatype(Bytes)
	core.JSONOut(ds)
}
