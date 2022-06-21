package bq

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestDataset(t *testing.T) {
	ns := NewNamespace("myproject")
	ds := ns.Dataset("mydataset")
	ds.Doc("dataset comment")
	tab := ds.Table("mytable")
	tab.Doc("my table")
	_ = tab.Column("id").Datatype(Bytes).Doc("my id")
	tab.Column("large").Datatype(BigNumeric(2, 10))
	core.JSONOut(ds)
}
