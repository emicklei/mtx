package pg

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var (
	registry            = mtx.NewTypeRegistry("pg.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

// https://www.postgresql.org/docs/current/datatype.html
var (
	Unknown         = registry.Standard("any", basic.Unknown)
	BigInt          = registry.Standard("bigint", basic.Integer)
	Boolean         = registry.Standard("boolean", basic.Boolean)
	Bytea           = registry.Standard("bytea", basic.Bytes)
	Date            = registry.Standard("date", basic.Date)
	DoublePrecision = registry.Standard("double precision", basic.Double)
	Float8          = DoublePrecision
	Integer         = registry.Standard("integer", basic.Integer)
	Int             = Integer
	Int4            = Integer
	JSON            = registry.Standard("json", basic.JSON)
	JSONB           = registry.Standard("jsonb", basic.Unknown)
	//INTERVAL         = registry.Standard("interval", basic.INTERVAL, basic.StandardType)
	// TODO
	Text       = registry.Standard("text", basic.String)
	Timestamp  = registry.Standard("timestamp", basic.DateTime)
	Timestampz = registry.Standard("timestampz", basic.Timestamp)
	UUID       = registry.Standard("uuid", basic.UUID).Set(mtx.GoTypeName, "pg.UUID")
)
