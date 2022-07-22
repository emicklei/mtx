package db

import "github.com/emicklei/mtx"

type DatabaseView struct {
	*mtx.Named
	ColumnExpressions []any
}
