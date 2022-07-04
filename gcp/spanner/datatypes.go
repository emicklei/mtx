package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx"
)

type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType) DType {
	dt := DType{
		Named: mtx.N("pg.Datatype", typename),
	}.WithAttributeType(at)
	return registry.Add(dt)
}

func MappedAttributeType(at mtx.AttributeType) DType {
	return registry.MappedAttributeType(at)
}

var BigInteger = DType{
	Named:      mtx.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithAttributeType(mtx.INTEGER)

var (
	UNKNOWN   = register("ANY", mtx.UNKNOWN)
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
