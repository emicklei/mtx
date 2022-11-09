package bq

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestNullableBQStringEntityString(t *testing.T) {
	tab := NewDataset("test").Table("test")
	tab.C("s", String, "").Nullable()
	ent := tab.ToEntity()
	s := ent.Attributes[0]
	if got, want := s.AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := s.IsNullable, true; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	t.Log(mtx.ToJSON(tab))
}
