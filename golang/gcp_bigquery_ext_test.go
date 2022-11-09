package golang

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

func TestInt64(t *testing.T) {
	at := basic.Integer.Set("bits", 64)
	bt := bigQueryTypeMapper(at, false)
	if got, want := bt.Name, "int64"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestJSON(t *testing.T) {
	bqJSON := mtx.Datatype{
		Named:         mtx.N("test", "JSON"),
		BasicDatatype: &basic.JSON,
	}
	dt := bigQueryTypeMapper(bqJSON, false)
	if got, want := dt.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
