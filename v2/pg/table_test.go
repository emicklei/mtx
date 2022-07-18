package pg

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestCreate(t *testing.T) {
	tab := NewTable("testtable")
	t.Log("\n", mtx.ToJSON(tab))
}
