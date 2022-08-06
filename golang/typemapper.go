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
	// is nullable
	if nt := dt.NullableAttributeDatatype; nt != nil {
		return *nt
	}
	return Type("*" + dt.Name)
}

var PgxTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	// TODO
	return at
}