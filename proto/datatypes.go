package proto

import (
	"github.com/emicklei/mtx"
)

// BEGIN: copy from datatypes.go.template
type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType, isUserDefined bool) DType {
	dt := DType{
		Named:         mtx.N("proto.Datatype", typename),
		IsUserDefined: isUserDefined,
	}.WithAttributeType(at)
	return registry.Add(dt)
}

func MappedAttributeType(at mtx.AttributeType) DType {
	return registry.MappedAttributeType(at)
}

func Type(name string) DType {
	dt, ok := registry.TypeNamed(name)
	if ok {
		return dt
	}
	return register(name, mtx.UNKNOWN, true)
}

// END: copy from datatypes.go.template

var (
	UNKNOWN = register("any", mtx.UNKNOWN, mtx.UserDefinedType)
	DOUBLE  = register("double", mtx.DOUBLE, mtx.StandardType)
	FLOAT   = register("float", mtx.FLOAT, mtx.StandardType)
	STRING  = register("string", mtx.STRING, mtx.StandardType)
	INT32   = register("int32", mtx.INTEGER, mtx.StandardType)
	INT64   = register("int64", mtx.INTEGER, mtx.StandardType) //.Set("bits", 64))
	BOOL    = register("bool", mtx.BOOLEAN, mtx.StandardType)
)
