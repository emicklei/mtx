package mtx

type TypeRegistry struct {
	class        string
	knownTypes   map[string]Datatype
	encodedTypes map[string]Datatype
	trace        bool
}

func NewTypeRegistry(class string) *TypeRegistry {
	return &TypeRegistry{
		class:        class,
		knownTypes:   map[string]Datatype{},
		encodedTypes: map[string]Datatype{},
	}
}

func (r *TypeRegistry) Trace()        { r.trace = true }
func (r *TypeRegistry) Class() string { return r.class }

// MappedAttributeType returns the best matching known or encodede type.
// Properties set on the argument type are copied into the new result
func (r *TypeRegistry) MappedAttributeType(attrType Datatype) Datatype {
	if !attrType.HasName() {
		if r.trace {
			trace("registry", r.class, "attrType", attrType, "msg", "no name")
		}
		return r.knownTypes["any"]
	}
	// check encoded types
	et, ok := r.encodedTypes[attrType.Name]
	if ok {
		if r.trace {
			trace("registry", r.class, "attrType", attrType, "encodedType", et)
		}
		return et.WithCopiedPropertiesFrom(attrType)
	}
	// check known types
	for _, each := range r.knownTypes {
		if dt := each.BasicDatatype; dt != nil && dt.Name == attrType.Name {
			return each.WithCopiedPropertiesFrom(attrType)
		}
	}
	// fallback
	a, ok := r.knownTypes["any"] // must have an any
	if !ok {
		panic("warning: missing any in " + r.class)
	}
	if r.trace {
		trace("registry", r.class, "attrType", attrType, "fallback to", a)
	}
	return a.WithCopiedPropertiesFrom(attrType)
}

func (r *TypeRegistry) EncodeAs(attrType Datatype, encodedType Datatype) {
	r.encodedTypes[attrType.Name] = encodedType
}

func (r *TypeRegistry) Add(d Datatype) Datatype {
	if d.Class != r.class {
		panic("wrong class")
	}
	r.knownTypes[d.GetName()] = d
	return d
}

func (r *TypeRegistry) Register(typename string, isUserDefined bool) Datatype {
	dt := Datatype{
		Named:         N(r.class, typename),
		IsUserDefined: isUserDefined,
	}
	return r.Add(dt)
}

func (r *TypeRegistry) Type(typename string) Datatype {
	dt, ok := r.knownTypes[typename]
	if ok {
		return dt
	}
	return r.RegisterType(typename, Unknown)
}

func (r *TypeRegistry) RegisterType(typename string, attrType Datatype) Datatype {
	dt := Datatype{
		Named:         N(r.class, typename),
		BasicDatatype: &attrType,
		IsUserDefined: true,
	}
	return r.Add(dt)
}

func (r *TypeRegistry) Standard(typename string, attrType Datatype) Datatype {
	dt := Datatype{
		Named:         N(r.class, typename),
		BasicDatatype: &attrType,
	}
	return r.Add(dt)
}
