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
	tab.Column("large").Datatype(BigNumeric(2, 10))
	core.JSONOut(ds)
}
