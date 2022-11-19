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
		if typeName, ok := gt.Get(mtx.GoNullableTypeName); ok {
			return Type(typeName.(string)) // no longer nullable
		}
	} else {
		// If GoName set then return that type
		if n, ok := gt.Get(mtx.GoTypeName); ok {
			return Type(n.(string))
		}
	}
	switch gt.Name {
	case "string":
		if gt.IsNullable {
			return Type("*string")
		}
		return String
	case basic.Integer.Name:
		if v, ok := gt.Get("bits"); ok {
			bits := v.(int)
			switch bits {
			case 64:
				return Type("int64")
			case 32:
				return Type("int32")
			}
		}
	case basic.Timestamp.Name:
		return Time
	case basic.Bytes.Name:
		return Bytes
	case basic.JSON.Name:
		return String
	case basic.Boolean.Name:
		return Bool
	}
	return mtx.Unknown
}
