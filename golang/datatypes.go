package golang

import "github.com/emicklei/mtx"

var (
	registry            = mtx.NewTypeRegistry("golang.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	STRING  = registry.Standard("string", mtx.STRING)
	BOOL    = registry.Standard("bool", mtx.BOOLEAN)
	BYTES   = registry.Standard("[]byte", mtx.BYTES)
	TIME    = registry.Standard("time.Time", mtx.TIMESTAMP)
	FLOAT32 = registry.Standard("float32", mtx.FLOAT)
	FLOAT64 = registry.Standard("float64", mtx.FLOAT)
	INT     = registry.Standard("int", mtx.INTEGER)

	MAP_STRING_ANY = registry.Standard("map[string]any", mtx.JSON)
)

func init() {
	registry.EncodeAs(mtx.DATE, TIME)
	registry.EncodeAs(mtx.TIMESTAMP, TIME)
	registry.EncodeAs(mtx.DATETIME, TIME)
}
