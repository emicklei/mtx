package mtx

type TypeRegistry[T HasAttributeType] struct {
	knownTypes map[string]T
}

type HasAttributeType interface {
	HasName
	AttrType() AttributeType
}

func NewTypeRegistry[T HasAttributeType]() *TypeRegistry[T] {
	return &TypeRegistry[T]{knownTypes: map[string]T{}}
}

// MappedAttributeType returns the best matching attribute type.
func (r *TypeRegistry[T]) MappedAttributeType(at AttributeType) T {
	for _, each := range r.knownTypes {
		if each.AttrType().Equals(at) {
			return each
		}
	}
	// TODO specials
	return r.knownTypes["any"]
}

func (r *TypeRegistry[T]) Add(dt T) T {
	r.knownTypes[dt.GetName()] = dt
	return dt
}

func (r *TypeRegistry[T]) TypeNamed(name string) (T, bool) {
	e, ok := r.knownTypes[name]
	return e, ok
}
