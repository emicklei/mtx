package db

type TableMigration struct {
	Base *Table
	Name string
}

func (t *Table) Migration(name string) *TableMigration {
	return &TableMigration{Base: t, Name: name}
}
