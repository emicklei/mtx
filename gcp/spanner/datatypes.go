package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx"
)

// BEGIN: copy from datatypes.go.template

var registry = mtx.NewTypeRegistry("spanner.Datatype")

func register(typename string, at mtx.AttributeType) mtx.Datatype {
	return registry.Register(typename, at, false)
}

func RegisterType(typename string, at mtx.AttributeType) mtx.Datatype {
	return registry.Register(typename, at, true)
}

func MappedAttributeType(at mtx.AttributeType) mtx.Datatype {
	return registry.MappedAttributeType(at)
}

func Type(typename string) mtx.Datatype {
	dt, ok := registry.TypeNamed(typename)
	if ok {
		return dt
	}
	return registry.Register(typename, mtx.UNKNOWN, true)
}

// END: copy from datatypes.go.template

// these are documented available types
var (
	BOOL      = register("BOOL", mtx.BOOLEAN)
	BYTES     = register("BYTES(MAX)", mtx.BYTES)
	DATE      = register("DATE", mtx.DATE)
	JSON      = register("JSON", mtx.JSON)
	TIMESTAMP = register("TIMESTAMP", mtx.TIMESTAMP)
	INT64     = register("INT64", mtx.INTEGER)
	FLOAT64   = register("FLOAT64", mtx.FLOAT)
	NUMERIC   = register("NUMERIC", mtx.DECIMAL) // suitable for financial calculations
	STRING    = register("STRING(MAX)", mtx.STRING)
)

func init() {
	INT64.Set("bits", "64")
	// define encoding for remaining standard types
	registry.EncodeAs(mtx.DURATION, STRING)
	registry.EncodeAs(mtx.UUID, String(32))
	//registry.EncodeAs(mtx.DATERANGE, Array(DATE))
}

var UNKNOWN = registry.Register("ANY", mtx.UNKNOWN, true)

var BigInteger = mtx.Datatype{
	Named:      mtx.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithAttributeType(mtx.INTEGER)

func String(max int) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N("spanner.Datatype", fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithAttributeType(mtx.STRING)
}

func Array(elementType mtx.Datatype) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N("spanner.Datatype", fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithAttributeType(mtx.Array(elementType.AttributeType))
}
