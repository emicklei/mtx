package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx"
)

var (
	registry            = mtx.NewTypeRegistry("spanner.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

// these are documented available types
var (
	BOOL      = registry.Standard("BOOL", mtx.BOOLEAN)
	BYTES     = registry.Standard("BYTES(MAX)", mtx.BYTES)
	DATE      = registry.Standard("DATE", mtx.DATE)
	JSON      = registry.Standard("JSON", mtx.JSON)
	TIMESTAMP = registry.Standard("TIMESTAMP", mtx.TIMESTAMP)
	INT64     = registry.Standard("INT64", mtx.INTEGER)
	FLOAT64   = registry.Standard("FLOAT64", mtx.FLOAT)
	NUMERIC   = registry.Standard("NUMERIC", mtx.DECIMAL) // suitable for financial calculations
	STRING    = registry.Standard("STRING(MAX)", mtx.STRING)
)

func init() {
	INT64.Set("bits", "64")
	// define encoding for remaining standard types
	registry.EncodeAs(mtx.DURATION, STRING)
	registry.EncodeAs(mtx.UUID, String(36))
	//registry.EncodeAs(mtx.DATERANGE, Array(DATE))
}

var UNKNOWN = registry.Register("ANY", true)

var BigInteger = mtx.Datatype{
	Named:      mtx.N(registry.Class(), "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithAttributeDatatype(mtx.INTEGER)

func String(max int) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N(registry.Class(), fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithAttributeDatatype(mtx.STRING)
}

func Array(elementType mtx.Datatype) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N(registry.Class(), fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithAttributeDatatype(mtx.Array(*elementType.AttributeDatatype))
}
