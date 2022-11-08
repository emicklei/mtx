package golang

import (
	"fmt"

	"github.com/emicklei/mtx/basic"
)

type FieldTagger func(attr *basic.Attribute, field *Field)

var JSONTagger = func(attr *basic.Attribute, field *Field) {
	field.Tags = append(field.Tags, Tag{
		Name:  "json",
		Value: fmt.Sprintf("%s,omitempty", attr.Name),
	})
}

type Tag struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// WithJSONTags is an Option that adds "json" tags to Go struct fields.
var WithJSONTags = func(b *StructBuilder) *StructBuilder {
	b.fieldTaggers = append(b.fieldTaggers, JSONTagger)
	return b
}
