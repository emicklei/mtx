package spanner

import (
	"github.com/emicklei/mtx"
)

func ToBasicType(st mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(st, registry.Class())

	bt := *st.BasicDatatype
	bt.Named = bt.Named.WithPropertiesCopiedFrom(st.Named)
	if st.IsNullable {
		return bt.WithNullable()
	}
	return bt
}
