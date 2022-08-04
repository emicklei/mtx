package mtx

import "testing"

func TestNamedSourceOn(t *testing.T) {
	n := N("clz", "nam").Doc("doc").Set("key", 42)
	if got, want := ToSource(n), `.Doc("doc").Set("key",42)`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestNamesspace(t *testing.T) {
	if got, want := String.Namespace(), "mtx"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
