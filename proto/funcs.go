package proto

import "github.com/emicklei/mtx/core"

type Package struct {
	messages []*Message
}

func NewPackage(name string) *Package {
	return &Package{}
}

type Message struct {
	*core.Named
	Fields []*Field
}

func (p *Package) Message(name string) *Message {
	m, ok := core.FindByName(p.messages, name)
	if ok {
		return m
	}
	m = &Message{Named: core.N("proto.Message", name)}
	p.messages = append(p.messages, m)
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

type FieldType struct {
	*core.Named
}

type Field struct {
	*core.Named
	Type     FieldType
	Repeated bool
	Optional bool
}

func (f *Field) FieldType(ft FieldType) *Field {
	f.Type = ft
	return f
}
