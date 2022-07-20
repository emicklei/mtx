package spanner

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestCoreMapping(t *testing.T) {
	if got, want := BOOL.AttributeDatatype.Name, "boolean"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := STRING.AttributeDatatype.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestMappedAttributeType(t *testing.T) {
	st := MappedAttributeType(mtx.STRING)
	if got, want := st, STRING; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	{
		st := MappedAttributeType(mtx.JSON)
		if got, want := st, JSON; got != want {
			t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
		}
	}
}
