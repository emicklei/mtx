package spanner

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
	"github.com/emicklei/mtx/golang"
)

func ToBasicType(dt mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(dt, registry.Class())
	if dt.Equal(String) {
		if dt.IsNullable {
			return basic.String.Set(golang.GoNullableTypeName, "spanner.NullString").Nullable()
		}
	}
	return mtx.Unknown
}
