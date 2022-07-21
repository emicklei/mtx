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

// MappedAttributeType returns the best matching type.
func (r *TypeRegistry) MappedAttributeType(at Datatype) Datatype {
	for _, each := range r.knownTypes {
		if dt := each.AttributeDatatype; dt != nil && dt.Name == at.Name {
			return each
		}
	}
	// check encoded types
	et, ok := r.encodedTypes[at.Name]
	if ok {
		return et
	}
	return r.knownTypes["any"] // TODO return the unknown
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
	// check existing
	dt, ok := r.knownTypes[d.GetName()]
	if ok {
		return dt
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
