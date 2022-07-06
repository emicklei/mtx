package pg

import (
	"github.com/emicklei/mtx"
)

// BEGIN: copy from datatypes.go.template
type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType, isUserDefined bool) DType {
	dt := DType{
		Named:         mtx.N("pg.Datatype", typename),
		IsUserDefined: isUserDefined,
	}.WithAttributeType(at)
	return registry.Add(dt)
}

func RegisterType(typename string, at mtx.AttributeType) DType {
	return register(typename, at, mtx.UserDefinedType)
}

func MappedAttributeType(at mtx.AttributeType) DType {
	return registry.MappedAttributeType(at)
}

func Type(name string) DType {
	dt, ok := registry.TypeNamed(name)
	if ok {
		return dt
	}
	return register(name, mtx.UNKNOWN, mtx.UserDefinedType)
}

// END: copy from datatypes.go.template

var (
	UNKNOWN = register("ANY", mtx.UNKNOWN, mtx.StandardType)
	STRING  = register("text", mtx.STRING, mtx.StandardType)
	DATE    = register("date", mtx.DATE, mtx.StandardType)
	BYTES   = register("bytes", mtx.BYTES, mtx.StandardType)
)
