package db

import (
	"io"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

type Extensions struct {
	ownerClass    string
	ColumClass    string `json:"-"`
	DatatypeClass string `json:"-"`
}

func (e *Extensions) withOwner(class string) *Extensions {
	return &Extensions{ownerClass: class}
}
func (e *Extensions) ColumnExtensions() ExtendsColumn                  { return e.withOwner(e.ColumClass) }
func (e *Extensions) Datatype() ExtendsDatatype                        { return e.withOwner(e.DatatypeClass) }
func (e *Extensions) OwnerClass() string                               { return e.ownerClass }
func (e *Extensions) PostBuildAttribute(*Column, *basic.Attribute)     {}
func (e *Extensions) SQLOn(tab *Table, w io.Writer)                    {}
func (e *Extensions) ValidateTable(tab *Table, ec *mtx.ErrorCollector) {}
