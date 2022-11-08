package m2m

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

type FieldMapping struct {
	From        *basic.Attribute
	NameMapping func(m *FieldMapping) string
	Type        *TypeMapping
}

func NewFieldMapping(attr *basic.Attribute) *FieldMapping {
	return &FieldMapping{
		From:        attr,
		NameMapping: func(m *FieldMapping) string { return m.From.Name },
	}
}

func (m *FieldMapping) RenameTo(newName string) *FieldMapping {
	m.NameMapping = func(m *FieldMapping) string { return newName }
	return m
}

func (m *FieldMapping) ToAttribute() *basic.Attribute {
	a := basic.NewAttribute(m.NameMapping(m))
	if m.Type != nil {
		a.AttributeType = m.Type.To
	}
	return a
}

type EntityMapping struct {
	From        *basic.Entity
	NameMapping func(m *EntityMapping) string
	Fields      []*FieldMapping
}

func NewEntityMapping(from *basic.Entity) *EntityMapping {
	fm := []*FieldMapping{}
	for _, each := range from.Attributes {
		fm = append(fm, NewFieldMapping(each))
	}
	return &EntityMapping{
		From:        from,
		NameMapping: func(m *EntityMapping) string { return m.From.Name },
		Fields:      fm,
	}
}

func (m *EntityMapping) SetName(newName string) *EntityMapping {
	m.NameMapping = func(m *EntityMapping) string { return newName }
	return m
}

func (m *EntityMapping) SetAttributeType(fromAttrName string, typ mtx.Datatype) *EntityMapping {
	var fm *FieldMapping
	for _, each := range m.Fields {
		if each.From.Name == fromAttrName {
			fm = each
			break
		}
	}
	if fm == nil {
		// ignore
		return m
	}
	fm.Type = NewTypeMapping(fm.From.AttributeType)
	fm.Type.To = typ
	return m
}

func (m *EntityMapping) SetAttributeName(fromAttrName, newName string) *EntityMapping {
	var fm *FieldMapping
	for _, each := range m.Fields {
		if each.From.Name == fromAttrName {
			fm = each
			break
		}
	}
	if fm == nil {
		// ignore
		return m
	}
	fm.RenameTo(newName)
	return m
}

func (m *EntityMapping) ToEntity() *basic.Entity {
	e := basic.NewEntity(m.NameMapping(m))
	for _, each := range m.Fields { // my mappings
		e.Attributes = append(e.Attributes, each.ToAttribute())
	}
	return e
}

type TypeMapping struct {
	From mtx.Datatype
	To   mtx.Datatype
}

func NewTypeMapping(from mtx.Datatype) *TypeMapping {
	return &TypeMapping{
		From: from,
		To:   from,
	}
}
