package golang

import (
	"fmt"

	"github.com/emicklei/mtx"
)

var WithBigQueryTypeMapper = func(b *StructBuilder) *StructBuilder {
	return b.WithTypeMapper(bigQueryTypeMapper)
}

// bigQueryTypeMapper maps Attribute types to Go types from the Google bigquery Go package
var bigQueryTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	if at.Name == mtx.Decimal.Name {
		// for both nullable and not
		return Type("*big.Rat")
	}
	if !nullable {
		return StandardTypeMapper(at, nullable)
	}
	// nullable
	// https://pkg.go.dev/cloud.google.com/go/bigquery#pkg-types
	switch at.Name {
	case mtx.String.Name:
		return Type("bigquery.NullString")
	case mtx.Boolean.Name:
		return Type("bigquery.NullBool")
	case mtx.Timestamp.Name:
		return Type("bigquery.NullTime")
	case mtx.Date.Name:
		return Type("bigquery.NullDate")
	case mtx.DateTime.Name:
		return Type("bigquery.NullDateTime")
	case mtx.Timestamp.Name:
		return Type("bigquery.NullTimestamp")
	case mtx.Integer.Name:
		return Type("bigquery.NullInt64")
	case mtx.Float.Name, mtx.Double.Name:
		return Type("bigquery.NullFloat64")
	default:
		return StandardTypeMapper(at, nullable)
	}
}

var bigQueryTagger = func(attr *mtx.Attribute, field *Field) {
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
