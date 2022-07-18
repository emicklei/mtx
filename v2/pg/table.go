package pg

import (
	"github.com/emicklei/mtx"
	v2 "github.com/emicklei/mtx/v2"
)

type Table struct {
	*v2.Table
}

func NewTable(name string) *Table {
	return &Table{
		Table: &v2.Table{
			Named: mtx.N("pg.Table", name),
		},
	}
}

type Column struct {
	*v2.Column
}

type Datatype struct {
	*v2.Datatype
}
