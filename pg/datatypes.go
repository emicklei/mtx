package pg

import (
	"github.com/emicklei/mtx"
)

// BEGIN: copy from datatypes.go.template

var registry = mtx.NewTypeRegistry("pg.Datatype")

func register(typename string, at mtx.Datatype) mtx.Datatype {
	dt := mtx.Datatype{
		Named:             mtx.N("pg.Datatype", typename),
		AttributeDatatype: &at,
	}
	return registry.Add(dt)
}

func RegisterType(typename string, at mtx.Datatype) mtx.Datatype {
	dt := mtx.Datatype{
		Named:             mtx.N("pg.Datatype", typename),
		AttributeDatatype: &at,
		IsUserDefined:     true,
	}
	return registry.Add(dt)
}

// MappedAttributeType returns the mapped proto type for a given attribute type
func MappedAttributeType(at mtx.Datatype) mtx.Datatype {
	return registry.MappedAttributeType(at)
}

func Type(typename string) mtx.Datatype {
	dt, ok := registry.TypeNamed(typename)
	if ok {
		return dt
	}
	return RegisterType(typename, mtx.UNKNOWN)
}

// END: copy from datatypes.go.template

// https://www.postgresql.org/docs/current/datatype.html
var (
	UNKNOWN          = register("ANY", mtx.UNKNOWN)
	BIGINT           = register("bigint", mtx.INTEGER)
	BOOLEAN          = register("boolean", mtx.BOOLEAN)
	BYTEA            = register("bytea", mtx.BYTES)
	DATE             = register("date", mtx.DATE)
	DOUBLE_PRECISION = register("double precision", mtx.DOUBLE)
	FLOAT8           = DOUBLE_PRECISION
	INTEGER          = register("integer", mtx.DOUBLE)
	INT              = INTEGER
	INT4             = INTEGER
	JSON             = register("json", mtx.JSON)
	JSONB            = register("jsonb", mtx.UNKNOWN)
	//INTERVAL         = register("interval", mtx.INTERVAL, mtx.StandardType)
	// TODO
	TEXT       = register("text", mtx.STRING)
	TIMESTAMP  = register("timestamp", mtx.DATETIME)
	TIMESTAMPZ = register("timestampz", mtx.TIMESTAMP)
	UUID       = register("uuid", mtx.UUID)
)
