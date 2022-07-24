package csv

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/emicklei/mtx"
)

var (
	registry            = mtx.NewTypeRegistry("csv.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	UNKNOWN   = registry.Standard("any", mtx.UNKNOWN)
	BOOLEAN   = registry.Standard("boolean", mtx.BOOLEAN)
	NUMBER    = registry.Standard("number", mtx.DECIMAL)
	STRING    = registry.Standard("string", mtx.STRING)
	TIMESTAMP = registry.Standard("timestamp", mtx.TIMESTAMP)
)

var timestampRegEx = regexp.MustCompile("[0-9][0-9][0-9][0-9]-[0-9][0-9]T[0-9][0-9]:[0-9][0-9].*")

func DetectType(content string) mtx.Datatype {
	if len(content) == 0 {
		return UNKNOWN
	}
	// is it a boolean
	if low := strings.ToLower(content); low == "true" || low == "false" { // language is english!
		return BOOLEAN
	}
	// is it a decimal
	if strings.Contains(content, ".") {
		_, err := strconv.ParseFloat(content, 64)
		if err == nil {
			return NUMBER
		}
	}
	// is it an int?
	_, err := strconv.Atoi(content)
	if err == nil {
		return NUMBER
	}
	// it is a Time? yyyy-mm-ddThh:mm:...
	if timestampRegEx.MatchString(content) {
		return TIMESTAMP
	}
	return STRING
}
