package golang

import (
	"fmt"

	"github.com/emicklei/mtx/basic"
)

var bigQueryTagger = func(attr *basic.Attribute, field *Field) {
	field.Tags = append(field.Tags, Tag{
		Name:  "bigquery",
		Value: fmt.Sprintf("%s", attr.Name),
	})
}

// WithBigQueryTags is an Option that adds "bigquery" tags to Go struct fields.
var WithBigQueryTags = func(b *StructBuilder) *StructBuilder {
	b.fieldTaggers = append(b.fieldTaggers, bigQueryTagger)
	return b
}
