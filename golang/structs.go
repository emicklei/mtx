package golang

import (
	"bytes"
	"fmt"
	"io"

	"github.com/emicklei/mtx"
	"github.com/iancoleman/strcase"
)

const (
	GoName             = "GoName"     // property key to bypass mapping Name
	GoTypeName         = "GoTypeName" // property key to bypass mapping Name of a Attribute Type
	GoNullableTypeName = "GoNullableTypeName"
)

type Package struct {
	*mtx.Named
	Structs []*Struct
}

func NewPackage(name string) *Package {
	return &Package{
		Named: mtx.N("golang.Package", name),
	}
}

// T is short for Type.Doc
func (p *Package) T(name string, doc string) *Struct {
	return p.Type(name).Doc(doc)
}

// Type finds or create a Struct type
func (p *Package) Type(name string) *Struct {
	if t, ok := mtx.FindByName(p.Structs, name); ok {
		return t
	}
	t := &Struct{
		Named:   mtx.N("golang.Struct", name),
		Package: p,
	}
	return t
}

type Struct struct {
	*mtx.Named
	Package *Package
	Fields  []*Field
}

func (s *Struct) Doc(doc string) *Struct {
	s.Named.Documentation = doc
	return s
}

// F is short for Field.Type.Doc
func (s *Struct) F(name string, dt mtx.Datatype, doc string) *Field {
	return s.Field(name).Type(dt).Doc(doc)
}

// Field finds or creates a Field.
func (s *Struct) Field(name string) *Field {
	if f, ok := mtx.FindByName(s.Fields, name); ok {
		return f
	}
	f := &Field{
		Named:     mtx.N("golang.Field", name),
		FieldType: mtx.Unknown,
	}
	s.Fields = append(s.Fields, f)
	return f
}

// Go returns the source in Go for defining this struct type.
func (s *Struct) Go() string {
	var buf bytes.Buffer
	s.GoOn(&buf)
	return buf.String()
}

func (s *Struct) GoOn(w io.Writer) {
	fmt.Fprintf(w, "// %s : %s\n", s.Name, s.Documentation)
	fmt.Fprintf(w, "type %s struct {\n", s.Name)
	for _, each := range s.Fields {
		fmt.Fprintf(w, "\t%s %s ", each.Name, each.FieldType.Name)
		// add tags
		if len(each.Tags) > 0 {
			fmt.Fprintf(w, "`")
			for _, tag := range each.Tags {
				fmt.Fprintf(w, "%s:\"%s\" ", tag.Name, tag.Value)
			}
			fmt.Fprintf(w, "` ")
		}
		fmt.Fprintf(w, "// %s\n", each.Documentation)
	}
	fmt.Fprint(w, "}\n")
}

type Field struct {
	*mtx.Named
	Category  string `json:"category,omitempty"`
	FieldType mtx.Datatype
	Tags      []mtx.Tag
}

func (f *Field) Doc(doc string) *Field {
	f.Named.Documentation = doc
	return f
}

func (f *Field) Type(dt mtx.Datatype) *Field {
	f.FieldType = dt
	return f
}

func ToStruct(ent *mtx.Entity) *Struct {
	n := strcase.ToCamel(ent.Name)
	if v, ok := ent.Get(GoTypeName); ok {
		n = v.(string)
	}
	str := &Struct{
		Named: mtx.N("golang.Struct", n),
	}
	str.Documentation = ent.Documentation
	for _, each := range ent.Attributes {
		f := &Field{
			Named:     mtx.N("golang.Field", goFieldName(each)),
			FieldType: goDatatype(each),
			Tags:      each.Tags,
		}
		f.Documentation = each.Documentation
		str.Fields = append(str.Fields, f)
	}
	return str
}

func goFieldName(a *mtx.Attribute) string {
	// TODO check override
	return strcase.ToCamel(a.Name)
}

func goDatatype(a *mtx.Attribute) mtx.Datatype {
	if n, ok := a.Get(GoTypeName); ok {
		return mtx.NewAttributeType(n.(string))
	}
	if a.IsNullable {
		if null := a.AttributeType.NullableAttributeDatatype; null != nil {
			return *null
		}
		mapped := registry.MappedAttributeType(a.AttributeType)
		if null := mapped.NullableAttributeDatatype; null != nil {
			return *null
		}
		return mtx.Datatype{
			Named: mtx.N("golang.Datatype", "*"+mapped.Name),
		}
	}
	return registry.MappedAttributeType(a.AttributeType)
}
