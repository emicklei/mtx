package tests

import (
	"testing"

	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/proto"
)

func TestGoStructFromPackageMessage(t *testing.T) {
	pkg := proto.NewPackage("my_pkg")
	msg := pkg.Message("MyMessage").Doc("Sample proto Message")
	msg.F("id", 1, proto.String, "id of the message")

	// create entity from proto message
	e := msg.ToEntity()

	if got, want := golang.ToStruct(e).Go(), `// MyMessage : Sample proto Message
type MyMessage struct {
	Id string // id of the message
}
`; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
