package proto

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

type Package struct {
	*mtx.Named
	Messages []*Message `json:"messages"`
}

func NewPackage(name string) *Package {
	return &Package{Named: mtx.N("proto.Package", name)}
}

func (p *Package) Validate(c *mtx.ErrorCollector) {
	p.Named.Validate(c)
	p.Named.CheckClass(c, "proto.Package")
	for _, each := range p.Messages {
		each.Validate(c)
	}
}

func (p *Package) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "pkg := proto.NewPackage(\"%s\")", p.Name)
	p.Named.SourceOn(w)
	for _, each := range p.Messages {
		fmt.Fprintln(w)
		each.SourceOn(w)
	}
}

type Message struct {
	*mtx.Named
	Fields []*Field `json:"fields"`
}

func (p *Package) Message(name string) *Message {
	m, ok := mtx.FindByName(p.Messages, name)
	if ok {
		return m
	}
	m = &Message{Named: mtx.N("proto.Message", name)}
	p.Messages = append(p.Messages, m)
	return m
}

func (p *Package) Doc(doc string) *Package {
	p.Named.Doc(doc)
	return p
}

func (m *Message) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "msg := pkg.Message(\"%s\")", m.Name)
	m.Named.SourceOn(w)
	for _, each := range m.Fields {
		fmt.Fprintln(w)
		each.SourceOn(w)
	}
}

func (m *Message) Validate(c *mtx.ErrorCollector) {
	m.Named.Validate(c)
	m.Named.CheckClass(c, "proto.Message")
	for _, each := range m.Fields {
		each.Validate(c)
	}
}

func ToEntity(m *Message) *basic.Entity {
	// temp
	return m.ToEntity()
}

func (m *Message) ToEntity() *basic.Entity {
	e := basic.NewEntity(m.Name)
	// share props
	e.Properties = m.Properties
	e.Doc(m.Documentation)
	for _, each := range m.Fields {
		_ = e.A(each.Name, *each.FieldType.BasicDatatype, each.Documentation)
		// if each.IsOptional {
		// 	attr.Nullable()
		// }
	}
	return e
}

func (m *Message) Field(name string) *Field {
	f, ok := mtx.FindByName(m.Fields, name)
	if ok {
		return f
	}
	f = &Field{Named: mtx.N("proto.Field", name)}
	m.Fields = append(m.Fields, f)
	return f
}

// F is a shortcut for Field.Number.Type.Doc
func (m *Message) F(name string, nr int, ft mtx.Datatype, doc string) *Field {
	return m.Field(name).Number(nr).Type(ft).Doc(doc)
}

// Compose adds a copy of a field without the sequence number
func (m *Message) Compose(f *Field) *Field {
	cp := *f
	cp.SequenceNumber = 0
	m.Fields = append(m.Fields, &cp)
	return &cp
}

// Doc overrides to return the receiver
func (m *Message) Doc(d string) *Message {
	m.Named.Documentation = d
	return m
}

var _ mtx.TypedLabel = new(Field)

type Field struct {
	*mtx.Named
	Category       string       `json:"category,omitempty"`
	FieldType      mtx.Datatype `json:"type"`
	IsRepeated     bool         `json:"repeated,omitempty"`
	IsOptional     bool         `json:"optional,omitempty"`
	SequenceNumber int          `json:"nr"` // zero means unknown
}

func (f *Field) GetDatatype() mtx.Datatype { return f.FieldType }

func (f *Field) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "msg.F(\"%s\",%d,", f.Name, f.SequenceNumber)
	f.FieldType.SourceOn(w)
	fmt.Fprintf(w, ",\"%s\")", f.Documentation)
}

func (f *Field) Validate(c *mtx.ErrorCollector) {
	f.Named.Validate(c)
	f.Named.CheckClass(c, "proto.Field")
	f.FieldType.CheckClass(c, "proto.Datatype")
}

func (f *Field) Type(ft mtx.Datatype) *Field {
	f.FieldType = ft
	return f
}

func (f *Field) Number(seq int) *Field {
	f.SequenceNumber = seq
	return f
}

func (f *Field) Optional() *Field {
	f.IsOptional = true
	return f
}

// Doc overrides to return the receiver
func (f *Field) Doc(d string) *Field {
	f.Named.Documentation = d
	return f
}

type DatatypeExtensions struct {
}

func (d DatatypeExtensions) OwnerClass() string { return "proto.Datatype" }
