package core

import (
	"fmt"
	"io"
)

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

func (n Named) SourceOn(w io.Writer) {
	if d := n.Documentation; d != "" {
		fmt.Fprintf(w, ".Doc(\"%s\")", d)
	}
	for k, v := range n.Properties {
		var vs string
		if s, ok := v.(string); ok {
			vs = s
		} else {
			vs = fmt.Sprintf("%v", v)
		}
		fmt.Fprintf(w, ".Set(\"%s\",%s)", k, vs)
	}
}

// Set add/overwrites a property that can used to pass context information.
func (n *Named) Set(key string, value any) *Named {
	if n.Properties == nil {
		n.Properties = map[string]any{key: value}
	}
	n.Properties[key] = value
	return n
}

func (n *Named) Get(key string) (any, bool) {
	if n.Properties == nil {
		return "", false
	}
	v, ok := n.Properties[key]
	return v, ok
}

func (n *Named) Doc(d string) *Named {
	n.Documentation = d
	return n
}

func N(class, name string) *Named { return &Named{Name: name, Class: class} }
