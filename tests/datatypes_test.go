package tests

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/gcp/spanner"
	"github.com/emicklei/mtx/golang"
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
	tab.C("s", bq.String, "")
	ent := tab.ToEntity()
	s := ent.Attributes[0]
	if got, want := s.AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
}

func TestNullableBQStringEntityString(t *testing.T) {
	tab := bq.NewDataset("test").Table("test")
	tab.C("s", bq.String, "").Nullable()
	ent := tab.ToEntity()
	str := golang.ToStruct(ent, golang.WithBigQueryTypeMapper)
	if got, want := str.Fields[0].Name, "S"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := str.Fields[0].FieldType.Name, "bigquery.NullString"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestProtoEncodedBytesForSpanner(t *testing.T) {
	et := proto.Bytes.EncodedFrom(mtx.JSON)
	st := spanner.MappedAttributeType(*et.AttributeDatatype)
	if got, want := st, spanner.JSON; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	{
		st := spanner.MappedAttributeType(mtx.Duration)
		if got, want := st, spanner.Type("STRING(MAX)"); got != want {
			t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
		}
	}
}
