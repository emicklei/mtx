package db

import (
	"io"

	"github.com/emicklei/mtx"
)

type ExtendsDatabase interface {
	Table() ExtendsTable
}

type ExtendsTable interface {
	OwnerClass() string
	Column() ExtendsColumn
	SQLOn(table *Table, w io.Writer)
}

type ExtendsColumn interface {
	Datatype() ExtendsDatatype
	OwnerClass() string
	PostBuildAttribute(*Column, *mtx.Attribute)
}

type ExtendsDatatype interface {
	OwnerClass() string
}
