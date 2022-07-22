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

func TestNullableBQStringEntityString(t *testing.T) {
	tab := bq.NewDataset("test").Table("test")
	tab.C("s", bq.STRING, "").Nullable()
	ent := tab.ToEntity()
	s := ent.Attributes[0]
	if got, want := s.AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestProtoEncodedBytesForSpanner(t *testing.T) {
	et := proto.BYTES.EncodedFrom(mtx.JSON)
	st := spanner.MappedAttributeType(*et.AttributeDatatype)
	if got, want := st, spanner.JSON; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	{
		st := spanner.MappedAttributeType(mtx.DURATION)
		if got, want := st, spanner.Type("STRING(MAX)"); got != want {
			t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
		}
	}
}
