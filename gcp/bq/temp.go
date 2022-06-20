package bq

import "github.com/emicklei/mtx/core"

type Dataset2 struct {
	core.Named
	Tables []*core.Table[core.Column[Datatype]]
}

func (d *Dataset2) Table(name string) *core.Table[core.Column[Datatype]] {
	t, ok := core.FindByName(d.Tables, name)
	if ok {
		return t
	}
	t = &core.Table[core.Column[Datatype]]{Named: core.N("bq.Table", name)}
	d.Tables = append(d.Tables, t)
	return t
}
