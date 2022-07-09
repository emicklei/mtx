package golang

import "github.com/emicklei/mtx"

var attributeTypeToGoTypeMapping = map[string]string{
	mtx.INTEGER.Name:   "int64",
	mtx.BOOLEAN.Name:   "bool",
	mtx.BYTES.Name:     "[]byte",
	mtx.DATE.Name:      "time.Time",
	mtx.TIMESTAMP.Name: "time.Time",
	mtx.FLOAT.Name:     "float32",
	mtx.DOUBLE.Name:    "float64",
	mtx.STRING.Name:    "string",
	mtx.UUID.Name:      "uuid.UUID",
	//mtx.DECIMAL: TODO
}
