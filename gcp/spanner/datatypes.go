package spanner

import "github.com/emicklei/mtx/core"

var BigInteger = core.Datatype[DatatypeExtensions]{
	Named:      core.N("spanner.Datatype", "BIGINT"),
	Extensions: DatatypeExtensions{Max: 1024},
}
