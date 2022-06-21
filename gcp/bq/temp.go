package bq

import "github.com/emicklei/mtx/core"

type Dataset2 struct {
	core.Named
	Tables []*Table2
}

type Table2 struct {
	core.Table
}

func (t *Table2) Column(name string) *Column2 {
	return nil
}

type Column2 struct {
	core.Column
}
