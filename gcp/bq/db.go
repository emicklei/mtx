package bq

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
)

func NewDataset(name string) *db.Database {
	return &db.Database{
		Named:      mtx.N("spanner.Dataset", name),
		Extensions: new(DatabaseExtensions),
	}
}
