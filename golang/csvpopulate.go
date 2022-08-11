package golang

import (
	"fmt"
	"strings"

	"github.com/emicklei/mtx"
	"github.com/iancoleman/strcase"
)

var WithCSVPopulate = func(b *StructBuilder) *StructBuilder {
	p := new(CSVPopulateMethodGenerator)
	b.methodBuilders = append(b.methodBuilders, p.Build)
	return b
}

type CSVPopulateMethodGenerator struct {
}

/**
if v, err := r.Identity.CSVPopulate(record); err == nil {
	r.Identity = v
} else {
	return r, err
}
**/

func (g *CSVPopulateMethodGenerator) Build(s *Struct) {
	buf := new(strings.Builder)
	buf.WriteString(fmt.Sprintf("func (r %s) CSVPopulate(record []string,offset int) (%s, error) {\n", s.Name, s.Name))
	offset := 0
	for _, each := range s.Fields {
		if _, ok := each.Get("IsEntity"); ok { // TODO
			buf.WriteString(fmt.Sprintf(`	if v, err := r.%s.CSVPopulate(record,%d); err == nil {
		r.%s = v
	} else {
		return r, err
	}`, each.Name, offset, each.Name))
			offset += each.GetInt("AttributeCount", -1)
		} else {
			buf.WriteString(fmt.Sprintf(`	if v := record[offset+%d]; v != "" {
		r.%s = %s
	}`, offset, each.Name, fromStringConvertFuncName(each.FieldType)))
			offset++
		}
		buf.WriteString("\n")
	}
	buf.WriteString("	return r, nil\n}\n")
	s.Methods = append(s.Methods, &Method{
		Named:  mtx.N("CSVPopulate", "golang.Method"),
		Source: buf.String(),
	})
}

func fromStringConvertFuncName(dt mtx.Datatype) string {
	if dt.Name == mtx.String.Name {
		return "v"
	}
	if strings.HasPrefix(dt.Name, "*") {
		return fmt.Sprintf("StringToPtr%s(v)", strcase.ToCamel(dt.Name[1:]))
	}
	return fmt.Sprintf("StringTo%s(v)", strcase.ToCamel(dt.Name))
}
