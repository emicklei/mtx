package proto

import "github.com/emicklei/mtx/core"

var (
	UNKNOWN = register(FieldType{Named: core.N("proto.Field", "any")})
	DOUBLE  = register(FieldType{Named: core.N("proto.Field", "double")}.WithCoreType(core.DOUBLE))
	FLOAT   = register(FieldType{Named: core.N("proto.Field", "float")})
	STRING  = register(FieldType{Named: core.N("proto.Field", "string")})
	INT32   = register(FieldType{Named: core.N("proto.Field", "int32")})
	INT64   = register(FieldType{Named: core.N("proto.Field", "int64")}.WithCoreType(core.INTEGER)) //.Set("bits", 64))
	BOOL    = register(FieldType{Named: core.N("proto.Field", "bool")}.WithCoreType(core.BOOLEAN))
)

var knownTypes = map[string]FieldType{}

func register(ft FieldType) FieldType {
	knownTypes[ft.Name] = ft
	return ft
}

func FieldTypeNamed(name string) FieldType {
	for k, v := range knownTypes {
		if k == name {
			return v
		}
	}
	return UNKNOWN
}
