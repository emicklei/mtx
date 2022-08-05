package golang

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestBigQueryTypeMapper(t *testing.T) {
	dt := BigQueryTypeMapper(mtx.String, true)
	if got, want := dt.Name, "bigquery.NullString"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	dt = BigQueryTypeMapper(mtx.String, false)
	if got, want := dt.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestStandardTypeMapper(t *testing.T) {
	dt := StandardTypeMapper(mtx.JSON, false)
	if got, want := dt.Name, MapStringAny.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
