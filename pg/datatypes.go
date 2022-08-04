package pg

import (
	"github.com/emicklei/mtx"
)

var (
	registry            = mtx.NewTypeRegistry("pg.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

// https://www.postgresql.org/docs/current/datatype.html
var (
	Unknown         = registry.Standard("any", mtx.Unknown)
	BigInt          = registry.Standard("bigint", mtx.Integer)
	Boolean         = registry.Standard("boolean", mtx.Boolean)
	Bytea           = registry.Standard("bytea", mtx.Bytes)
	Date            = registry.Standard("date", mtx.Date)
	DoublePrecision = registry.Standard("double precision", mtx.Double)
	Float8          = DoublePrecision
	Integer         = registry.Standard("integer", mtx.Integer)
	Int             = Integer
	Int4            = Integer
	JSON            = registry.Standard("json", mtx.JSON)
	JSONB           = registry.Standard("jsonb", mtx.Unknown)
	//INTERVAL         = registry.Standard("interval", mtx.INTERVAL, mtx.StandardType)
	// TODO
	Text       = registry.Standard("text", mtx.String)
	Timestamp  = registry.Standard("timestamp", mtx.DateTime)
	Timestampz = registry.Standard("timestampz", mtx.Timestamp)
	UUID       = registry.Standard("uuid", mtx.UUID)
)
