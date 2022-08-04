package csv

import (
	"testing"
)

func TestFixture(t *testing.T) {
	s, err := ScanSheet("fixture.csv")
	if err != nil {
		t.Fatal(err)
	}
	tab := s.Tab("main")
	if got, want := tab.Columns[0].ColumnType, Boolean; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := tab.Columns[1].ColumnType, Number; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := tab.Columns[2].ColumnType, String; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := tab.Columns[5].ColumnType, Number; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
