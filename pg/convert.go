package pg

import (
	"github.com/emicklei/mtx"
)

func FromBasicType(bt mtx.Datatype) mtx.Datatype {
	pt := registry.MappedAttributeType(bt)
	if bt.IsNullable {
		return pt.Set(mtx.GoNullableTypeName, "pgtype.Text")
	}
	return pt
}

// ToBasicType returns a mapped basic Datatype
func ToBasicType(dt mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(dt, registry.Class())

	bt := *dt.BasicDatatype
	bt = bt.WithCopiedPropertiesFrom(dt)
	if dt.IsNullable {
		return bt.WithNullable()
	}
	return bt
}
