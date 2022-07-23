package tests

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/proto"
)

func TestPackageMessage(t *testing.T) {
	pkg := proto.NewPackage("my_pkg")
	msg := pkg.Message("MyMessage").Doc("Sample proto Message")
	msg.F("id", 1, proto.STRING, "id of the message")
	t.Log("\n", mtx.ToJSON(msg))
	t.Log("\n", mtx.ToSource(pkg))

	// create entity from proto message
	e := msg.ToEntity()
	t.Log("\n", mtx.ToJSON(e))
	t.Log("\n", golang.ToStruct(e).Go())
}
