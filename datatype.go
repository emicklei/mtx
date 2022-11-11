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

var Unknown = NewBasicType("Unknown")

type Datatype struct {
	*Named
	BasicDatatype *Datatype `json:"attribute_type,omitempty"`
	IsNullable    bool      `json:"is_nullable"`
	IsUserDefined bool      `json:"is_userdefined"`
	// for arrays
	ElementType *Datatype       `json:"element_type,omitempty"`
	Extensions  ExtendsDatatype `json:"ext,omitempty"`
}

func NewBasicType(name string) Datatype {
	return Datatype{
		Named: N("basic.Datatype", name),
	}
}

func (d Datatype) HasName() bool { return d.Named != nil && d.Name != "" }

func (d Datatype) BasicType() Datatype {
	return *d.BasicDatatype
}

// TODO needed?
func (d Datatype) EncodedFrom(at Datatype) Datatype {
	d.BasicDatatype = &at
	return d
}

func (d Datatype) WithBasicDatatype(dt Datatype) Datatype {
	d.BasicDatatype = &dt
	return d
}

func (d Datatype) Nullable() Datatype {
	d.IsNullable = true
	return d
}

// Equal matches only name and class
func (d Datatype) Equal(o Datatype) bool {
	return d.Name == o.Name && d.Class == o.Class
}

func (d Datatype) String() string {
	required := " "
	if d.IsNullable {
		required = " ? "
	}
	if d.Named == nil {
		return fmt.Sprintf("*unnamed*%sDatatype", required)
	}
	if d.Properties != nil {
		doc, _ := json.Marshal(d.Properties)
		return fmt.Sprintf("%s%s(%s) %s", d.Name, required, d.Class, string(doc))
	}
	return fmt.Sprintf("%s%s(%s)", d.Name, required, d.Class)
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

func Array(elementType Datatype) Datatype {
	return Datatype{
		Named:       N("mtx.Datatype", "array"),
		ElementType: &elementType,
	}
}
