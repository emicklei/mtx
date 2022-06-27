package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import "github.com/emicklei/mtx/core"

func simple(typename string) core.Datatype[DatatypeExtensions] {
	return core.Datatype[DatatypeExtensions]{
		Named: core.N("spanner.Datatype", typename),
	}
}

var BigInteger = core.Datatype[DatatypeExtensions]{
	Named:      core.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}

var (
	BYTES     = simple("BYTES")
	ARRAY     = simple("ARRAY")
	DATE      = simple("DATE")
	TIMESTAMP = simple("TIMESTAMP")
	JSON      = simple("JSON")
	INT64     = simple("INT64")
	STRING    = simple("STRING")
)
