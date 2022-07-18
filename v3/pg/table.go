package pg

import (
	"io"

	"github.com/emicklei/mtx"
	v3 "github.com/emicklei/mtx/v3"
)

type TableExtensions struct{}

func (e *TableExtensions) OwnerClass() string           { return "" }
func (e *TableExtensions) SQLOn(table any, w io.Writer) {}

func NewTable(name string) *v3.Table {
	return &v3.Table{
		Named:      mtx.N("pg.Table", name),
		Extensions: new(TableExtensions),
	}
}
