package tests

import (
	"testing"

	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/pg"
)

func TestNullableString(t *testing.T) {
	bs := bq.String.Nullable()
	es := bq.ToBasicType(bs)
	gt := golang.FromBasicType(es)
	pt := pg.FromBasicType(es)
	t.Log(bs)
	t.Log(es)
	t.Log(gt)
	t.Log(pt)
}
