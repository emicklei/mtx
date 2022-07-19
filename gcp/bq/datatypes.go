package bq

import "github.com/emicklei/mtx"

// BEGIN: copy from datatypes.go.template

var registry = mtx.NewTypeRegistry("bq.Datatype")

func register(typename string, at mtx.AttributeType) mtx.Datatype {
	return registry.Register(typename, at, false)
}

func RegisterType(typename string, at mtx.AttributeType) mtx.Datatype {
	return registry.Register(typename, at, true)
}

func MappedAttributeType(at mtx.AttributeType) mtx.Datatype {
	return registry.MappedAttributeType(at)
}

func Type(typename string) mtx.Datatype {
	dt, ok := registry.TypeNamed(typename)
	if ok {
		return dt
	}
	return registry.Register(typename, mtx.UNKNOWN, true)
}

// END: copy from datatypes.go.template

var BYTES = register("BYTES", mtx.BYTES)

func MaxBytes(max int64) mtx.Datatype {
	return mtx.Datatype{
		Named:      mtx.N("bq.Datatype", "BYTES"),
		Extensions: DatatypeExtensions{Max: max},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var DATE = register("DATE", mtx.DATE)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DATETIME = register("DATETIME", mtx.DATETIME)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var GEOGRAPHY = register("GEOGRAPHY", mtx.STRING)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var INTERVAL = register("INTERVAL", mtx.UNKNOWN)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#integer_type
var INT64 = mtx.Datatype{Named: mtx.N("bq.Datatype", "INT64")}
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
func BigDecimal(p, s int) mtx.Datatype {
	return mtx.Datatype{
		Named: mtx.N("bq.Datatype", "BIGDECIMAL"), Extensions: DatatypeExtensions{Scale: s, Precision: p},
	}
}
