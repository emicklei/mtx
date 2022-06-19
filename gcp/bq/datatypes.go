package bq

import "github.com/emicklei/mtx/core"

type Datatype struct {
	core.Named
	Max int64
}

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
