package spanner

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

func NewDatabase(name string) *db.Database {
	return &db.Database{
		Named:      mtx.N("spanner.Database", name),
		Extensions: new(DatabaseExtensions),
	}
}
