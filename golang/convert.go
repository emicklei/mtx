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
	mtx.CheckClass(gt, basic.Boolean.Class)

	if gt.IsNullable {
		if typeName, ok := gt.Get(GoNullableTypeName); ok {
			return Type(typeName.(string)) // no longer nullable
		}
	} else {
		// If GoName set then return that type
		if n, ok := gt.Get(GoName); ok {
			return Type(n.(string))
		}
	}

	if gt.Name == "string" {
		if gt.IsNullable {
			return Type("*string")
		}
		return String
	}
	if gt.Name == "decimal" {

	}
	return mtx.Unknown
}
