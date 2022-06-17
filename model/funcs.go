package model

type Model struct {
	Name       string
	Attributes []*Attribute
}

func New(name string) *Model {
	return &Model{
		Name:       name,
		Attributes: []*Attribute{}}
}

// if modetype is given then create the attribute if missing
func (m *Model) Attr(name string, modeltype ...AttributeType) *Attribute {
	if len(modeltype) > 0 {
		att := &Attribute{Name: name, Type: modeltype[0]}
		m.Attributes = append(m.Attributes, att)
		return att
	}
	// find it by name
	for _, each := range m.Attributes {
		if each.Name == name {
			return each
		}
	}
	// not found
	return nil
}

type Attribute struct {
	Name string
	Type AttributeType
}

type Association struct {
	One   *Model
	Other *Model
}

func ToMany(m *Model) *Association {
	return &Association{
		Other: m,
	}
}

func (m *Model) Relation(name string, assoc *Association) *Association {
	assoc.One = m
	return assoc
}
