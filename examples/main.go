package main

import (
	"github.com/emicklei/mtx/core"
	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/proto"
)

func main() {
	// world := mtx.NewWorld()
	// people := world.Namespace("people")
	// // spec a domain entity
	// pm := people.Model("Person")
	// pm.Attr("id", model.Identifier)
	// pm.Attr("name", model.String)
	// pm.Attr("age", model.Integer)
	// pm.Relation("parents", model.ToMany(pm))
	// pm.Relation("children", model.ToMany(pm))
	// json.NewEncoder(os.Stdout).Encode(pm)

	// spec a proto
	pkg := proto.NewPackage("people")
	msg := pkg.Message("Person")
	msg.Field("name").FieldType(proto.String)
	msg.Field("children").FieldType(proto.String)

	/**
	// spec a entity<->proto mapping
	entityToProto := mapping.New()
	entityToProto.Field(pm.Attr("name"), msg.Field("name"))
	entityToProto.Field(pm.Attr("children"), personMsg.Field("children").Project(pm.Attr("id")))

	// spec a database
	peopleTab := db.NewTable("persons")
	peopleTab.Column("id", db.UUID)

	// spec a entity<->db mapping
	**/
	{
		ns := bq.NewNamespace("world")
		ds := ns.Dataset("mydataset")
		tab := ds.Table("mytable")
		_ = tab.Column("id").Datatype(bq.Bytes)
		core.JSONOut(ds)
	}
}
