package proto

import (
	"fmt"
	"io"
	"strings"

	"github.com/emicklei/mtx"
)

var (
	UNKNOWN = register(FieldType{Named: mtx.N("proto.FieldType", "any")}.WithCoreType(mtx.UNKNOWN))
	DOUBLE  = register(FieldType{Named: mtx.N("proto.FieldType", "double")}.WithCoreType(mtx.DOUBLE))
	FLOAT   = register(FieldType{Named: mtx.N("proto.FieldType", "float")}.WithCoreType(mtx.FLOAT))
	STRING  = register(FieldType{Named: mtx.N("proto.FieldType", "string")}.WithCoreType(mtx.STRING))
	INT32   = register(FieldType{Named: mtx.N("proto.FieldType", "int32")}.WithCoreType(mtx.INTEGER))
	INT64   = register(FieldType{Named: mtx.N("proto.FieldType", "int64")}.WithCoreType(mtx.INTEGER)) //.Set("bits", 64))
	BOOL    = register(FieldType{Named: mtx.N("proto.FieldType", "bool")}.WithCoreType(mtx.BOOLEAN))
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
	return RegisterType(name)
}

type FieldType struct {
	*mtx.Named
	IsCustom      bool              `json:"is_custom"`
	AttributeType mtx.AttributeType `json:"-"`
}

func (ft FieldType) SourceOn(w io.Writer) {
	if ft.IsCustom {
		fmt.Fprintf(w, "proto.RegisterType(\"%s\")", ft.Name)
		return
	}
	fmt.Fprintf(w, "proto.%s", strings.ToUpper(ft.Name))
}

func (ft FieldType) WithCoreType(at mtx.AttributeType) FieldType {
	ft.AttributeType = at
	return ft
}

func RegisterType(name string) FieldType {
	// exists?
	if ft, ok := knownTypes[name]; ok {
		return ft
	}
	// new!
	ft := FieldType{Named: mtx.N("proto.FieldType", name), IsCustom: true}.WithCoreType(mtx.RegisterType(name))
	return register(ft)
}
