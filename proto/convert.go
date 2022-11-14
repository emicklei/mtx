package proto

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

func ToEntity(m *Message) *basic.Entity {
	// temp
	return m.ToEntity()
}

func ToBasicType(pt mtx.Datatype) mtx.Datatype {
	mtx.CheckClass(pt, registry.Class())

	bt := *pt.BasicDatatype
	bt.CopyPropertiesFrom(pt.Named)
	return bt
}
