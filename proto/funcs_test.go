package proto

import (
	"testing"

	"github.com/emicklei/mtx/core"
)

func TestPackageSourceOn(t *testing.T) {
	pkg := NewPackage("pkg").Doc("pkg-doc")
	msg := pkg.Message("MSG").Doc("msg-doc")
	msg.F("field", 1, INT32, "field-doc")
	if got, want := core.ToSource(pkg), `pkg := proto.NewPackage("pkg").Doc("pkg-doc")
msg := pkg.Message("MSG").Doc("msg-doc")
msg.F("field",1,proto.INT32,"field-doc")`; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
