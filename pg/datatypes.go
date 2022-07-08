package pg

import (
	"github.com/emicklei/mtx"
)

// BEGIN: copy from datatypes.go.template
type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType, isUserDefined bool) DType {
	dt := DType{
		Named:         mtx.N("pg.Datatype", typename),
		IsUserDefined: isUserDefined,
	}.WithAttributeType(at)
	return registry.Add(dt)
}

func RegisterType(typename string, at mtx.AttributeType) DType {
	return register(typename, at, mtx.UserDefinedType)
}

func MappedAttributeType(at mtx.AttributeType) DType {
	return registry.MappedAttributeType(at)
}

func Type(name string) DType {
	dt, ok := registry.TypeNamed(name)
	if ok {
		return dt
	}
	return register(name, mtx.UNKNOWN, mtx.UserDefinedType)
}

// END: copy from datatypes.go.template

// https://www.postgresql.org/docs/current/datatype.html
var (
	UNKNOWN          = register("ANY", mtx.UNKNOWN, mtx.StandardType)
	BIGINT           = register("bigint", mtx.INTEGER, mtx.StandardType)
	BOOLEAN          = register("boolean", mtx.BOOLEAN, mtx.StandardType)
	BYTEA            = register("bytea", mtx.BYTES, mtx.StandardType)
	DATE             = register("date", mtx.DATE, mtx.StandardType)
	DOUBLE_PRECISION = register("double precision", mtx.DOUBLE, mtx.StandardType)
	FLOAT8           = DOUBLE_PRECISION
	INTEGER          = register("integer", mtx.DOUBLE, mtx.StandardType)
	INT              = INTEGER
	INT4             = INTEGER
	JSON             = register("json", mtx.JSON, mtx.StandardType)
	JSONB            = register("jsonb", mtx.UNKNOWN, mtx.StandardType)
	//INTERVAL         = register("interval", mtx.INTERVAL, mtx.StandardType)
	// TODO
	TEXT       = register("text", mtx.STRING, mtx.StandardType)
	TIMESTAMP  = register("timestamp", mtx.DATETIME, mtx.StandardType)
	TIMESTAMPZ = register("timestampz", mtx.TIMESTAMP, mtx.StandardType)
	UUID       = register("uuid", mtx.UUID, mtx.StandardType)
)
