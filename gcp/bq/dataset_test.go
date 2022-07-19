package bq

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestDataset(t *testing.T) {
	ds := NewDataset("mydataset").Doc("dataset comment")
	tab := ds.Table("mytable").Doc("my table")
	tab.Column("id").Type(BYTES).NotNull().Doc("my id")
	tab.Column("large").Type(BigNumeric(2, 10))
	tab.C("moment", DATE, "location independent moment in time")
	t.Log(mtx.ToJSON(ds))
	t.Log("\n", ToJSONSchema(tab))
}
