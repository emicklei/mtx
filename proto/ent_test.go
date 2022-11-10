package proto

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

func TestMsgToEntity(t *testing.T) {
	p := NewPackage("test")
	m := p.Message("Test")
	m.F("name", 1, String, "to call")

	e := ToEntity(m)
	if got, want := e.Attribute("name").AttributeType.Name, basic.String.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	t.Log(mtx.ToJSON(e))
}
