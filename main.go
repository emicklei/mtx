package main

import (
	"github.com/emicklei/mtx/db"
	"github.com/emicklei/mtx/mapping"
	"github.com/emicklei/mtx/model"
	"github.com/emicklei/mtx/namespace"
	"github.com/emicklei/mtx/proto"
)

func main() {
	// spec a domain entity
	people := namespace.New("people")
	pm := people.Model("Person")
	pm.Attr("id", model.Identifier)
	pm.Attr("name", model.String)
	pm.Attr("age", model.Integer)
	pm.Relation("parents", model.ToMany(pm))
	pm.Relation("children", model.ToMany(pm))

	// spec a proto
	peoplePkg := proto.NewPackage("people")
	personMsg := peoplePkg.NewMessage("Person")
	personMsg.Field("name", proto.String)
	personMsg.Field("children", proto.Repeated(proto.String))

	// spec a entity<->proto mapping
	entityToProto := mapping.New()
	entityToProto.Field(pm.Attr("name"), personMsg.Field("name"))
	entityToProto.Field(pm.Attr("children"), personMsg.Field("children").Project(pm.Attr("id")))

	// spec a database
	peopleTab := db.NewTable("persons")
	peopleTab.Column("id", db.UUID)

	// spec a entity<->db mapping
}
