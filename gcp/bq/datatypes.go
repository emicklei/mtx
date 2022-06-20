package bq

import "github.com/emicklei/mtx/core"

var Bytes = Datatype{Named: core.N("bq.Datatype", "BYTES")}

func MaxBytes(max int64) Datatype {
	return Datatype{Named: core.N("bq.Datatype", "BYTES"), Max: max}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var Date = Datatype{Named: core.N("bq.Datatype", "DATE")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DateTime = Datatype{Named: core.N("bq.Datatype", "DATETIME")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var Geography = Datatype{Named: core.N("bq.Datatype", "GEOGRAPHY")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var Interval = Datatype{Named: core.N("bq.Datatype", "Interval")}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var INT64 = Datatype{Named: core.N("bq.Datatype", "INT64")}
var INT, SMALLINT, INTEGER, BIGINT, TINYINT, BYTEINT = INT64, INT64, INT64, INT64, INT64, INT64

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Numeric(p, s int) Datatype {
	return Datatype{
		Named: core.N("bq.Datatype", "NUMERIC"), Scale: s, Precision: p,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func Decimal(p, s int) Datatype {
	return Datatype{
		Named: core.N("bq.Datatype", "DECIMAL"), Scale: s, Precision: p,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigNumeric(p, s int) Datatype {
	return Datatype{
		Named: core.N("bq.Datatype", "BIGNUMERIC"), Scale: s, Precision: p,
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#parameterized_decimal_type
func BigDecimal(p, s int) Datatype {
	return Datatype{
		Named: core.N("bq.Datatype", "BIGDECIMAL"), Scale: s, Precision: p,
	}
}
