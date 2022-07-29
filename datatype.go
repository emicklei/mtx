package mtx

import (
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
	return fmt.Sprintf("%s (%s)", d.Name, d.Class)
}

func (d Datatype) SourceOn(w io.Writer) {
	pkg := d.Class[0:strings.Index(d.Class, ".")]
	if d.IsUserDefined {
		fmt.Fprintf(w, "%s.Type(\"%s\")", pkg, d.Name)
		return
	}
	fmt.Fprintf(w, "%s.%s", pkg, strings.ToUpper(d.Name))
}

// Set overrides Named.Set to preserve return type
func (d Datatype) Set(key string, value any) Datatype {
	d.Named.Set(key, value)
	return d
}

func (d Datatype) Validate(c *ErrorCollector) {
	d.Named.Validate(c)
}
