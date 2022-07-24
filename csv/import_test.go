package csv

import (
	std "encoding/csv"
	"io"
	"os"
	"testing"
)

func TestFixture(t *testing.T) {
	s := NewSheet("test")
	tab := s.Tab("test")

	f, _ := os.Open("fixture.csv")
	defer f.Close()
	r := std.NewReader(f)
	gotNames := false
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if !gotNames {
			for _, each := range record {
				tab.Column(each).Type(UNKNOWN)
			}
			gotNames = true
		} else {
			// resolve unknown type
			for i, each := range record {
				typ := tab.Columns[i]
				if typ.GetDatatype() == UNKNOWN {
					typ.Type(DetectType(each))
				}
			}
		}
		//t.Log(record)
	}
	//t.Log("\n", mtx.ToJSON(s))
	if got, want := tab.Columns[0].ColumnType, BOOLEAN; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := tab.Columns[1].ColumnType, NUMBER; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := tab.Columns[2].ColumnType, STRING; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := tab.Columns[5].ColumnType, NUMBER; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
