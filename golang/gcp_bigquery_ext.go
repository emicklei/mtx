package golang

import (
	"fmt"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var WithBigQueryTypeMapper = func(b *StructBuilder) *StructBuilder {
	return b.WithTypeMapper(bigQueryTypeMapper)
}

// bigQueryTypeMapper maps Attribute types to Go types from the Google bigquery Go package
var bigQueryTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	// TODO Date -> basic.Register("civil.Date")

	if at.Name == basic.Decimal.Name {
		// for both nullable and not
		return Type("*big.Rat")
	}
	if !nullable {
		if at.Name == basic.Integer.Name {
			// check bits
			if at.GetInt("bits", 0) == 64 {
				return Int64
			}
			return StandardTypeMapper(at, nullable)
		}
		if at.Name == basic.JSON.Name {
			return String
		}
		return StandardTypeMapper(at, nullable)
	}
	// nullable
	// https://pkg.go.dev/cloud.google.com/go/bigquery#pkg-types
	switch at.Name {
	case basic.String.Name, basic.JSON.Name:
		return Type("bigquery.NullString")
	case basic.Boolean.Name:
		return Type("bigquery.NullBool")
	case basic.Timestamp.Name:
		return Type("bigquery.NullTime")
	case basic.Date.Name:
		return Type("bigquery.NullDate")
	case basic.DateTime.Name:
		return Type("bigquery.NullDateTime")
	case basic.Timestamp.Name:
		return Type("bigquery.NullTimestamp")
	case basic.Integer.Name:
		return Type("bigquery.NullInt64")
	case basic.Float.Name, basic.Double.Name:
		return Type("bigquery.NullFloat64")
	case basic.Bytes.Name:
		return Bytes // empty bytes are considered null
	default:
		return StandardTypeMapper(at, nullable)
	}
}

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
