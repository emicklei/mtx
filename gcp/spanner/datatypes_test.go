package spanner

import (
	"testing"

	"github.com/emicklei/mtx/basic"
)

func TestCoreMapping(t *testing.T) {
	if got, want := Bool.BasicDatatype.Name, "boolean"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := String.BasicDatatype.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestMappedAttributeType(t *testing.T) {
	st := MappedAttributeType(basic.String)
	if got, want := st, String; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	{
		st := MappedAttributeType(basic.JSON)
		if got, want := st, JSON; got != want {
			t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
		}
	}
}
