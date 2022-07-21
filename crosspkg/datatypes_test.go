package crosspkg

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/gcp/spanner"
	"github.com/emicklei/mtx/proto"
)

func TestDefaultDataType(t *testing.T) {
	proto.RegisterType("common.IString", mtx.JSON)
	pt := proto.Type("common.IString")
	st := spanner.MappedAttributeType(*pt.AttributeDatatype)
	t.Log(st)
}

func TestBQStringMapsToGoString(t *testing.T) {
	tab := bq.NewDataset("test").Table("test")
	tab.C("s", bq.STRING, "")
	ent := tab.ToEntity()
	s := ent.Attributes[0]
	if got, want := s.AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestNullableBQStringBecomesNullString(t *testing.T) {
	tab := bq.NewDataset("test").Table("test")
	tab.C("s", bq.STRING, "").Nullable()
	ent := tab.ToEntity()
	s := ent.Attributes[0]
	if got, want := s.AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}
