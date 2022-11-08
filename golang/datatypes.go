package golang

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var (
	registry            = mtx.NewTypeRegistry("golang.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	Any     = registry.Standard("any", basic.Unknown)
	String  = registry.Standard("string", basic.String)
	Bool    = registry.Standard("bool", basic.Boolean)
	Bytes   = registry.Standard("[]byte", basic.Bytes)
	Time    = registry.Standard("time.Time", basic.Timestamp)
	Float32 = registry.Standard("float32", basic.Float)
	Float64 = registry.Standard("float64", basic.Float)
	Int     = registry.Standard("int", basic.Integer)
	Int64   = registry.Standard("int64", basic.Integer.Set("bits", 64))

	BigRat       = registry.Standard("*big.Rat", basic.Decimal)
	MapStringAny = registry.Standard("map[string]any", basic.JSON)
)

func init() {
	registry.EncodeAs(basic.Date, Time)
	registry.EncodeAs(basic.Timestamp, Time)
	registry.EncodeAs(basic.DateTime, Time)
	// TODO
	//registry.Trace()
}
