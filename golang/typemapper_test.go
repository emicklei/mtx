package golang

import (
	"testing"

	"github.com/emicklei/mtx/basic"
)

func TestBigQueryTypeMapper(t *testing.T) {
	dt := bigQueryTypeMapper(basic.String, true)
	if got, want := dt.Name, "bigquery.NullString"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	dt = bigQueryTypeMapper(basic.String, false)
	if got, want := dt.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestStandardTypeMapper(t *testing.T) {
	dt := StandardTypeMapper(basic.JSON, false)
	if got, want := dt.Name, MapStringAny.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	dt = StandardTypeMapper(basic.String, false)
	if got, want := dt.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	dt = StandardTypeMapper(basic.String, true)
	if got, want := dt.Name, "*string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	dt = StandardTypeMapper(basic.Bytes, true)
	if got, want := dt.Name, "*[]byte"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
