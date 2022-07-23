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
	UNKNOWN = registry.Standard("any", mtx.UNKNOWN) // bytes is the fallback
	DOUBLE  = registry.Standard("double", mtx.DOUBLE)
	FLOAT   = registry.Standard("float", mtx.FLOAT)
	STRING  = registry.Standard("string", mtx.STRING)
	INT32   = registry.Standard("int32", mtx.INTEGER)
	INT64   = registry.Standard("int64", mtx.INTEGER) //.Set("bits", 64))
	BOOL    = registry.Standard("bool", mtx.BOOLEAN)
	BYTES   = registry.Standard("bytes", mtx.BYTES)
)
