package tests

import (
	"fmt"
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func ExampleEntity() {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person").Doc("A human")
	e.A("firstName", mtx.String, "calling name")

	// create Go struct source from entity
	fmt.Println(golang.ToStruct(e).Go())
	// Output:
	// // Person : A human
	// type Person struct {
	// 	FirstName string // calling name
	// }
}

func TestMAP_STRING_ANY(t *testing.T) {
	e := mtx.NewEntity("Test")
	e.A("m", mtx.JSON, "map")
	s := golang.ToStruct(e)
	f := s.Fields[0]
	if got, want := f.FieldType.Name, "map[string]any"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestTIMESTAMP_TIME(t *testing.T) {
	e := mtx.NewEntity("Test")
	e.A("t", mtx.Timestamp, "")
	e.A("t_n", mtx.Timestamp, "").Nullable()
	s := golang.ToStruct(e)
	f := s.Fields[0]
	if got, want := f.FieldType.Name, "time.Time"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	f = s.Fields[1]
	if got, want := f.FieldType.Name, "*time.Time"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestCustomTypeCivilDate(t *testing.T) {
	t.Skip() // changes global which breaks others
	golang.RegisterType("civil.Date", mtx.Date)

	e := mtx.NewEntity("Test")
	e.A("c", mtx.Date, "")
	e.A("c_n", mtx.Date, "").Nullable()
	s := golang.ToStruct(e)
	f := s.Fields[0]
	if got, want := f.FieldType.Name, "civil.Date"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	f = s.Fields[1]
	if got, want := f.FieldType.Name, "*civil.Date"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
