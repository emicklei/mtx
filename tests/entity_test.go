package tests

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func TestPersonModel(t *testing.T) {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person").Doc("A human")
	e.A("firstName", mtx.STRING, "calling name")

	// create Go struct source from entity
	t.Log("\n", golang.ToStruct(e).Go())
}

func TestOneToManyRelation(t *testing.T) {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person")
	r := p.OneToMany(e, e)
	r.Left("parent")
	r.Right("children")
	t.Log("\n", mtx.ToJSON(p))
}

func TestManyToManyRelation(t *testing.T) {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person")
	r := p.ManyToMany(e, e)
	r.Name = "followers"
	t.Log("\n", mtx.ToJSON(p))
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
	e.A("t", mtx.TIMESTAMP, "")
	e.A("t_n", mtx.TIMESTAMP, "").Nullable()
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
	golang.RegisterType("civil.Date", mtx.DATE)
	defer golang.RegisterType("time.Time", mtx.DATE)

	e := mtx.NewEntity("Test")
	e.A("c", mtx.DATE, "")
	e.A("c_n", mtx.DATE, "").Nullable()
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
