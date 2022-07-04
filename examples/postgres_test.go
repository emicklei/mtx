package main

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/pg"
)

func TestPostgresTable(t *testing.T) {
	db := new(pg.Database)
	tab := db.Table("persons")
	tab.C("id", pg.STRING, "identifier of a person")
	tab.C("birthDay", pg.DATE, "day of birth")
	t.Log("\n", mtx.ToJSON(tab))

	// create SQL to create table
	t.Log("\n", tab.SQL())

	// create entity from table
	e := tab.ToEntity()
	t.Log("\n", mtx.ToJSON(e))

	// create Go struct source from entity
	t.Log("\n", golang.Source(e))
}
