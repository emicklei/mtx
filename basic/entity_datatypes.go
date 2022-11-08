package basic

import "github.com/emicklei/mtx"

var (
	registry            = mtx.NewTypeRegistry("mtx.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

func Register(typename string) mtx.Datatype {
	return registry.Register(typename, true)
}

var (
	Unknown   = registry.Add(mtx.NewAttributeType("any"))
	Boolean   = registry.Add(mtx.NewAttributeType("boolean"))
	String    = registry.Add(mtx.NewAttributeType("string"))
	Integer   = registry.Add(mtx.NewAttributeType("integer"))
	ID        = registry.Add(mtx.NewAttributeType("identifier"))
	Date      = registry.Add(mtx.NewAttributeType("date"))      // yyyy mm dd
	DateTime  = registry.Add(mtx.NewAttributeType("datetime"))  // yyyy mm dd hh mm ss
	Timestamp = registry.Add(mtx.NewAttributeType("timestamp")) // yyyy mm dd hh mm ss + zone
	Bytes     = registry.Add(mtx.NewAttributeType("bytes"))
	Float     = registry.Add(mtx.NewAttributeType("float"))
	Double    = registry.Add(mtx.NewAttributeType("double"))
	Decimal   = registry.Add(mtx.NewAttributeType("decimal"))
	JSON      = registry.Add(mtx.NewAttributeType("JSON"))
	Duration  = registry.Add(mtx.NewAttributeType("duration")) // y,m,d,h,m,s
	UUID      = registry.Add(mtx.NewAttributeType("uuid"))
)
