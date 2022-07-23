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
	UNKNOWN          = registry.Standard("ANY", mtx.UNKNOWN)
	BIGINT           = registry.Standard("bigint", mtx.INTEGER)
	BOOLEAN          = registry.Standard("boolean", mtx.BOOLEAN)
	BYTEA            = registry.Standard("bytea", mtx.BYTES)
	DATE             = registry.Standard("date", mtx.DATE)
	DOUBLE_PRECISION = registry.Standard("double precision", mtx.DOUBLE)
	FLOAT8           = DOUBLE_PRECISION
	INTEGER          = registry.Standard("integer", mtx.DOUBLE)
	INT              = INTEGER
	INT4             = INTEGER
	JSON             = registry.Standard("json", mtx.JSON)
	JSONB            = registry.Standard("jsonb", mtx.UNKNOWN)
	//INTERVAL         = registry.Standard("interval", mtx.INTERVAL, mtx.StandardType)
	// TODO
	TEXT       = registry.Standard("text", mtx.STRING)
	TIMESTAMP  = registry.Standard("timestamp", mtx.DATETIME)
	TIMESTAMPZ = registry.Standard("timestampz", mtx.TIMESTAMP)
	UUID       = registry.Standard("uuid", mtx.UUID)
)
