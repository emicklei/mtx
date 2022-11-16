package mtx

import (
	"testing"
)

func TestDatatype_WithNullable(t *testing.T) {
	s := NewBasicType("string")
	s = s.Set("k", "v")
	n := s.WithNullable()
	if got, want := len(n.Properties), 1; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
