package bq

import "github.com/emicklei/mtx"

type DType = mtx.Datatype[DatatypeExtensions]

var BYTES = DType{Named: mtx.N("bq.Datatype", "BYTES")}

func MaxBytes(max int64) DType {
	return DType{
		Named:      mtx.N("bq.Datatype", "BYTES"),
		Extensions: DatatypeExtensions{Max: max},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var DATE = DType{Named: mtx.N("bq.Datatype", "DATE")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DATETIME = DType{Named: mtx.N("bq.Datatype", "DATETIME")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var GEOGRAPHY = DType{Named: mtx.N("bq.Datatype", "GEOGRAPHY")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var INTERVAL = DType{Named: mtx.N("bq.Datatype", "INTERVAL")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var INT64 = DType{Named: mtx.N("bq.Datatype", "INT64")}
var INT, SMALLINT, INTEGER, BIGINT, TINYINT, BYTEINT = INT64, INT64, INT64, INT64, INT64, INT64

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Numeric(p, s int) DType {
	return DType{
		Named: mtx.N("bq.Datatype", "NUMERIC"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Decimal(p, s int) DType {
	return DType{
		Named: mtx.N("bq.Datatype", "DECIMAL"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigNumeric(p, s int) DType {
	return DType{
		Named: mtx.N("bq.Datatype", "BIGNUMERIC"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigDecimal(p, s int) DType {
	return DType{
		Named: mtx.N("bq.Datatype", "BIGDECIMAL"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}
