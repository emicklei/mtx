package core

type HasName interface{ HasName(name string) bool }

func FindByName[T HasName](elements []T, name string) (T, bool) {
	for _, each := range elements {
		if each.HasName(name) {
			return each, true
		}
	}
	var t T
	return t, false
}

type Named struct {
	Name  string
	Class string
	Doc   string `json:"Doc,omitempty"`
}

func (n Named) HasName(v string) bool {
	return n.Name == v
}

func N(class, name string) Named { return Named{Name: name, Class: class} }

type Namespace struct {
	Name     string
	elements map[string]Named
}

func NewNamespace(name string) *Namespace {
	return &Namespace{
		Name:     name,
		elements: map[string]Named{},
	}
}
