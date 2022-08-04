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
	Bool      = registry.Standard("BOOL", mtx.Boolean)
	Bytes     = registry.Standard("BYTES(MAX)", mtx.Bytes)
	Date      = registry.Standard("DATE", mtx.Date)
	JSON      = registry.Standard("JSON", mtx.JSON)
	Timestamp = registry.Standard("TIMESTAMP", mtx.Timestamp)
	Int64     = registry.Standard("INT64", mtx.Integer)
	Float64   = registry.Standard("FLOAT64", mtx.Float)
	Numeric   = registry.Standard("NUMERIC", mtx.Decimal) // suitable for financial calculations
	String    = registry.Standard("STRING(MAX)", mtx.String)
)

func init() {
	Int64.Set("bits", "64")
	// define encoding for remaining standard types
	registry.EncodeAs(mtx.Duration, String)
	registry.EncodeAs(mtx.UUID, StringMax(36))
	//registry.EncodeAs(mtx.DATERANGE, Array(DATE))
}

var UNKNOWN = registry.Register("any", true)

var BigInteger = mtx.Datatype{
	Named:      mtx.N(registry.Class(), "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithAttributeDatatype(mtx.Integer)

func StringMax(max int) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N(registry.Class(), fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithAttributeDatatype(mtx.String)
}

func Array(elementType mtx.Datatype) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N(registry.Class(), fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithAttributeDatatype(mtx.Array(*elementType.AttributeDatatype))
}
