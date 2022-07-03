package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx"
)

type dtType = mtx.Datatype[DatatypeExtensions]

var knownTypes = map[string]dtType{}

func register(dt dtType) dtType {
	knownTypes[dt.Name] = dt
	return dt
}

// MappedAttributeType returns the best matching spanner type.
func MappedAttributeType(at mtx.AttributeType) dtType {
	for _, each := range knownTypes {
		if each.AttributeType.Equals(at) {
			return each
		}
	}
	// TODO specials
	return STRING
}

func simple(typename string, at mtx.AttributeType) dtType {
	return register(dtType{
		Named: mtx.N("spanner.Datatype", typename),
	}.WithCoreType(at))
}

var BigInteger = register(dtType{
	Named:      mtx.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithCoreType(mtx.INTEGER))

var (
	UNKNOWN   = simple("ANY", mtx.UNKNOWN)
	BOOL      = simple("BOOL", mtx.BOOLEAN)
	BYTES     = simple("BYTES(MAX)", mtx.BYTES)
	DATE      = simple("DATE", mtx.DATE)
	JSON      = simple("JSON", mtx.JSON)
	TIMESTAMP = simple("TIMESTAMP", mtx.TIMESTAMP)
	INT64     = simple("INT64", mtx.INTEGER)
	FLOAT64   = simple("FLOAT64", mtx.FLOAT)
	NUMERIC   = simple("NUMERIC", mtx.DECIMAL) // suitable for financial calculations
	STRING    = simple("STRING(MAX)", mtx.STRING)
)

func init() {
	INT64.Set("bits", "64")
}

func String(max int) dtType {
	return dtType{
		Named:      mtx.N("spanner.Datatype", fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithCoreType(mtx.STRING)
}

func Array(elementType dtType) dtType {
	return dtType{
		Named: mtx.N("spanner.Datatype", fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithCoreType(mtx.Array(elementType.AttributeType))
}
