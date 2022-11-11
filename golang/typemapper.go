package golang

import (
	"github.com/emicklei/mtx"
)

type TypeMapper func(at mtx.Datatype, nullable bool) mtx.Datatype

var StandardTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	dt := registry.MappedAttributeType(at)
	if !nullable {
		return dt
	}
	return Type("*" + dt.Name)
}

var PgxTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	// TODO
	return at
}

func Datatype(edt mtx.Datatype) mtx.Datatype {
	if edt.IsNullable {
		if typeName, ok := edt.Get(GoNullableTypeName); ok {
			return Type(typeName.(string)) // no longer nullable
		}
	}
	if edt.Name == "string" {
		if edt.IsNullable {
			return Type("*string")
		}
		return String
	}
	return mtx.Unknown
}
