package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx/core"
)

type dtType = core.Datatype[DatatypeExtensions]

var knownTypes = map[string]dtType{}

func register(dt dtType) dtType {
	knownTypes[dt.Name] = dt
	return dt
}

// MappedAttributeType returns the best matching spanner type.
func MappedAttributeType(at core.AttributeType) dtType {
	for _, each := range knownTypes {
		if each.AttributeType.Equals(at) {
			return each
		}
	}
	// TODO specials
	return STRING
}

func simple(typename string, at core.AttributeType) dtType {
	return register(dtType{
		Named: core.N("spanner.Datatype", typename),
	}.WithCoreType(at))
}

var BigInteger = register(dtType{
	Named:      core.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithCoreType(core.INTEGER))

var (
	UNKNOWN   = simple("ANY", core.UNKNOWN)
	BOOL      = simple("BOOL", core.BOOLEAN)
	BYTES     = simple("BYTES(MAX)", core.BYTES)
	DATE      = simple("DATE", core.DATE)
	JSON      = simple("JSON", core.JSON)
	TIMESTAMP = simple("TIMESTAMP", core.TIMESTAMP)
	INT64     = simple("INT64", core.INTEGER)
	FLOAT64   = simple("FLOAT64", core.FLOAT)
	NUMERIC   = simple("NUMERIC", core.DECIMAL) // suitable for financial calculations
	STRING    = simple("STRING(MAX)", core.STRING)
)

func init() {
	INT64.Set("bits", "64")
}

func String(max int) dtType {
	return dtType{
		Named:      core.N("spanner.Datatype", fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithCoreType(core.STRING)
}

func Array(elementType dtType) dtType {
	return dtType{
		Named: core.N("spanner.Datatype", fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithCoreType(core.Array(elementType.AttributeType))
}
