package spanner

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func TestSpannerTable(t *testing.T) {
	db := NewDatabase("testdb")
	tab := db.Table("my_table").Doc("Awesome table")
	tab.C("id", Int64, "identifier of a row").Primary()
	mtx.ToJSON(tab)

	// create SQL to create table
	tab.ToSQL()

	// create entity from spanner table
	e := tab.ToEntity()
	mtx.ToJSON(e)

	// create Go struct source from entity
	golang.ToStruct(e).ToGo()

	// write to file, read it back
	js := mtx.ToJSON(db)
	fn := "TestSpannerTable.json"
	os.WriteFile(fn, []byte(js), os.ModePerm)
	defer os.Remove(fn)
	data, _ := os.ReadFile(fn)
	db2 := NewDatabase("testdb2")
	json.Unmarshal(data, db2)
	t.Log("\n", mtx.ToJSON(db))

}
