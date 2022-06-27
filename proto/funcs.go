package proto

import "github.com/emicklei/mtx/core"

type Package struct {
	Messages []*Message `json:"messages"`
}

func NewPackage(name string) *Package {
	return &Package{}
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

func (m *Message) Field(name string) *Field {
	f, ok := core.FindByName(m.Fields, name)
	if ok {
		return f
	}
	f = &Field{Named: core.N("proto.Field", name)}
	m.Fields = append(m.Fields, f)
	return f
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

type FieldType struct {
	*core.Named
}

type Field struct {
	*core.Named
	FieldType      FieldType `json:"type"`
	Repeated       bool
	Optional       bool
	SequenceNumber int `json:"nr"` // zero means unknown
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
