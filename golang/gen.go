package golang

import (
	"bytes"
	"fmt"

	"github.com/emicklei/mtx/core"
	"github.com/iancoleman/strcase"
)

type Option struct {
	Package string
}

func Package(pkg string) Option { return Option{Package: pkg} }

func Source(e *core.Entity, options ...Option) string {
	buf := new(bytes.Buffer)
	if d := e.Documentation; d != "" {
		fmt.Fprintf(buf, "// %s", e.Documentation)
	}
	fmt.Fprintf(buf, "type %s struct {\n", e.Name)
	for _, each := range e.Attributes {
		fmt.Fprintf(buf, "\t%s %s `json:\"%s,omitempty\" `// %s\n", goFieldName(each), goTypeSource(each), each.Name, each.Documentation)
	}
	fmt.Fprintf(buf, "}")
	return buf.String()
}

// TODO handle nullable
func goTypeSource(a *core.Attribute) string {
	if gt, ok := a.Get(core.GoTypeName); ok {
		return gt.(string)
	}
	typ, ok := attributeTypeToGoTypeMapping[a.AttributeType.Name]
	if ok {
		return typ
	}
	// fallback
	return "any"
}

func goFieldName(a *core.Attribute) string {
	// TODO check override
	return strcase.ToCamel(a.Name)
}
