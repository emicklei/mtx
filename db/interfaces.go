package db

import (
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

type ExtendsDatabase interface {
	TableExtensions() ExtendsTable
}

type ExtendsTable interface {
	OwnerClass() string
	ColumnExtensions() ExtendsColumn
	SQLOn(table *Table, w io.Writer)
}

type ExtendsColumn interface {
	Datatype() ExtendsDatatype
	OwnerClass() string
	PostBuildAttribute(*Column, *basic.Attribute)
	ToBasicType(dt mtx.Datatype) mtx.Datatype
}

type ExtendsDatatype interface {
	OwnerClass() string
}
