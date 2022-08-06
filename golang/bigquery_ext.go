package golang

import (
	"fmt"

	"github.com/emicklei/mtx"
)

var WithBigQueryTypeMapper = func(b *StructBuilder) *StructBuilder {
	return b.WithTypeMapper(BigQueryTypeMapper)
}

// BigQueryTypeMapper maps Attribute types to Go types from the Google bigquery Go package
var BigQueryTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	if !nullable {
		return StandardTypeMapper(at, nullable)
	}
	// nullable
	switch at.Name {
	case mtx.String.Name:
		return Type("bigquery.NullString")
	default:
		return StandardTypeMapper(at, nullable)
	}
}

var BigQueryTagger = func(attr *mtx.Attribute, field *Field) {
	field.Tags = append(field.Tags, Tag{
		Name:  "bigquery",
		Value: fmt.Sprintf("%s,omitempty", attr.Name),
	})
}

// WithBigQueryTagger is an Option that adds "bigquery" tags to Go struct fields.
var WithBigQueryTagger = func(b *StructBuilder) *StructBuilder {
	b.fieldTaggers = append(b.fieldTaggers, BigQueryTagger)
	return b
}
