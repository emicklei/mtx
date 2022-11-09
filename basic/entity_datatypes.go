package basic

import "github.com/emicklei/mtx"

var (
	registry            = mtx.NewTypeRegistry("basic.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

func Register(typename string) mtx.Datatype {
	return registry.Register(typename, true)
}

var (
	Unknown   = registry.Add(mtx.NewBasicType("any"))
	Boolean   = registry.Add(mtx.NewBasicType("boolean"))
	String    = registry.Add(mtx.NewBasicType("string"))
	Integer   = registry.Add(mtx.NewBasicType("integer"))
	ID        = registry.Add(mtx.NewBasicType("identifier"))
	Date      = registry.Add(mtx.NewBasicType("date"))      // yyyy mm dd
	DateTime  = registry.Add(mtx.NewBasicType("datetime"))  // yyyy mm dd hh mm ss
	Timestamp = registry.Add(mtx.NewBasicType("timestamp")) // yyyy mm dd hh mm ss + zone
	Bytes     = registry.Add(mtx.NewBasicType("bytes"))
	Float     = registry.Add(mtx.NewBasicType("float"))
	Double    = registry.Add(mtx.NewBasicType("double"))
	Decimal   = registry.Add(mtx.NewBasicType("decimal"))
	JSON      = registry.Add(mtx.NewBasicType("JSON"))
	Duration  = registry.Add(mtx.NewBasicType("duration")) // y,m,d,h,m,s
	UUID      = registry.Add(mtx.NewBasicType("uuid"))
)
