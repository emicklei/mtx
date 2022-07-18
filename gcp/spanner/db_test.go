package spanner

import (
	"testing"
)

func TestTable(t *testing.T) {
	tab := NewDatabase("testdb").Table("test")
	tab.Set("custom", "prop")
	tab.Column("col").Type(BigInteger)
	tab.ToEntity()
}
