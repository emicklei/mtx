package mtx

import (
	"fmt"
)

// Returns nil if valid
func Validate(v Validates) *ErrorCollector {
	c := new(ErrorCollector)
	v.Validate(c)
	if len(c.list) > 0 {
		return c
	}
	return nil
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

func MustBeValid(v Validates) {
	if ec := Validate(v); ec != nil {
		ec.Print()
		panic("invalid")
	}
}
