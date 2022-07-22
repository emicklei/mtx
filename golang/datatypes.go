package golang

import "github.com/emicklei/mtx"

// BEGIN: copy from datatypes.go.template

var registry = mtx.NewTypeRegistry("golang.Datatype")

func register(typename string, at mtx.Datatype) mtx.Datatype {
	dt := mtx.Datatype{
		Named:             mtx.N("golang.Datatype", typename),
		AttributeDatatype: &at,
	}
	return registry.Add(dt)
}

func RegisterType(typename string, at mtx.Datatype) mtx.Datatype {
	dt := mtx.Datatype{
		Named:             mtx.N("golang.Datatype", typename),
		AttributeDatatype: &at,
		IsUserDefined:     true,
	}
	return registry.Add(dt)
}

// MappedAttributeType returns the mapped proto type for a given attribute type
func MappedAttributeType(at mtx.Datatype) mtx.Datatype {
	return registry.MappedAttributeType(at)
}

func Type(typename string) mtx.Datatype {
	dt, ok := registry.TypeNamed(typename)
	if ok {
		return dt
	}
	return RegisterType(typename, mtx.UNKNOWN)
}

// END: copy from datatypes.go.template

var (
	STRING  = register("string", mtx.STRING)
	BOOL    = register("bool", mtx.BOOLEAN)
	BYTES   = register("[]byte", mtx.BYTES)
	TIME    = register("time.Time", mtx.TIMESTAMP)
	FLOAT32 = register("float32", mtx.FLOAT)
	FLOAT64 = register("float64", mtx.FLOAT)
	INT     = register("int", mtx.INTEGER)
)

func init() {
	registry.EncodeAs(mtx.DATE, TIME)
}
