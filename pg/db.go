package pg

import "github.com/emicklei/mtx"

func NewDatabase(name string) *mtx.Database {
	return &mtx.Database{
		Named:      mtx.N("pg.Database", name),
		Extensions: new(DatabaseExtensions),
	}
}
