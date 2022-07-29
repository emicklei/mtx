package mtx

var (
	registry            = NewTypeRegistry("mtx.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

func Register(typename string) Datatype {
	return registry.Register(typename, true)
}

var (
	UNKNOWN   = registry.Add(NewAttributeType("any"))
	BOOLEAN   = registry.Add(NewAttributeType("boolean"))
	STRING    = registry.Add(NewAttributeType("string"))
	INTEGER   = registry.Add(NewAttributeType("integer"))
	ID        = registry.Add(NewAttributeType("identifier"))
	DATE      = registry.Add(NewAttributeType("date"))      // yyyy mm dd
	DATETIME  = registry.Add(NewAttributeType("datetime"))  // yyyy mm dd hh mm ss
	TIMESTAMP = registry.Add(NewAttributeType("timestamp")) // yyyy mm dd hh mm ss + zone
	BYTES     = registry.Add(NewAttributeType("bytes"))
	FLOAT     = registry.Add(NewAttributeType("float"))
	DOUBLE    = registry.Add(NewAttributeType("double"))
	DECIMAL   = registry.Add(NewAttributeType("decimal"))
	JSON      = registry.Add(NewAttributeType("json"))
	DURATION  = registry.Add(NewAttributeType("duration")) // y,m,d,h,m,s
	UUID      = registry.Add(NewAttributeType("uuid"))
)

func Array(elementType Datatype) Datatype {
	dt := Datatype{
		Named:       N("mtx.Datatype", "array"),
		ElementType: &elementType,
	}
	return registry.Add(dt)
}
