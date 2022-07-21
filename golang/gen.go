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
		fmt.Fprintf(buf, "\t%s %s ", goFieldName(each), GoTypeSource(each))
		// add tags
		if len(each.Tags) > 0 {
			fmt.Fprintf(buf, "`")
			for _, tag := range each.Tags {
				fmt.Fprintf(buf, "%s:\"%s\" ", tag.Name, tag.Value)
			}
			fmt.Fprintf(buf, "` ")
		}
		fmt.Fprintf(buf, "// %s\n", each.Documentation)
	}
	fmt.Fprintf(buf, "}")
	return buf.String()
}

// TODO handle nullable
func GoTypeSource(a *mtx.Attribute) string {
	if gt, ok := a.Get(mtx.GoTypeName); ok {
		// if typename is overridden then it should have taken care of nullable
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
