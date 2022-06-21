package core

type HasName interface{ HasName(name string) bool }

func FindByName[T HasName](elements []*T, name string) (*T, bool) {
	for _, each := range elements {
		if (*each).HasName(name) {
			return each, true
		}
	}
	var t T
	return &t, false
}

type Named struct {
	Name       string
	Class      string
	Properties map[string]string `json:"Properties,omitempty"`
	Doc        string            `json:"Doc,omitempty"`
}

func (n Named) HasName(v string) bool {
	return n.Name == v
}

func (n *Named) Set(key, value string) {
	if n.Properties == nil {
		n.Properties = map[string]string{key: value}
	}
	n.Properties[key] = value
}

func N(class, name string) *Named { return &Named{Name: name, Class: class} }
