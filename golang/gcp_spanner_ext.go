package golang

import (
	"fmt"

	"github.com/emicklei/mtx/basic"
)

var spannerTagger = func(attr *basic.Attribute, field *Field) {
	field.Tags = append(field.Tags, Tag{
		Name:  "spanner",
		Value: fmt.Sprintf("%s", attr.Name),
	})
}

// WithSpannerTags is an Option that adds "spanner" tags to Go struct fields.
var WithSpannerTags = func(b *StructBuilder) *StructBuilder {
	b.fieldTaggers = append(b.fieldTaggers, spannerTagger)
	return b
}
