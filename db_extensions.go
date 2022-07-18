package mtx

import "io"

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
}

type ExtendsDatatype interface {
	OwnerClass() string
}

type SQLWriter interface{ SQLOn(w io.Writer) }
