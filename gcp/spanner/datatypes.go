package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx/core"
)

type dtType = core.Datatype[DatatypeExtensions]

func simple(typename string) dtType {
	return dtType{
		Named: core.N("spanner.Datatype", typename),
	}
}

var BigInteger = dtType{
	Named:      core.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithCoreType(core.INTEGER)

var (
	BOOL      = simple("BOOL").WithCoreType(core.BOOLEAN)
	BYTES     = simple("BYTES(MAX)").WithCoreType(core.BYTES)
	DATE      = simple("DATE").WithCoreType(core.DATE)
	JSON      = simple("JSON")
	TIMESTAMP = simple("TIMESTAMP").WithCoreType(core.TIMESTAMP)
	INT64     = simple("INT64").WithCoreType(core.INTEGER)
	FLOAT64   = simple("FLOAT64").WithCoreType(core.FLOAT)
	NUMERIC   = simple("NUMERIC").WithCoreType(core.DECIMAL) // suitable for financial calculations
	STRING    = simple("STRING(MAX)").WithCoreType(core.STRING)
)

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
