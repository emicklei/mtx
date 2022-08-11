package golang

import (
	"fmt"
	"strings"

	"github.com/emicklei/mtx"
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
		buf.WriteString(fmt.Sprintf(`	if v := record[%d]; v != "" {
		r.%s = bigquery.NullString{StringVal: v, Valid: true}
	}`, i, each.Name))
		buf.WriteString("\n")
	}
	buf.WriteString("	return r, nil\n}\n")
	s.Methods = append(s.Methods, &Method{
		Named:  mtx.N("Test", "golang.Method"),
		Source: buf.String(),
	})
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
