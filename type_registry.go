package mtx

const (
	// TODO why is this needed?
	StandardType    = false
	UserDefinedType = true
)

type TypeRegistry[T HasAttributeType] struct {
	knownTypes   map[string]T
	encodedTypes map[string]T
}

type HasAttributeType interface {
	HasName
	AttrType() AttributeType
}

func NewTypeRegistry[T HasAttributeType]() *TypeRegistry[T] {
	return &TypeRegistry[T]{
		knownTypes:   map[string]T{},
		encodedTypes: map[string]T{},
	}
}

// MappedAttributeType returns the best matching type.
func (r *TypeRegistry[T]) MappedAttributeType(at AttributeType) T {
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

func (r *TypeRegistry[T]) EncodeAs(at AttributeType, dt T) {
	// check existing
	_, ok := r.encodedTypes[at.Name]
	if ok {
		panic("duplicate encoded key:" + at.Name)
	}
	r.encodedTypes[at.Name] = dt
}

func (r *TypeRegistry[T]) Add(dt T) T {
	// check existing
	_, ok := r.knownTypes[dt.GetName()]
	if ok {
		panic("duplicate known key:" + dt.GetName())
	}
	r.knownTypes[dt.GetName()] = dt
	return dt
}

func (r *TypeRegistry[T]) TypeNamed(name string) (T, bool) {
	e, ok := r.knownTypes[name]
	return e, ok
}
