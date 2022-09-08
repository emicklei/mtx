package mtx

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type ExtendsDatatype interface {
	OwnerClass() string
}

type Datatype struct {
	*Named
	AttributeDatatype         *Datatype       `json:"attribute_type,omitempty"`
	NullableAttributeDatatype *Datatype       `json:"nullable_attribute_type,omitempty"`
	IsUserDefined             bool            `json:"is_user_defined,omitempty"`
	ElementType               *Datatype       `json:"element_type,omitempty"`
	Extensions                ExtendsDatatype `json:"ext,omitempty"`
}

func NewAttributeType(name string) Datatype {
	return Datatype{
		Named: N("mtx.Datatype", name),
	}
}

func (d Datatype) HasName() bool { return d.Named != nil && d.Name != "" }

func (d Datatype) EncodedFrom(at Datatype) Datatype {
	d.AttributeDatatype = &at
	return d
}

func (d Datatype) WithAttributeDatatype(dt Datatype) Datatype {
	d.AttributeDatatype = &dt
	return d
}

func (d Datatype) WithNullable(dt Datatype) Datatype {
	d.NullableAttributeDatatype = &dt
	return d
}

// Equal matches only name and class
func (d Datatype) Equal(o Datatype) bool {
	return d.Name == o.Name && d.Class == o.Class
}

func (d Datatype) String() string {
	if d.Named == nil {
		return "*unnamed* Datatype"
	}
	if d.Properties != nil {
		doc, _ := json.Marshal(d.Properties)
		return fmt.Sprintf("%s (%s) %s", d.Name, d.Class, string(doc))
	}
	return fmt.Sprintf("%s (%s) {}", d.Name, d.Class)
}

func (d Datatype) SourceOn(w io.Writer) {
	pkg := d.Class[0:strings.Index(d.Class, ".")]
	if d.IsUserDefined {
		fmt.Fprintf(w, "%s.Type(\"%s\")", pkg, d.Name)
		return
	}
	fmt.Fprintf(w, "%s.%s", pkg, strings.Title(d.Name))
}

// Set overrides Named.Set to preserve return type
func (d Datatype) Set(key string, value any) Datatype {
	d.Named.Set(key, value)
	return d
}

func (d Datatype) Validate(c *ErrorCollector) {
	d.Named.Validate(c)
}

// Return a Datatype with all properties copied from the argument.
// If no properties are in the argument then return a copy of the receiver.
func (d Datatype) WithCopiedPropertiesFrom(o Datatype) Datatype {
	if o.Properties == nil {
		return d
	}
	if len(o.Properties) == 0 {
		return d
	}
	d.CopyPropertiesFrom(o.Named)
	return d
}
