package spanner

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestTable(t *testing.T) {
	tab := new(Database).Table("test")
	tab.Column("col").Datatype(BigInteger)
	core.JSONOut(tab)
}
