package bq

import "github.com/emicklei/mtx/core"

type Dataset struct {
	core.Named
	Tables []*Table
}

func (d *Dataset) Table(name string) *Table {
	t, ok := core.FindByName(d.Tables, name)
	if ok {
		return t
	}
	t = &Table{Named: core.N("bq.Table", name)}
	d.Tables = append(d.Tables, t)
	return t
}

type Table struct {
	core.Named
	Columns []*Column
}

func (t *Table) Column(name string) *Column {
	c, ok := core.FindByName(t.Columns, name)
	if ok {
		return c
	}
	c = &Column{Named: core.N("bq.Column", name)}
	t.Columns = append(t.Columns, c)
	return c
}

type Column struct {
	core.Named
	Type Datatype
}

func (c *Column) Datatype(dt Datatype) *Column {
	c.Type = dt
	return c
}
