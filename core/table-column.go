package core

type Table[C any] struct {
	Named
	Columns       []*C
	columnFactory func(name string) C
}

type Column[D any] struct {
	Named
	IsPrimary bool
	Type      D
}

func (c *Column[D]) Doc(d string) *Column[D] {
	c.Named.Doc = d
	return c
}

func (c *Column[D]) Datatype(dt D) *Column[D] {
	c.Type = dt
	return c
}

func (t *Table[C]) Column(name string) *C {
	c, ok := FindByName(t.Columns, name)
	if ok {
		return c
	}
	tc := t.columnFactory(name)
	t.Columns = append(t.Columns, &tc)
	return &tc
}
