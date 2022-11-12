package golang

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

// ToStruct builds a Struct. Option control build strategies.
func ToStruct(ent *basic.Entity, options ...Option) *Struct {
	b := NewStructBuilder(ent)
	for _, each := range options {
		b = each(b)
	}
	return b.Build()
}

func FromBasicType(gt mtx.Datatype) mtx.Datatype {
	if gt.IsNullable {
		if typeName, ok := gt.Get(GoNullableTypeName); ok {
			return Type(typeName.(string)) // no longer nullable
		}
	}
	if gt.Name == "string" {
		if gt.IsNullable {
			return Type("*string")
		}
		return String
	}
	return mtx.Unknown
}
