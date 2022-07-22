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
