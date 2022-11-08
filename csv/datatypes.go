package csv

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

var (
	registry            = mtx.NewTypeRegistry("csv.Datatype")
	Type                = registry.Type
	RegisterType        = registry.RegisterType
	MappedAttributeType = registry.MappedAttributeType
)

var (
	Unknown   = registry.Standard("any", mtx.Unknown)
	Boolean   = registry.Standard("boolean", basic.Boolean)
	Number    = registry.Standard("number", basic.Decimal)
	String    = registry.Standard("string", basic.String)
	Timestamp = registry.Standard("timestamp", basic.Timestamp)
)

var timestampRegEx = regexp.MustCompile("[0-9][0-9][0-9][0-9]-[0-9][0-9]T[0-9][0-9]:[0-9][0-9].*")

func DetectType(content string) mtx.Datatype {
	if len(content) == 0 {
		return Unknown
	}
	// is it a boolean
	if low := strings.ToLower(content); low == "true" || low == "false" { // language is english!
		return Boolean
	}
	// is it a decimal
	if strings.Contains(content, ".") {
		_, err := strconv.ParseFloat(content, 64)
		if err == nil {
			return Number
		}
	}
	// is it an int?
	_, err := strconv.Atoi(content)
	if err == nil {
		return Number
	}
	// it is a Time? yyyy-mm-ddThh:mm:...
	if timestampRegEx.MatchString(content) {
		return Timestamp
	}
	return String
}
