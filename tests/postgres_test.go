package tests

import (
	"testing"

	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/pg"
)

func TestGoRowStructFromPostgresTable(t *testing.T) {
	db := pg.NewDatabase("all")
	tab := db.Table("all")
	tab.C("ctext", pg.Text, "ctext")
	tab.C("ctext_n", pg.Text, "ctext_n").Nullable()
	tab.C("cdate", pg.Date, "cdate")
	tab.C("cdate_n", pg.Date, "cdate_n").Nullable()
	e := tab.ToEntity()
	s := golang.ToStruct(e)
	if got, want := s.Field("CdateN").FieldType.Name, "*time.Time"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
