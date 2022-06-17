package db

type Table struct {
	Columns []Column
}

type Column struct {
	Name     string
	Typename string
}

func NewTable(table string) *Table { return nil }

func (t *Table) Column(name string, columntype ...int)
