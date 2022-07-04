package bq

import "github.com/emicklei/mtx"

// BEGIN: copy from datatypes.go.template
type DType = mtx.Datatype[DatatypeExtensions]

var registry = mtx.NewTypeRegistry[DType]()

func register(typename string, at mtx.AttributeType, isUserDefined bool) DType {
	dt := DType{
		Named:         mtx.N("bq.Datatype", typename),
		IsUserDefined: isUserDefined,
	}.WithAttributeType(at)
	return registry.Add(dt)
}

func MappedAttributeType(at mtx.AttributeType) DType {
	return registry.MappedAttributeType(at)
}

func Type(name string) DType {
	dt, ok := registry.TypeNamed(name)
	if ok {
		return dt
	}
	return register(name, mtx.UNKNOWN, true)
}

// END: copy from datatypes.go.template

var BYTES = DType{Named: mtx.N("bq.Datatype", "BYTES")}

func MaxBytes(max int64) DType {
	return DType{
		Named:      mtx.N("bq.Datatype", "BYTES"),
		Extensions: DatatypeExtensions{Max: max},
	}
}

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#date_type
// YYYY-[M]M-[D]D
var DATE = register("DATE", mtx.DATE, mtx.StandardType)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#datetime_type
// YYYY-[M]M-[D]D[( |T)[H]H:[M]M:[S]S[.F]]
var DATETIME = register("DATETIME", mtx.DATETIME, mtx.StandardType)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#geography_type
var GEOGRAPHY = register("GEOGRAPHY", mtx.STRING, mtx.StandardType)

// https://cloud.google.com/bigquery/docs/reference/standard-sql/data-types#interval_type
// [sign]Y-M [sign]D [sign]H:M:S[.F]
var INTERVAL = register("INTERVAL", mtx.UNKNOWN, mtx.UserDefinedType)

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
