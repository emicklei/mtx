package db

import "github.com/emicklei/mtx/core"

type Table struct {
	core.Named
	Columns    []*Column
	Properties map[string]string
}

func (t *Table) Doc(c string) *Table {
	t.Named.Doc = c
	return t
}

func (t *Table) Set(key, value string) *Table {
	if t.Properties == nil {
		t.Properties = map[string]string{}
	}
	t.Properties[key] = value
	return t
}

func (t *Table) Column(name string) *Column {
	c, ok := core.FindByName(t.Columns, name)
	if ok {
		return c
	}
	c = &Column{Named: core.N("db.Column", name)}
	t.Columns = append(t.Columns, c)
	return c
}

type Column struct {
	core.Named
	Type       Datatype
	IsPrimary  bool
	Properties map[string]string
}

func (c *Column) Doc(d string) *Column {
	c.Named.Doc = d
	return c
}

func (c *Column) BePrimary() *Column {
	c.IsPrimary = true
	return c
}

func (c *Column) Set(key, value string) *Column {
	if c.Properties == nil {
		c.Properties = map[string]string{}
	}
	c.Properties[key] = value
	return c
}

func (c *Column) Datatype(dt Datatype) *Column {
	c.Type = dt
	return c
}

type Datatype struct {
	core.Named
}
