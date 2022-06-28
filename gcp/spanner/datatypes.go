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
}

var (
	BOOL      = simple("BOOL")
	BYTES     = simple("BYTES(MAX)")
	DATE      = simple("DATE")
	JSON      = simple("JSON")
	TIMESTAMP = simple("TIMESTAMP")
	INT64     = simple("INT64")
	FLOAT64   = simple("FLOAT64")
	NUMERIC   = simple("NUMERIC") // suitable for financial calculations
	STRING    = simple("STRING(MAX)")
)

func String(max int) dtType {
	return dtType{
		Named:      core.N("spanner.Datatype", fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}
}

func Array(elementType dtType) dtType {
	return dtType{
		Named: core.N("spanner.Datatype", fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}
}
