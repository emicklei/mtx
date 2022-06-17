package proto

type Package struct{}

func NewPackage(name string) *Package { return new(Package) }

type Message struct{}

func (p *Package) NewMessage(name string) *Message { return nil }

func (m *Message) Field(name string, fieldtype FieldType) any { return nil }

type FieldType int
