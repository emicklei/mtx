package bq

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestDataset(t *testing.T) {
	ds := NewDataset("mydataset").Doc("dataset comment")
	tab := ds.Table("mytable").Doc("my table")
	tab.Column("id").Type(BYTES).Doc("my id")
	tab.Column("large").Type(BigNumeric(2, 10))
	t.Log(mtx.ToJSON(ds))
}
