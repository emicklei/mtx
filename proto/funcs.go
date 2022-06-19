package proto

import "github.com/emicklei/mtx/core"

type Package struct {
	messages []*Message
}

func NewPackage(name string) *Package {
	return &Package{}
}

type Message struct {
	core.Named
	Fields map[string]*Field
}

func (p *Package) Message(name string) *Message {
	m, ok := core.FindByName(p.messages, name)
	if ok {
		return m
	}
	m = &Message{Fields: map[string]*Field{}}
	p.messages = append(p.messages, m)
	return m
}

func (m *Message) Field(name string, fieldtype FieldType) any { return nil }

func (m *Message) Name() string { return m.Name() }

type FieldType int

type Field struct {
	Type     FieldType
	Repeated bool
	Optional bool
}
