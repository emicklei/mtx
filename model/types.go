package model

type AttributeType struct {
	Typename string
}

var String = AttributeType{Typename: "string"}
var Integer = AttributeType{Typename: "integer"}
var Identifier = AttributeType{Typename: "uuid"}
