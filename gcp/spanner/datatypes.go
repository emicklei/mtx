package spanner

// https://cloud.google.com/spanner/docs/reference/standard-sql/data-types

import (
	"fmt"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var (
	registry            = mtx.NewTypeRegistry("spanner.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

// these are documented available types
var (
	Bool      = registry.Standard("BOOL", basic.Boolean)
	Bytes     = registry.Standard("BYTES(MAX)", basic.Bytes)
	Date      = registry.Standard("DATE", basic.Date)
	JSON      = registry.Standard("JSON", basic.JSON)
	Timestamp = registry.Standard("TIMESTAMP", basic.Timestamp)
	Int64     = registry.Standard("INT64", basic.Integer)
	Float64   = registry.Standard("FLOAT64", basic.Float)
	Numeric   = registry.Standard("NUMERIC", basic.Decimal) // suitable for financial calculations
	String    = registry.Standard("STRING(MAX)", basic.String)
)

func init() {
	Int64.Set("bits", "64")
	// define encoding for remaining standard types
	registry.EncodeAs(basic.Duration, String)
	registry.EncodeAs(basic.UUID, StringMax(36))
	//registry.EncodeAs(mtx.DATERANGE, Array(DATE))
}

var UNKNOWN = registry.Register("any", true)

var BigInteger = mtx.Datatype{
	Named:      mtx.N(registry.Class(), "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}.WithAttributeDatatype(basic.Integer)

func StringMax(max int) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N(registry.Class(), fmt.Sprintf("STRING(%d)", max)),
		Extensions: DatatypeExtensions{Max: int64(max)},
	}.WithAttributeDatatype(basic.String)
}

func Array(elementType mtx.Datatype) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N(registry.Class(), fmt.Sprintf("ARRAY(%s)", elementType.Name)),
	}.WithAttributeDatatype(mtx.Array(*elementType.AttributeDatatype))
}
