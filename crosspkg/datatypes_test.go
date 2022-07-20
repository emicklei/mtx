package crosspkg

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/gcp/spanner"
	"github.com/emicklei/mtx/proto"
)

func TestDefaultDataType(t *testing.T) {
	proto.RegisterType("common.IString", mtx.JSON)
	pt := proto.Type("common.IString")
	st := spanner.MappedAttributeType(*pt.AttributeDatatype)
	t.Log(st)
}
