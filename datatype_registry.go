package mtx

const (
	// TODO why is this needed?
	StandardType    = false
	UserDefinedType = true
)

type TypeRegistry struct {
	class        string
	knownTypes   map[string]Datatype
	encodedTypes map[string]Datatype
}

func NewTypeRegistry(class string) *TypeRegistry {
	return &TypeRegistry{
		class:        class,
		knownTypes:   map[string]Datatype{},
		encodedTypes: map[string]Datatype{},
	}
}

func (r *TypeRegistry) Class() string { return r.class }

// MappedAttributeType returns the best matching known or encodede type.
func (r *TypeRegistry) MappedAttributeType(attrType Datatype) Datatype {
	for _, each := range r.knownTypes {
		if dt := each.AttributeDatatype; dt != nil && dt.Name == attrType.Name {
			return each
		}
	}
	// check encoded types
	et, ok := r.encodedTypes[attrType.Name]
	if ok {
		return et
	}
	return r.knownTypes["any"]
}

func (r *TypeRegistry) EncodeAs(at Datatype, dt Datatype) {
	// check existing
	_, ok := r.encodedTypes[at.Name]
	if ok {
		return
	}
	r.encodedTypes[at.Name] = dt
}

func (r *TypeRegistry) Add(d Datatype) Datatype {
	if d.Class != r.class {
		panic("wrong class")
	}
	r.knownTypes[d.GetName()] = d
	return d
}

func (r *TypeRegistry) TypeNamed(name string) (Datatype, bool) {
	e, ok := r.knownTypes[name]
	return e, ok
}

func (r *TypeRegistry) Register(typename string, isUserDefined bool) Datatype {
	dt := Datatype{
		Named:         N(r.class, typename),
		IsUserDefined: isUserDefined,
	}
	return r.Add(dt)
}

func (r *TypeRegistry) Type(typename string) Datatype {
	dt, ok := r.TypeNamed(typename)
	if ok {
		return dt
	}
	return r.RegisterType(typename, UNKNOWN)
}

func (r *TypeRegistry) RegisterType(typename string, attrType Datatype) Datatype {
	dt := Datatype{
		Named:             N(r.class, typename),
		AttributeDatatype: &attrType,
		IsUserDefined:     true,
	}
	return r.Add(dt)
}

func (r *TypeRegistry) Standard(typename string, attrType Datatype) Datatype {
	dt := Datatype{
		Named:             N(r.class, typename),
		AttributeDatatype: &attrType,
	}
	return r.Add(dt)
}
