package spanner

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestTable(t *testing.T) {
	tab := new(Database).Table("test")
	tab.Set("custom", "prop")
	tab.Column("col").Type(BigInteger)
	core.JSONOut(tab)
	core.JSONOut(tab.ToEntity())
}
