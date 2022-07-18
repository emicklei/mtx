package bq

import "github.com/emicklei/mtx"

func NewDataset(name string) *mtx.Database {
	return &mtx.Database{
		Named:      mtx.N("spanner.Dataset", name),
		Extensions: new(DatabaseExtensions),
	}
}
