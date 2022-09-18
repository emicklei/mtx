package mtx

type FieldMapping struct {
	From        *Attribute
	NameMapping func(m *FieldMapping) string
	Type        *TypeMapping
}

func NewFieldMapping(attr *Attribute) *FieldMapping {
	return &FieldMapping{
		From:        attr,
		NameMapping: func(m *FieldMapping) string { return m.From.Name },
	}
}

func (m *FieldMapping) RenameTo(newName string) *FieldMapping {
	m.NameMapping = func(m *FieldMapping) string { return newName }
	return m
}

func (m *FieldMapping) ToAttribute() *Attribute {
	a := NewAttribute(m.NameMapping(m))
	if m.Type != nil {
		a.AttributeType = m.Type.To
	}
	return a
}

type EntityMapping struct {
	From        *Entity
	NameMapping func(m *EntityMapping) string
	Fields      []*FieldMapping
}

func NewEntityMapping(from *Entity) *EntityMapping {
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

func (m *EntityMapping) SetAttributeType(fromAttrName string, typ Datatype) *EntityMapping {
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

func (m *EntityMapping) ToEntity() *Entity {
	e := NewEntity(m.NameMapping(m))
	for _, each := range m.Fields { // my mappings
		e.Attributes = append(e.Attributes, each.ToAttribute())
	}
	return e
}

type TypeMapping struct {
	From Datatype
	To   Datatype
}

func NewTypeMapping(from Datatype) *TypeMapping {
	return &TypeMapping{
		From: from,
		To:   from,
	}
}
