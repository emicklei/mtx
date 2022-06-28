package core

type AttributeType struct {
	Typename string `json:"name"`
}

var (
	BOOLEAN   = AttributeType{Typename: "boolean"}
	STRING    = AttributeType{Typename: "string"}
	INTEGER   = AttributeType{Typename: "integer"}
	ID        = AttributeType{Typename: "identifier"}
	DATE      = AttributeType{Typename: "date"}      // yyyy mm dd
	DATETIME  = AttributeType{Typename: "datetime"}  // yyyy mm dd hh mm ss
	TIMESTAMP = AttributeType{Typename: "timestamp"} // yyyy mm dd hh mm ss + zone
	BYTES     = AttributeType{Typename: "bytes"}
)
