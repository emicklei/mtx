package proto

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var (
	registry            = mtx.NewTypeRegistry("proto.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	Unknown = registry.Standard("any", mtx.Unknown) // bytes is the fallback
	Double  = registry.Standard("double", basic.Double)
	Float   = registry.Standard("float", basic.Float)
	String  = registry.Standard("string", basic.String)
	Int32   = registry.Standard("int32", basic.Integer)
	Int64   = registry.Standard("int64", basic.Integer) //.Set("bits", 64))
	Bool    = registry.Standard("bool", basic.Boolean)
	Bytes   = registry.Standard("bytes", basic.Bytes)
)
