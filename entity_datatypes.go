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
	Unknown   = registry.Add(NewAttributeType("any"))
	Boolean   = registry.Add(NewAttributeType("boolean"))
	String    = registry.Add(NewAttributeType("string"))
	Integer   = registry.Add(NewAttributeType("integer"))
	ID        = registry.Add(NewAttributeType("identifier"))
	Date      = registry.Add(NewAttributeType("date"))      // yyyy mm dd
	DateTime  = registry.Add(NewAttributeType("datetime"))  // yyyy mm dd hh mm ss
	Timestamp = registry.Add(NewAttributeType("timestamp")) // yyyy mm dd hh mm ss + zone
	Bytes     = registry.Add(NewAttributeType("bytes"))
	Float     = registry.Add(NewAttributeType("float"))
	Double    = registry.Add(NewAttributeType("double"))
	Decimal   = registry.Add(NewAttributeType("decimal"))
	JSON      = registry.Add(NewAttributeType("JSON"))
	Duration  = registry.Add(NewAttributeType("duration")) // y,m,d,h,m,s
	UUID      = registry.Add(NewAttributeType("uuid"))
)

func Array(elementType Datatype) Datatype {
	dt := Datatype{
		Named:       N("mtx.Datatype", "array"),
		ElementType: &elementType,
	}
	return registry.Add(dt)
}
