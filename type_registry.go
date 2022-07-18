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

type HasAttributeType interface {
	HasName
	AttrType() AttributeType
}

func NewTypeRegistry(class string) *TypeRegistry {
	return &TypeRegistry{
		class:        class,
		knownTypes:   map[string]Datatype{},
		encodedTypes: map[string]Datatype{},
	}
}

// MappedAttributeType returns the best matching type.
func (r *TypeRegistry) MappedAttributeType(at AttributeType) Datatype {
	for _, each := range r.knownTypes {
		if each.AttrType().Equals(at) {
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

func (r *TypeRegistry) EncodeAs(at AttributeType, dt Datatype) {
	// check existing
	_, ok := r.encodedTypes[at.Name]
	if ok {
		panic("duplicate encoded key:" + at.Name)
	}
	r.encodedTypes[at.Name] = dt
}

func (r *TypeRegistry) Add(d Datatype) Datatype {
	// check existing
	_, ok := r.knownTypes[d.GetName()]
	if ok {
		panic("duplicate known key:" + d.GetName())
	}
	r.knownTypes[d.GetName()] = d
	return d
}

func (r *TypeRegistry) TypeNamed(name string) (Datatype, bool) {
	e, ok := r.knownTypes[name]
	return e, ok
}

func (r *TypeRegistry) Register(typename string, at AttributeType, isUserDefined bool) Datatype {
	dt := Datatype{
		Named:         N(r.class, typename),
		IsUserDefined: isUserDefined,
		AttributeType: at,
	}
	return r.Add(dt)
}
