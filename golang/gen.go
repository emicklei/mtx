package golang

import (
	"bytes"
	"fmt"

	"github.com/emicklei/mtx"
	"github.com/iancoleman/strcase"
)

type Option struct {
	Package string
}

func Package(pkg string) Option { return Option{Package: pkg} }

func Source(e *mtx.Entity, options ...Option) string {
	buf := new(bytes.Buffer)
	if d := e.Documentation; d != "" {
		fmt.Fprintf(buf, "// %s\n", e.Documentation)
	}
	fmt.Fprintf(buf, "type %s struct {\n", e.Name)
	for _, each := range e.Attributes {
		fmt.Fprintf(buf, "\t%s %s `json:\"%s,omitempty\" `// %s\n", goFieldName(each), goTypeSource(each), each.Name, each.Documentation)
	}
	fmt.Fprintf(buf, "}")
	return buf.String()
}

// TODO handle nullable
func goTypeSource(a *mtx.Attribute) string {
	if gt, ok := a.Get(mtx.GoTypeName); ok {
		if a.IsNullable {
			// TODO too simple
			return "*" + gt.(string)
		}
		return gt.(string)
	}
	if a.AttributeType.Named == nil {
		panic("missing Named in attribute.attributetype")
	}
	typ, ok := attributeTypeToGoTypeMapping[a.AttributeType.Name]
	if ok {
		if a.IsNullable {
			// TODO too simple
			return "*" + typ
		}
		return typ
	}
	// fallback
	return "any"
}

func goFieldName(a *mtx.Attribute) string {
	// TODO check override
	return strcase.ToCamel(a.Name)
}
