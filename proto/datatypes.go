package proto

import (
	"github.com/emicklei/mtx"
)

type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType) DType {
	dt := DType{
		Named: mtx.N("proto.Datatype", typename),
	}.WithAttributeType(at)
	return registry.Add(dt)
}

func TypeNamed(name string) DType {
	dt, ok := registry.TypeNamed(name)
	if ok {
		return dt
	}
	return register(name, mtx.UNKNOWN)
}

func MappedAttributeType(at mtx.AttributeType) DType {
	return registry.MappedAttributeType(at)
}

var (
	UNKNOWN = register("any", mtx.UNKNOWN)
	// DOUBLE  = register(FieldType{Named: mtx.N("proto.FieldType", "double")}.WithAttributeType(mtx.DOUBLE))
	// FLOAT   = register(FieldType{Named: mtx.N("proto.FieldType", "float")}.WithAttributeType(mtx.FLOAT))
	// STRING  = register(FieldType{Named: mtx.N("proto.FieldType", "string")}.WithAttributeType(mtx.STRING))
	INT32 = register("int32", mtx.INTEGER)
	// INT64   = register(FieldType{Named: mtx.N("proto.FieldType", "int64")}.WithAttributeType(mtx.INTEGER)) //.Set("bits", 64))
	// BOOL    = register(FieldType{Named: mtx.N("proto.FieldType", "bool")}.WithAttributeType(mtx.BOOLEAN))
)
