package basic

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
)

type Attribute struct {
	*mtx.Named
	Category      string       `json:"category,omitempty"`
	AttributeType mtx.Datatype `json:"type"`
	// IsNullable = true means the value can be NULL/nil
	IsNullable bool `json:"is_nullable,omitempty"`
}

func NewAttribute(name string) *Attribute {
	attr := &Attribute{
		IsNullable: false, // required by default
	}
	attr.Named = mtx.N(EntityAttributeClass, name)
	return attr
}

func (a *Attribute) Type(t mtx.Datatype) *Attribute {
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

func (a *Attribute) Validate(c *mtx.ErrorCollector) {
	a.Named.Validate(c)
	a.Named.CheckClass(c, "basic.Attribute")
	a.AttributeType.Validate(c)
	a.AttributeType.Named.CheckClass(c, "mtx.Datatype")
}

func (a *Attribute) String() string {
	return fmt.Sprintf("%s:%s:%s", a.Name, a.AttributeType.Name, a.Class)
}
