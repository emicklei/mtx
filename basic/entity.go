package basic

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

const (
	EntityClass          = "basic.Entity"
	EntityAttributeClass = "basic.Attribute"
	EntityName           = "Entity.Name"    // property key to override Name of an Entity
	AttributeName        = "Attribute.Name" // property key to override Name of an Entity Attribute

)

type Entity struct {
	*mtx.Named
	pkg        *Package
	Attributes []*Attribute `json:"attributes"`
}

func NewEntity(name string) *Entity {
	return &Entity{Named: mtx.N(EntityClass, name)}
}

func (e *Entity) A(name string, typ mtx.Datatype, doc string) *Attribute {
	return e.Attribute(name).Type(typ).Doc(doc)
}

func (e *Entity) Doc(doc string) *Entity {
	e.Documentation = doc
	return e
}

func (e *Entity) Attribute(name string) *Attribute {
	attr, ok := mtx.FindByName(e.Attributes, name)
	if ok {
		return attr
	}
	attr = NewAttribute(name)
	e.Attributes = append(e.Attributes, attr)
	return attr
}

func (e *Entity) Validate(c *mtx.ErrorCollector) {
	e.Named.Validate(c)
	e.Named.CheckClass(c, "basic.Entity")
	for _, each := range e.Attributes {
		each.Validate(c)
	}
}

func (e *Entity) String() string {
	return fmt.Sprintf("%s:%s", e.Name, e.Class)
}

// SourceOn writes Go source to recreate the receiver.
func (e *Entity) SourceOn(w io.Writer) {
}

//var _ TypedLabel = new(Attribute)
