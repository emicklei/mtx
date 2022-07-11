package mtx

import (
	"fmt"
	"io"
)

const (
	EntityClass          = "Entity"
	EntityAttributeClass = "Attribute"
	EntityName           = "Entity.Name" // property key to override Name of an Entity
	GoTypeName           = "GoTypeName"  // property key to bypass mapping Name of a Attribute Type
)

type Entity struct {
	*Named
	pkg        *Package
	Attributes []*Attribute `json:"attributes"`
}

func NewEntity(name string) *Entity {
	return &Entity{Named: N(EntityClass, name)}
}

func (e *Entity) A(name string, typ AttributeType, doc string) *Attribute {
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
		IsRequired: true, // required by default
	}
	attr.Named = N(EntityAttributeClass, name)
	e.Attributes = append(e.Attributes, attr)
	return attr
}

// SourceOn writes Go source to recreate the receiver.
func (e *Entity) SourceOn(w io.Writer) {
}

type Attribute struct {
	*Named
	AttributeType AttributeType `json:"type"`
	IsRequired    bool          `json:"required"`
}

func (a *Attribute) Type(t AttributeType) *Attribute {
	a.AttributeType = t
	return a
}

func (a *Attribute) Doc(d string) *Attribute {
	a.Documentation = d
	return a
}

func (a *Attribute) Optional() *Attribute {
	a.IsRequired = false
	return a
}

// SourceOn writes Go source to recreate the receiver.
func (a *Attribute) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "ent.Attribute(\"%s\")", a.Name)
}
