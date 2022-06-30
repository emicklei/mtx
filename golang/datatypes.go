package golang

import (
	"github.com/emicklei/mtx/core"
)

var attributeTypeToGoTypeMapping = map[string]string{
	core.INTEGER.Name:   "int64",
	core.BOOLEAN.Name:   "bool",
	core.BYTES.Name:     "[]byte",
	core.DATE.Name:      "time.Time",
	core.TIMESTAMP.Name: "time.Time",
	core.FLOAT.Name:     "float32",
	core.DOUBLE.Name:    "float64",
	core.STRING.Name:    "string",
	//core.DECIMAL: TODO
}
