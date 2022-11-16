package mtx

import "testing"

func TestNamedSourceOn(t *testing.T) {
	n := N("clz", "nam").Doc("doc").Set("key", "42")
	if got, want := ToSource(n), `.Doc("doc").Set("key","42")`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestNamesspace(t *testing.T) {
	if got, want := Unknown.Namespace(), "basic"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestNamedCopy(t *testing.T) {
	n1 := N("clz", "nam")
	n2 := n1.Set("k", "v")
	if got, want := n1.Properties == nil, true; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := len(n2.Properties), 1; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
