package pg

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
)

func FromBasicType(bt mtx.Datatype) mtx.Datatype {
	pt := registry.MappedAttributeType(bt)
	if bt.IsNullable {
		pt.Set(golang.GoNullableTypeName, "pgtype.Text")
	}
	return pt
}
