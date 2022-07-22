package golang

import (
	"strings"
	"testing"

	"github.com/emicklei/mtx"
)

func TestStructGoSource(t *testing.T) {
	s := new(Struct)
	s.Named = mtx.N("golang.Struct", "Test")
	if got, want := s.Go(), `type Test struct {
}
`; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}

}

func TestStructWithFieldsGoSource(t *testing.T) {
	s := new(Struct)
	s.Named = mtx.N("golang.Struct", "Test")
	s.Fields = append(s.Fields, &Field{
		Named:     mtx.N("golang.Field", "Test"),
		FieldType: STRING,
	})
	// space after //
	if got, want := s.Go(), `type Test struct {
	Test string // 
}
`; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}

}

func TestStructFull(t *testing.T) {
	p := NewPackage("test")
	s := p.Type("Test")
	s.F("Example", STRING, "some example")
	if got, want := s.Go(), `type Test struct {
	Example string // some example
}
`; got != want {
		t.Log(flatten(got))
		t.Log(flatten(want))
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

// reaplce tabs and newlines and spaces
func flatten(s string) string {
	return strings.Replace((strings.Replace(s, "\n", "(n)", -1)), "\t", "(t)", -1)
}
