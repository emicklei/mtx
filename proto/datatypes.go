package proto

import (
	"github.com/emicklei/mtx"
)

var registry = mtx.NewTypeRegistry("proto.Datatype")

func register(typename string, at mtx.AttributeType) mtx.Datatype {
	return registry.Register(typename, at, false)
}

func RegisterType(typename string, at mtx.AttributeType) mtx.Datatype {
	return registry.Register(typename, at, true)
}

// MappedAttributeType returns the mapped proto type for a given attribute type
func MappedAttributeType(at mtx.AttributeType) mtx.Datatype {
	return registry.MappedAttributeType(at)
}

func Type(typename string) mtx.Datatype {
	dt, ok := registry.TypeNamed(typename)
	if ok {
		return dt
	}
	return registry.Register(typename, mtx.UNKNOWN, true)
}

func init() {
}

var (
	UNKNOWN = register("any", mtx.UNKNOWN) // bytes is the fallback
	DOUBLE  = register("double", mtx.DOUBLE)
	FLOAT   = register("float", mtx.FLOAT)
	STRING  = register("string", mtx.STRING)
	INT32   = register("int32", mtx.INTEGER)
	INT64   = register("int64", mtx.INTEGER) //.Set("bits", 64))
	BOOL    = register("bool", mtx.BOOLEAN)
	BYTES   = register("bytes", mtx.BYTES)
)
