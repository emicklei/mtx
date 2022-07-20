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
	AttributeDatatype *Datatype
	// This means : can the datatype be used to capture a NULL value
	CanRepresentNull bool `json:"can_present_null,omitempty"`
	IsUserDefined    bool `json:"is_user_defined,omitempty"`
	ElementType      *Datatype
	Extensions       ExtendsDatatype `json:"ext"`
}

// func (d Datatype) EncodedFrom(at AttributeType) Datatype {
// 	d.AttributeType = at
// 	return d
// }

func (d Datatype) Optional() Datatype {
	d.CanRepresentNull = true
	return d
}

func (d Datatype) WithAttributeDatatype(dt Datatype) Datatype {
	d.AttributeDatatype = &dt
	return d
}

func (d Datatype) String() string {
	return fmt.Sprintf("%s (%s) : %T", d.Name, d.Class, d)
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
