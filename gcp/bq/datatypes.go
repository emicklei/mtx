package bq

import (
	"fmt"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var (
	registry            = mtx.NewTypeRegistry("bq.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	Bytes  = registry.Standard("BYTES", basic.Bytes)
	String = registry.Standard("STRING", basic.String, mtx.GoNullableTypeName, "bigquery.NullString")
	Record = registry.Standard("RECORD", mtx.Unknown)
	Bool   = registry.Standard("BOOL", basic.Boolean, mtx.GoNullableTypeName, "bigquery.NullBool")
)

func init() {
	registry.EncodeAs(basic.JSON, basic.Bytes)
}

func MaxBytes(max int64) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N("bq.Datatype", "BYTES"),
		Extensions: DatatypeExtensions{Max: max},
	}
}

var (
	// https://cloud.google.com/bigquery/docs/reference/standard-sql/json-data#sql
	JSON = registry.Standard("JSON", basic.JSON).Set(mtx.GoNullableTypeName, "bigquery.NullString")
)

// TODO look at civil package https://pkg.go.dev/cloud.google.com/go/bigquery#InferSchema

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var Date = registry.Standard("DATE", basic.Date).
	Set(mtx.GoNullableTypeName, "bigquery.NullDate").
	Set(mtx.GoName, "civil.Date")

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DateTime = registry.Standard("DATETIME", basic.DateTime).Set(mtx.GoNullableTypeName, "bigquery.NullDateTime")

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var Geography = registry.Standard("GEOGRAPHY", basic.Unknown)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var Interval = registry.Standard("INTERVAL", basic.Unknown)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var Int64 = registry.Standard("INT64", basic.Integer).Set("bits", 64)
var Int, SmallInt, Integer, BigInt, TinyInt, ByteInt = Int64, Int64, Int64, Int64, Int64, Int64

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Numeric(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named:         mtx.N("bq.Datatype", "NUMERIC"),
		Extensions:    DatatypeExtensions{Scale: s, Precision: p},
		BasicDatatype: &basic.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Decimal(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named:         mtx.N("bq.Datatype", "DECIMAL"),
		Extensions:    DatatypeExtensions{Scale: s, Precision: p},
		BasicDatatype: &basic.Decimal,
	}.Set(mtx.GoNullableTypeName, "*big.Rat").Set(mtx.GoName, "*big.Rat")
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigNumeric(precision, scale int) mtx.Datatype {
	return mtx.Datatype{
		Named:         mtx.N("bq.Datatype", fmt.Sprintf("BIGNUMERIC(%d,%d)", precision, scale)),
		Extensions:    DatatypeExtensions{Scale: scale, Precision: precision},
		BasicDatatype: &basic.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigDecimal(precision, scale int) mtx.Datatype {
	return mtx.Datatype{
		Named:         mtx.N("bq.Datatype", fmt.Sprintf("BIGDECIMAL(%d,%d)", precision, scale)),
		Extensions:    DatatypeExtensions{Scale: scale, Precision: precision},
		BasicDatatype: &basic.Decimal,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#timestamp_type
var Timestamp = registry.Standard("TIMESTAMP", basic.Timestamp).Set(mtx.GoNullableTypeName, "bigquery.NullTimestamp")
