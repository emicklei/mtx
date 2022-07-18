package v2

import "github.com/emicklei/mtx"

type Table struct {
	*mtx.Named
	Columns []*Column
}

type Column struct {
	*mtx.Named
}

type Datatype struct {
	*mtx.Named
}
