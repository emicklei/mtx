package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx"
)

// BEGIN: copy from datatypes.go.template
type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType, isUserDefined bool) DType {
	dt := DType{
		Named:         mtx.N("spanner.Datatype", typename),
		IsUserDefined: isUserDefined,
	}.WithAttributeType(at)
	return registry.Add(dt)
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

var BigInteger = DType{
	Named:      mtx.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithAttributeType(mtx.INTEGER)

var (
	UNKNOWN   = register("ANY", mtx.UNKNOWN, mtx.UserDefinedType)
	BOOL      = register("BOOL", mtx.BOOLEAN, mtx.StandardType)
	BYTES     = register("BYTES(MAX)", mtx.BYTES, mtx.StandardType)
	DATE      = register("DATE", mtx.DATE, mtx.StandardType)
	JSON      = register("JSON", mtx.JSON, mtx.StandardType)
	TIMESTAMP = register("TIMESTAMP", mtx.TIMESTAMP, mtx.StandardType)
	INT64     = register("INT64", mtx.INTEGER, mtx.StandardType)
	FLOAT64   = register("FLOAT64", mtx.FLOAT, mtx.StandardType)
	NUMERIC   = register("NUMERIC", mtx.DECIMAL, mtx.StandardType) // suitable for financial calculations
	STRING    = register("STRING(MAX)", mtx.STRING, mtx.StandardType)
)

func init() {
	INT64.Set("bits", "64")
}

func String(max int) DType {
	return DType{
		Named:      mtx.N("spanner.Datatype", fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithAttributeType(mtx.STRING)
}

func Array(elementType DType) DType {
	return DType{
		Named: mtx.N("spanner.Datatype", fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithAttributeType(mtx.Array(elementType.AttributeType))
}
