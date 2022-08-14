package golang

import (
	"fmt"

	"github.com/emicklei/mtx"
)

var WithSpannerTypeMapper = func(b *StructBuilder) *StructBuilder {
	return b.WithTypeMapper(bigQueryTypeMapper)
}

// spannerTypeMapper maps Attribute types to Go types from the Google spanner Go package
var spannerTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	if !nullable {
		return StandardTypeMapper(at, nullable)
	}
	// nullable
	switch at.Name {
	case mtx.String.Name:
		return Type("spanner.NullString")
	case mtx.JSON.NullableAttributeDatatype.Name:
		return Type("spanner.NullJSON")
	default:
		return StandardTypeMapper(at, nullable)
	}
}

var spannerTagger = func(attr *mtx.Attribute, field *Field) {
	field.Tags = append(field.Tags, Tag{
		Name:  "spanner",
		Value: fmt.Sprintf("%s,omitempty", attr.Name),
	})
}

// WithSpannerTags is an Option that adds "spanner" tags to Go struct fields.
var WithSpannerTags = func(b *StructBuilder) *StructBuilder {
	b.fieldTaggers = append(b.fieldTaggers, spannerTagger)
	return b
}
