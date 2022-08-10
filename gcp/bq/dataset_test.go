package bq

import (
	"testing"
)

func TestDataset(t *testing.T) {
	ds := NewDataset("mydataset").Doc("dataset comment")
	tab := ds.Table("mytable").Doc("my table")
	tab.Column("id").Type(Bytes).Doc("my id")
	tab.Column("large").Type(BigNumeric(2, 10))
	tab.C("moment", Date, "location independent moment in time")
	t.Log("\n", ToJSONSchema(tab))
}
