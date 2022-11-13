package pg

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func FromBasicType(bt mtx.Datatype) mtx.Datatype {
	pt := registry.MappedAttributeType(bt)
	if bt.IsNullable {
		return pt.Set(golang.GoNullableTypeName, "pgtype.Text")
	}
	return pt
}

// ToBasicType returns a mapped basic Datatype
func ToBasicType(dt mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(dt, registry.Class())

	return dt
}
