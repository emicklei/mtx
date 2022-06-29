package core

const (
	EntityClass          = "Entity"
	EntityAttributeClass = "Attribute"
	EntityName           = "Entity.Name" // property key to override Name of an Entity
	GoTypeName           = "GoTypeName"  // property key to bypass mapping Name of a Attribute Type
)

type Entity struct {
	*Named
	Attributes []*Attribute `json:"attributes"`
}

func NewEntity(name string) *Entity {
	return &Entity{Named: N(EntityClass, name)}
}

// if modetype is given then create the attribute if missing
func (m *Entity) Attribute(name string) *Attribute {
	attr, ok := FindByName(m.Attributes, name)
	if ok {
		return attr
	}
	attr = &Attribute{
		IsRequired: true, // required by default
	}
	attr.Named = N(EntityAttributeClass, name)
	m.Attributes = append(m.Attributes, attr)
	return attr
}

type Attribute struct {
	*Named
	AttributeType AttributeType `json:"type"`
	IsRequired    bool          `json:"required"`
}

func (a *Attribute) Type(t AttributeType) *Attribute {
	a.AttributeType = t
	return a
}

func (a *Attribute) Doc(d string) *Attribute {
	a.Documentation = d
	return a
}

func (a *Attribute) Optional() *Attribute {
	a.IsRequired = false
	return a
}
