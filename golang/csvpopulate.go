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

func (g *CSVPopulateMethodGenerator) Build(s *Struct) {
	buf := new(strings.Builder)
	buf.WriteString(fmt.Sprintf("func (r %s) CSVPopulate(record []string) (%s, error) {\n", s.Name, s.Name))
	for i, each := range s.Fields {
		fmt.Println("string -> ", each.FieldType.Name)

		buf.WriteString(fmt.Sprintf(`	if v := record[%d]; v != "" {
		r.%s = %s
	}`, i, each.Name, fromStringConvertFuncName(each.FieldType)))
		buf.WriteString("\n")
	}
	buf.WriteString("	return r, nil\n}\n")
	s.Methods = append(s.Methods, &Method{
		Named:  mtx.N("Test", "golang.Method"),
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

/**
func (r Identity) CSVPopulate(record []string) (Identity, error) {
	if v := record[0]; v != "" {
		r.LineItemId = bigquery.NullString{StringVal: v, Valid: true}
	}
	if v := record[1]; v != "" {
		r.TimeInterval = bigquery.NullString{StringVal: v, Valid: true}
	}
	return r, nil
}
**/
