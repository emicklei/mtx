package bq

import "github.com/emicklei/mtx/core"

type DType = core.Datatype[DatatypeExtensions]

var Bytes = DType{Named: core.N("bq.Datatype", "BYTES")}

func MaxBytes(max int64) DType {
	return DType{
		Named:      core.N("bq.Datatype", "BYTES"),
		Extensions: DatatypeExtensions{Max: max},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var Date = DType{Named: core.N("bq.Datatype", "DATE")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DateTime = DType{Named: core.N("bq.Datatype", "DATETIME")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var Geography = DType{Named: core.N("bq.Datatype", "GEOGRAPHY")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var Interval = DType{Named: core.N("bq.Datatype", "Interval")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var INT64 = DType{Named: core.N("bq.Datatype", "INT64")}
var INT, SMALLINT, INTEGER, BIGINT, TINYINT, BYTEINT = INT64, INT64, INT64, INT64, INT64, INT64

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Numeric(p, s int) DType {
	return DType{
		Named: core.N("bq.Datatype", "NUMERIC"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Decimal(p, s int) DType {
	return DType{
		Named: core.N("bq.Datatype", "DECIMAL"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigNumeric(p, s int) DType {
	return DType{
		Named: core.N("bq.Datatype", "BIGNUMERIC"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigDecimal(p, s int) DType {
	return DType{
		Named: core.N("bq.Datatype", "BIGDECIMAL"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}
