package pg

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestTableCreate(t *testing.T) {
	tab := NewTable("testje")
	t.Log(mtx.ToJSON(tab))
}
