package main

import (
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
}
