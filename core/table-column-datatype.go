package core

type DatabaseTable interface {
	OwnerClass() string
}

type TableColumn interface {
	OwnerClass() string
}

type ColumnDatatype interface {
	OwnerClass() string
}

type Table[T any, C TableColumn, D ColumnDatatype] struct {
	*Named
	Columns    []*Column[C, D]
	Extensions T
}

func (t *Table[T, C, D]) Doc(c string) *Table[T, C, D] {
	t.Named.Doc = c
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

type Column[C TableColumn, D ColumnDatatype] struct {
	*Named
	Type       Datatype[D]
	Extensions C
}

func (c *Column[C, D]) Doc(d string) *Column[C, D] {
	c.Named.Doc = d
	return c
}

func (c *Column[C, D]) Datatype(dt Datatype[D]) *Column[C, D] {
	c.Type = dt
	return c
}

type Datatype[D ColumnDatatype] struct {
	*Named
	Extensions D
	Scale      int `json:"Scale,omitempty"`     // Maximum scale range: 0 ≤ S ≤ 9
	Precision  int `json:"Precision,omitempty"` // Maximum precision range: max(1, S) ≤ P ≤ S + 29
}
