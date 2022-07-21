package mtx

var registry = NewTypeRegistry("Datatype")

func register(typename string) Datatype {
	return registry.Register(typename, false)
}

func Register(typename string) Datatype {
	return registry.Register(typename, true)
}

var (
	UNKNOWN   = register("any")
	BOOLEAN   = register("boolean")
	STRING    = register("string")
	INTEGER   = register("integer")
	ID        = register("identifier")
	DATE      = register("date")      // yyyy mm dd
	DATETIME  = register("datetime")  // yyyy mm dd hh mm ss
	TIMESTAMP = register("timestamp") // yyyy mm dd hh mm ss + zone
	BYTES     = register("bytes")
	FLOAT     = register("float")
	DOUBLE    = register("double")
	DECIMAL   = register("decimal")
	JSON      = register("json")
	DURATION  = register("duration") // y,m,d,h,m,s
	UUID      = register("uuid")
)

func Array(elementType Datatype) Datatype {
	dt := Datatype{
		Named:       N("Datatype", "array"),
		ElementType: &elementType,
	}
	return registry.Add(dt)
}

func (a Datatype) Equals(o Datatype) bool {
	if a.Name != o.Name {
		return false
	}
	// TODO
	return true
}
