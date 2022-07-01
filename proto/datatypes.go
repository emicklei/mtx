package proto

import (
	"fmt"
	"io"
	"strings"

	"github.com/emicklei/mtx/core"
)

var (
	UNKNOWN = register(FieldType{Named: core.N("proto.FieldType", "any")}.WithCoreType(core.UNKNOWN))
	DOUBLE  = register(FieldType{Named: core.N("proto.FieldType", "double")}.WithCoreType(core.DOUBLE))
	FLOAT   = register(FieldType{Named: core.N("proto.FieldType", "float")}.WithCoreType(core.FLOAT))
	STRING  = register(FieldType{Named: core.N("proto.FieldType", "string")}.WithCoreType(core.STRING))
	INT32   = register(FieldType{Named: core.N("proto.FieldType", "int32")}.WithCoreType(core.INTEGER))
	INT64   = register(FieldType{Named: core.N("proto.FieldType", "int64")}.WithCoreType(core.INTEGER)) //.Set("bits", 64))
	BOOL    = register(FieldType{Named: core.N("proto.FieldType", "bool")}.WithCoreType(core.BOOLEAN))
)

var knownTypes = map[string]FieldType{}

func register(ft FieldType) FieldType {
	knownTypes[ft.Name] = ft
	return ft
}

func TypeNamed(name string) FieldType {
	for k, v := range knownTypes {
		if k == name {
			return v
		}
	}
	return UNKNOWN
}

type FieldType struct {
	*core.Named
	AttributeType core.AttributeType `json:"-"`
}

func (ft FieldType) SourceOn(w io.Writer) {
	fmt.Fprintf(w, "proto.%s", strings.ToUpper(ft.Name))
}

func (ft FieldType) WithCoreType(at core.AttributeType) FieldType {
	ft.AttributeType = at
	return ft
}

func RegisterType(name string) FieldType {
	ft := FieldType{Named: core.N("proto.FieldType", name)}.WithCoreType(core.RegisterType(name))
	return register(ft)
}
