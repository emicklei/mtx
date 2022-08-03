package tests

import (
	"testing"

	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/pg"
)

func TestGoRowStructFromPostgresTable(t *testing.T) {
	db := pg.NewDatabase("all")
	tab := db.Table("all")
	tab.C("ctext", pg.TEXT, "ctext")
	tab.C("ctext_n", pg.TEXT, "ctext_n").Nullable()
	tab.C("cdate", pg.DATE, "cdate")
	tab.C("cdate_n", pg.DATE, "cdate_n").Nullable()
	e := tab.ToEntity()
	s := golang.ToStruct(e)
	if got, want := s.Field("CdateN").FieldType.Name, "*time.Time"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
