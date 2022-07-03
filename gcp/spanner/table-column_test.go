package spanner

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestTable(t *testing.T) {
	tab := new(Database).Table("test")
	tab.Set("custom", "prop")
	tab.Column("col").Type(BigInteger)
	mtx.JSONOut(tab)
	mtx.JSONOut(tab.ToEntity())
}
