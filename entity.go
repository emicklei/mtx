package mtx

import (
	"fmt"
	"io"
)

const (
	EntityClass          = "mtx.Entity"
	EntityAttributeClass = "mtx.Attribute"
	EntityName           = "Entity.Name"    // property key to override Name of an Entity
	AttributeName        = "Attribute.Name" // property key to override Name of an Entity Attribute

)

type Entity struct {
	*Named
	pkg        *Package
	Attributes []*Attribute `json:"attributes"`
}

func NewEntity(name string) *Entity {
	return &Entity{Named: N(EntityClass, name)}
}

func (e *Entity) A(name string, typ Datatype, doc string) *Attribute {
	return e.Attribute(name).Type(typ).Doc(doc)
}

func (e *Entity) Doc(doc string) *Entity {
	e.Documentation = doc
	return e
}

func (e *Entity) Attribute(name string) *Attribute {
	attr, ok := FindByName(e.Attributes, name)
	if ok {
		return attr
	}
	attr = &Attribute{
		IsNullable: false, // required by default
	}
	attr.Named = N(EntityAttributeClass, name)
	e.Attributes = append(e.Attributes, attr)
	return attr
}

func (e *Entity) Validate(c *ErrorCollector) {
	e.Named.Validate(c)
	e.Named.CheckClass(c, "mtx.Entity")
	for _, each := range e.Attributes {
		each.Validate(c)
	}
}

// SourceOn writes Go source to recreate the receiver.
func (e *Entity) SourceOn(w io.Writer) {
}

//var _ TypedLabel = new(Attribute)

type Attribute struct {
	*Named
	Category      string   `json:"category,omitempty"`
	AttributeType Datatype `json:"type"`
	// IsNullable = true means the value can be NULL/nil
	IsNullable bool  `json:"is_nullable,omitempty"`
	Tags       []Tag `json:"tags,omitempty"`
}

func (a *Attribute) Type(t Datatype) *Attribute {
	a.AttributeType = t
	return a
}

func (a *Attribute) Doc(d string) *Attribute {
	a.Documentation = d
	return a
}

func (a *Attribute) Nullable() *Attribute {
	a.IsNullable = true
	return a
}

// SourceOn writes Go source to recreate the receiver.
func (a *Attribute) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "ent.Attribute(\"%s\")", a.Name)
}

func (a *Attribute) Validate(c *ErrorCollector) {
	a.Named.Validate(c)
	a.Named.CheckClass(c, "mtx.Attribute")
	a.AttributeType.Validate(c)
	a.AttributeType.Named.CheckClass(c, "mtx.Datatype")
}

type Tag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
