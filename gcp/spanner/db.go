package spanner

import "github.com/emicklei/mtx"

func NewDatabase(name string) *mtx.Database {
	return &mtx.Database{
		Named:      mtx.N("spanner.Database", name),
		Extensions: new(DatabaseExtensions),
	}
}
