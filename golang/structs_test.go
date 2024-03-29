package golang

import (
	"fmt"
	"strings"
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

func TestStructGoSource(t *testing.T) {
	s := new(Struct)
	s.Named = mtx.N("golang.Struct", "Test")
	if got, want := s.ToGo(), `// Test 
type Test struct {
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
		FieldType: String,
	})
	// space after //
	if got, want := s.ToGo(), `// Test 
type Test struct {
	Test string // 
}
`; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}

}

func TestStructFull(t *testing.T) {
	p := NewPackage("test")
	s := p.Type("Test")
	s.F("Example", String, "some example")
	if got, want := s.ToGo(), `// Test 
type Test struct {
	Example string // some example
}
`; got != want {
		tokenCompare(got, want)
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestStructBuilder(t *testing.T) {
	p := basic.NewPackage("test")
	e := p.Entity("test")
	e.A("name", basic.String, "nameless")
	b := NewStructBuilder(e)
	s := b.Build()
	if got, want := s.ToGo(), `// Test 
type Test struct {
	Name string // nameless
}
`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestStructWithCSVPopulate(t *testing.T) {
	t.Skip()
	p := basic.NewPackage("test")
	e := p.Entity("test")
	e.A("name", basic.String, "required name")
	e.A("null_name", basic.String, "nullable name")
	s := ToStruct(e, WithCSVPopulate)
	if got, want := s.ToGo(), `// Test 
type Test struct {
	Name string // required name
	NullName *string // nullable name
}

func (r Test) CSVPopulate(record []string,offset int) (Test, error) {
	if v := record[offset+0]; v != "" {
		r.Name = v
	}
	if v := record[offset+1]; v != "" {
		r.NullName = StringToPtrString(v)
	}
	return r, nil
}

`; got != want {
		tokenCompare(got, want)
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func tokenCompare(l, r string) {
	left := strings.Split(l, "\n")
	right := strings.Split(r, "\n")
	for i := 0; i < len(left) && i < len(right); i++ {
		ls, rs := left[i], right[i]
		if ls != rs {
			fmt.Printf("%d:[%s][%s]\n", i, ls, rs)
		}
	}
}
