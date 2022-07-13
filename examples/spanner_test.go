package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/gcp/spanner"
	"github.com/emicklei/mtx/golang"
)

func TestSpannerTable(t *testing.T) {
	db := new(spanner.Database)
	tab := db.Table("my_table").Doc("Awesome table")
	tab.C("id", spanner.INT64, "identifier of a row").Primary()
	t.Log("\n", mtx.ToJSON(tab))

	// create SQL to create table
	t.Log("\n", tab.SQL())

	// create entity from spanner table
	e := tab.ToEntity()
	t.Log("\n", mtx.ToJSON(e))

	// create Go struct source from entity
	t.Log("\n", golang.Source(e))

	// write to file, read it back
	js := mtx.ToJSON(db)
	fn := "TestSpannerTable.json"
	os.WriteFile(fn, []byte(js), os.ModePerm)
	defer os.Remove(fn)
	data, _ := os.ReadFile(fn)
	db2 := new(spanner.Database)
	json.Unmarshal(data, db2)
	t.Log("\n", mtx.ToJSON(db))

}
