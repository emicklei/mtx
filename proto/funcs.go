package proto

import (
	"fmt"
	"io"

	"github.com/emicklei/mtx/core"
)

type Package struct {
	*core.Named
	Messages []*Message `json:"messages"`
}

func NewPackage(name string) *Package {
	return &Package{Named: core.N("proto.Package", name)}
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
	*core.Named
	Fields []*Field `json:"fields"`
}

func (p *Package) Message(name string) *Message {
	m, ok := core.FindByName(p.Messages, name)
	if ok {
		return m
	}
	m = &Message{Named: core.N("proto.Message", name)}
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

func (m *Message) Field(name string) *Field {
	f, ok := core.FindByName(m.Fields, name)
	if ok {
		return f
	}
	f = &Field{Named: core.N("proto.Field", name)}
	m.Fields = append(m.Fields, f)
	return f
}

// F is a shortcut for Field.Number.Type.Doc
func (m *Message) F(name string, nr int, ft FieldType, doc string) *Field {
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

type Field struct {
	*core.Named
	FieldType      FieldType `json:"type"`
	Repeated       bool
	Optional       bool
	SequenceNumber int `json:"nr"` // zero means unknown
}

func (f *Field) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "msg.F(\"%s\",%d,", f.Name, f.SequenceNumber)
	f.FieldType.SourceOn(w)
	fmt.Fprintf(w, ",\"%s\")", f.Documentation)
}

func (f *Field) Type(ft FieldType) *Field {
	f.FieldType = ft
	return f
}

func (f *Field) Number(seq int) *Field {
	f.SequenceNumber = seq
	return f
}

// Doc overrides to return the receiver
func (f *Field) Doc(d string) *Field {
	f.Named.Documentation = d
	return f
}
