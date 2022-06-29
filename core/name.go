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
	Name          string         `json:"name"`
	Class         string         `json:"class"`
	Properties    map[string]any `json:"properties,omitempty"`
	Documentation string         `json:"documentation,omitempty"`
}

func (n Named) HasName(v string) bool {
	return n.Name == v
}

// Set add/overwrites a property that can used to pass context information.
func (n *Named) Set(key string, value any) {
	if n.Properties == nil {
		n.Properties = map[string]any{key: value}
	}
	n.Properties[key] = value
}

func (n *Named) Get(key string) (any, bool) {
	if n.Properties == nil {
		return "", false
	}
	v, ok := n.Properties[key]
	return v, ok
}

func (n *Named) Doc(d string) {
	n.Documentation = d
}

func N(class, name string) *Named { return &Named{Name: name, Class: class} }
