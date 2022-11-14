package spanner

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
	"github.com/emicklei/mtx/golang"
)

func ToBasicType(st mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(st, registry.Class())
	if st.IsNullable {
		switch st.Name {
		case String.Name:
			basic.String.Set(golang.GoNullableTypeName, "spanner.NullString")
		}
		bt := *st.BasicDatatype
		return bt.Nullable()
	}
	return *st.BasicDatatype
}
