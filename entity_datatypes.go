package mtx

type AttributeType struct {
	Name string `json:"name"`
	// array or dictionary
	ElementType *AttributeType `json:"element_type,omitempty"`
}

var (
	UNKNOWN   = AttributeType{Name: "any"}
	BOOLEAN   = AttributeType{Name: "boolean"}
	STRING    = AttributeType{Name: "string"}
	INTEGER   = AttributeType{Name: "integer"}
	ID        = AttributeType{Name: "identifier"}
	DATE      = AttributeType{Name: "date"}      // yyyy mm dd
	DATETIME  = AttributeType{Name: "datetime"}  // yyyy mm dd hh mm ss
	TIMESTAMP = AttributeType{Name: "timestamp"} // yyyy mm dd hh mm ss + zone
	BYTES     = AttributeType{Name: "bytes"}
	FLOAT     = AttributeType{Name: "float"}
	DOUBLE    = AttributeType{Name: "double"}
	DECIMAL   = AttributeType{Name: "decimal"}
	JSON      = AttributeType{Name: "JSON"}
)

func Array(elementType AttributeType) AttributeType {
	return AttributeType{Name: "array", ElementType: &elementType}
}

func (a AttributeType) Equals(o AttributeType) bool {
	if a.Name != o.Name {
		return false
	}
	// TODO
	return true
}

func RegisterType(name string) AttributeType {
	// register(
	return AttributeType{Name: name}
}
