package proto

import (
	"github.com/emicklei/mtx"
)

var (
	registry            = mtx.NewTypeRegistry("proto.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	Unknown = registry.Standard("any", mtx.Unknown) // bytes is the fallback
	Double  = registry.Standard("double", mtx.Double)
	Float   = registry.Standard("float", mtx.Float)
	String  = registry.Standard("string", mtx.String)
	Int32   = registry.Standard("int32", mtx.Integer)
	Int64   = registry.Standard("int64", mtx.Integer) //.Set("bits", 64))
	Bool    = registry.Standard("bool", mtx.Boolean)
	Bytes   = registry.Standard("bytes", mtx.Bytes)
)
