package v3

import "github.com/emicklei/mtx"

type Table struct {
	*mtx.Named
	Extensions mtx.ExtendsTable `json:"ext"`
	Columns    []*Column        `json:"columns,omitempty"`
}

type Column struct {
	*mtx.Named
	Extensions mtx.ExtendsColumn
}
