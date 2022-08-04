package spanner

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestCoreMapping(t *testing.T) {
	if got, want := Bool.AttributeDatatype.Name, "boolean"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := String.AttributeDatatype.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestMappedAttributeType(t *testing.T) {
	st := MappedAttributeType(mtx.String)
	if got, want := st, String; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	{
		st := MappedAttributeType(mtx.JSON)
		if got, want := st, JSON; got != want {
			t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
		}
	}
}
