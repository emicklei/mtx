package golang

import (
	"github.com/emicklei/mtx"
)

type TypeMapper func(at mtx.Datatype, nullable bool) mtx.Datatype

var StandardTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	if nullable {
		return Type("*" + at.Name)
	}
	return at
}

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

var PgxTypeMapper = func(at mtx.Datatype, nullable bool) mtx.Datatype {
	// TODO
	return at
}
