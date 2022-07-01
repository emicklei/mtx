package spanner

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestCoreMapping(t *testing.T) {
	if got, want := BOOL.AttributeType.Name, "boolean"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	if got, want := STRING.AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestMappedAttributeType(t *testing.T) {
	st := MappedAttributeType(core.STRING)
	if got, want := st, STRING; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
