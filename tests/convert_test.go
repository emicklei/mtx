package tests

import (
	"testing"

	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/golang"
)

func TestNullableString(t *testing.T) {
	bs := bq.String.Nullable()
	es := bq.ToBasicType(bs)
	gt := golang.Datatype(es)
	t.Log(bs, es, gt)
}
