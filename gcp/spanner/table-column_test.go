package spanner

import (
	"testing"
)

func TestTable(t *testing.T) {
	tab := new(Database).Table("test")
	tab.Set("custom", "prop")
	tab.Column("col").Type(BigInteger)
	tab.ToEntity()
}
