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
	Unknown = registry.Standard("any", mtx.Unknown)
	Bytes   = registry.Standard("BYTES", mtx.Bytes)
	String  = registry.Standard("STRING", mtx.String).WithNullable(mtx.Register("bigquery.NullString"))
	Record  = registry.Standard("RECORD", mtx.Unknown)
)

func init() {
	registry.EncodeAs(mtx.JSON, mtx.Bytes)
}

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
var Date = registry.Standard("DATE", mtx.Register("civil.Date"))

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DateTime = registry.Standard("DATETIME", mtx.DateTime)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var Geography = registry.Standard("GEOGRAPHY", mtx.Unknown)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var Interval = registry.Standard("INTERVAL", mtx.Unknown)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var Int64 = registry.Standard("INT64", mtx.Integer).Set("bits", 64)
var Int, SmallInt, Integer, BigInt, TinyInt, ByteInt = Int64, Int64, Int64, Int64, Int64, Int64

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Numeric(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named:             mtx.N("bq.Datatype", "NUMERIC"),
		Extensions:        DatatypeExtensions{Scale: s, Precision: p},
		AttributeDatatype: &mtx.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Decimal(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named:             mtx.N("bq.Datatype", "DECIMAL"),
		Extensions:        DatatypeExtensions{Scale: s, Precision: p},
		AttributeDatatype: &mtx.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigNumeric(precision, scale int) mtx.Datatype {
	return mtx.Datatype{
		Named:             mtx.N("bq.Datatype", fmt.Sprintf("BIGNUMERIC(%d,%d)", precision, scale)),
		Extensions:        DatatypeExtensions{Scale: scale, Precision: precision},
		AttributeDatatype: &mtx.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigDecimal(precision, scale int) mtx.Datatype {
	return mtx.Datatype{
		Named:             mtx.N("bq.Datatype", fmt.Sprintf("BIGDECIMAL(%d,%d)", precision, scale)),
		Extensions:        DatatypeExtensions{Scale: scale, Precision: precision},
		AttributeDatatype: &mtx.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#timestamp_type
var Timestamp = registry.Standard("TIMESTAMP", mtx.Timestamp)
