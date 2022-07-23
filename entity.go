package mtx

import (
	"fmt"
	"io"
)

const (
	EntityClass          = "Entity"
	EntityAttributeClass = "Attribute"
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

// SourceOn writes Go source to recreate the receiver.
func (e *Entity) SourceOn(w io.Writer) {
}

//var _ TypedLabel = new(Attribute)

type Attribute struct {
	*Named
	AttributeType Datatype `json:"type"`
	// IsNullable = true means the value can be NULL/nil
	IsNullable bool  `json:"is_nullable,omitempty"`
	Tags       []Tag `json:"tags,omitempty"`
}

// func (a *Attribute) GetDatatype() Datatype {
// 	return Datatype{Named: a.Named, AttributeType: a.AttributeType}
// }

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

type Tag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}
