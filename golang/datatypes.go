package golang

import "github.com/emicklei/mtx"

var (
	registry            = mtx.NewTypeRegistry("golang.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	Any     = registry.Standard("any", mtx.Unknown)
	String  = registry.Standard("string", mtx.String)
	Bool    = registry.Standard("bool", mtx.Boolean)
	Bytes   = registry.Standard("[]byte", mtx.Bytes)
	Time    = registry.Standard("time.Time", mtx.Timestamp)
	Float32 = registry.Standard("float32", mtx.Float)
	Float64 = registry.Standard("float64", mtx.Float)
	Int     = registry.Standard("int", mtx.Integer)

	BigRat       = registry.Standard("*big.Rat", mtx.Decimal)
	MapStringAny = registry.Standard("map[string]any", mtx.JSON)
)

func init() {
	registry.EncodeAs(mtx.Date, Time)
	registry.EncodeAs(mtx.Timestamp, Time)
	registry.EncodeAs(mtx.DateTime, Time)
	// TODO
	//registry.Trace()
}
