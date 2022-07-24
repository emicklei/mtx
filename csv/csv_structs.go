package csv

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

type Sheet struct {
	*mtx.Named
	Tabs []*db.Table
}

func NewSheet(name string) *Sheet {
	return &Sheet{
		Named: mtx.N("csv.Sheet", name),
	}
}

func (s *Sheet) Tab(name string) *db.Table {
	if t, ok := mtx.FindByName(s.Tabs, name); ok {
		return t
	}
	t := &db.Table{
		Named: mtx.N("csv.Tab", name),
		Extensions: &db.Extensions{
			ColumClass:    "csv.Column",
			DatatypeClass: "csv.Datatype",
		},
	}
	s.Tabs = append(s.Tabs, t)
	return t
}
