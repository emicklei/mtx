package main

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func TestPersonModel(t *testing.T) {
	e := mtx.NewEntity("Person").Doc("A human")
	e.A("firstName", mtx.STRING, "calling name")
	t.Log("\n", mtx.ToJSON(e))

	// create Go struct source from entity
	t.Log("\n", golang.Source(e))
}
