package csv

import "github.com/emicklei/mtx"

// BEGIN: copy from datatypes.go.template

var registry = mtx.NewTypeRegistry("csv.Datatype")

func register(typename string, at mtx.Datatype) mtx.Datatype {
	dt := mtx.Datatype{
		Named:             mtx.N(registry.Class(), typename),
		AttributeDatatype: &at,
	}
	return registry.Add(dt)
}

var RegisterType = registry.RegisterType

// MappedAttributeType returns the mapped proto type for a given attribute type
func MappedAttributeType(at mtx.Datatype) mtx.Datatype {
	return registry.MappedAttributeType(at)
}

var Type = registry.Type

// END: copy from datatypes.go.template
