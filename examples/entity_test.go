package main

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func TestPersonModel(t *testing.T) {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person").Doc("A human")
	e.A("firstName", mtx.STRING, "calling name")
	t.Log("\n", mtx.ToJSON(e))

	// create Go struct source from entity
	t.Log("\n", golang.Source(e))
}

func TestOneToManyRelation(t *testing.T) {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person")
	r := p.OneToMany(e, e)
	r.One("parent")
	r.Many("children")
	t.Log("\n", mtx.ToJSON(p))
}

func TestManyToManyRelation(t *testing.T) {
	p := mtx.NewPackage("persons")
	e := p.Entity("Person")
	r := p.ManyToMany(e, e)
	r.Name = "followers"
	t.Log("\n", mtx.ToJSON(p))
}
