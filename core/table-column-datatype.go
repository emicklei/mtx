package core

type ExtendsTable interface {
	OwnerClass() string
}

type ExtendsColumn interface {
	OwnerClass() string
}

type ExtendsDatatype interface {
	OwnerClass() string
}

type Table[T ExtendsTable, C ExtendsColumn, D ExtendsDatatype] struct {
	*Named
	Columns    []*Column[C, D]
	Extensions T
}

func (t *Table[T, C, D]) Doc(d string) *Table[T, C, D] {
	t.Documentation = d
	return t
}

func (t *Table[T, C, D]) Column(name string) *Column[C, D] {
	c, ok := FindByName(t.Columns, name)
	if ok {
		return c
	}
	c = new(Column[C, D])
	c.Named = N(c.Extensions.OwnerClass(), name)
	t.Columns = append(t.Columns, c)
	return c
}

type Column[C ExtendsColumn, D ExtendsDatatype] struct {
	*Named
	Type       Datatype[D]
	Extensions C
}

func (c *Column[C, D]) Doc(d string) *Column[C, D] {
	c.Documentation = d
	return c
}

func (c *Column[C, D]) Datatype(dt Datatype[D]) *Column[C, D] {
	c.Type = dt
	return c
}

type Datatype[D ExtendsDatatype] struct {
	*Named
	Extensions D
}
