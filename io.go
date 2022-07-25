package mtx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func ToJSON(what any) string {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetIndent("", "  ")
	enc.Encode(what)
	return buf.String()
}

func ToSource(what SourceWriteable) string {
	buf := new(bytes.Buffer)
	what.SourceOn(buf)
	return buf.String()
}

type SourceWriteable interface {
	SourceOn(io.Writer)
}

func Validate(v Validates) (*ErrorCollector, bool) {
	c := new(ErrorCollector)
	v.Validate(c)
	return c, len(c.list) > 0
}

type ErrorCollector struct {
	ns   string
	list []ErrorWithOrigin
}

type ErrorWithOrigin struct {
	Origin *Named
	Err    error
}

func (e *ErrorCollector) Print() {
	for _, each := range e.list {
		fmt.Println(each.Origin.Name, each.Err)
	}
}
func (e *ErrorCollector) Add(who *Named, err error) {
	e.list = append(e.list, ErrorWithOrigin{who, err})
}
