package pg

import (
	"github.com/emicklei/mtx"
)

type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType) DType {
	dt := DType{
		Named: mtx.N("pg.Datatype", typename),
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
	return register(name, mtx.UNKNOWN)
}

var (
	UNKNOWN = register("ANY", mtx.UNKNOWN)
	STRING  = register("string", mtx.STRING)
)
