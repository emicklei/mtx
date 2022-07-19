package mtx

import "io"

type TypedLabel interface {
	GetName() string
	GetDatatype() Datatype
	// IsRequired() bool
}

type SQLWriter interface{ SQLOn(w io.Writer) }
