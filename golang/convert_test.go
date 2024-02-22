package golang

import (
	"testing"

	"github.com/emicklei/mtx/basic"
)

func TestFromBasicType(t *testing.T) {
	dt := basic.Integer.Set("bits", 64)
	FromBasicType(dt)
}
