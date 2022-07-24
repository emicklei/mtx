package bq

import (
	"fmt"

	"github.com/emicklei/mtx"
)

var (
	registry            = mtx.NewTypeRegistry("bq.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	BYTES  = registry.Standard("BYTES", mtx.BYTES)
	STRING = registry.Standard("STRING", mtx.STRING).WithNullable(mtx.Register("bigquery.NullString"))
)

func MaxBytes(max int64) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N("bq.Datatype", "BYTES"),
		Extensions: DatatypeExtensions{Max: max},
	}
}

var (
	// https://cloud.google.com/bigquery/docs/reference/standard-sql/json-data#sql
	JSON = registry.Standard("JSON", mtx.JSON)
)

// TODO look at civil package https://pkg.go.dev/cloud.google.com/go/bigquery#InferSchema

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var DATE = registry.Standard("DATE", mtx.Register("civil.Date"))

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DATETIME = registry.Standard("DATETIME", mtx.DATETIME)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var GEOGRAPHY = registry.Standard("GEOGRAPHY", mtx.UNKNOWN)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var INTERVAL = registry.Standard("INTERVAL", mtx.UNKNOWN)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var INT64 = registry.Standard("INT64", mtx.INTEGER)
var INT, SMALLINT, INTEGER, BIGINT, TINYINT, BYTEINT = INT64, INT64, INT64, INT64, INT64, INT64

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Numeric(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N("bq.Datatype", "NUMERIC"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Decimal(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N("bq.Datatype", "DECIMAL"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigNumeric(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N("bq.Datatype", "BIGNUMERIC"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigDecimal(precision, scale int) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N("bq.Datatype", fmt.Sprintf("BIGDECIMAL(%d,%d)", precision, scale)), Extensions: DatatypeExtensions{Scale: scale, Precision: precision},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#timestamp_type
var TIMESTAMP = registry.Standard("TIMESTAMP", mtx.TIMESTAMP)
