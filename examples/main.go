package main

import (
	"encoding/json"
	"os"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/model"
)

func main() {
	world := mtx.NewWorld()
	people := world.Namespace("people")
	// spec a domain entity
	pm := people.Model("Person")
	pm.Attr("id", model.Identifier)
	pm.Attr("name", model.String)
	pm.Attr("age", model.Integer)
	pm.Relation("parents", model.ToMany(pm))
	pm.Relation("children", model.ToMany(pm))
	json.NewEncoder(os.Stdout).Encode(pm)

	/**
	// spec a proto
	peoplePkg := world.ProtoPackage("people")
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
	**/
}
