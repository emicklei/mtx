package mtx

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type HasName interface{ GetName() string }

func FindByName[T HasName](elements []*T, name string) (*T, bool) {
	for _, each := range elements {
		if (*each).GetName() == name {
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
	Documentation string         `json:"doc,omitempty"`
}

func (n *Named) Namespace() string {
	return n.Class[0:strings.LastIndex(n.Class, ".")]
}

func (n *Named) Validate(c *ErrorCollector) {
	if n.Name == "" {
		c.Add(n, errors.New("empty name"))
	}
}

func (n *Named) CheckClass(c *ErrorCollector, must string) {
	if got, want := n.Class, must; got != want {
		c.Add(n, fmt.Errorf("got %s want %s", n.Class, must))
	}
}

func (n *Named) GetName() string {
	return n.Name
}

func (n *Named) SourceOn(w io.Writer) {
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

func (n *Named) GetInt(key string, absent int) int {
	if n.Properties == nil {
		return absent
	}
	v, ok := n.Properties[key]
	if !ok {
		return absent
	}
	i, ok := v.(int)
	if !ok {
		return absent
	}
	return i
}

func (n *Named) CopyPropertiesFrom(n2 *Named) {
	if n2.Properties == nil {
		return
	}
	if n.Properties == nil {
		n.Properties = n2.Properties
		return
	}
	for k, v := range n2.Properties {
		// detect override?
		n.Set(k, v)
	}
}

func (n *Named) Doc(d string) *Named {
	n.Documentation = d
	return n
}

func N(class, name string) *Named { return &Named{Name: name, Class: class} }
