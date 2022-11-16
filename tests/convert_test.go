package tests

import (
	"testing"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/gcp/spanner"
	"github.com/emicklei/mtx/golang"
)

func TestDatatypeMappingGolang(t *testing.T) {
	for i, each := range []struct {
		In           mtx.Datatype
		Convert      func(mtx.Datatype) mtx.Datatype
		TypeName     string
		NullTypeName string
	}{
		{
			In:           bq.String,
			Convert:      bq.ToBasicType,
			TypeName:     "string",
			NullTypeName: "bigquery.NullString",
		},
		{
			In:           bq.Bytes,
			Convert:      bq.ToBasicType,
			TypeName:     "[]byte",
			NullTypeName: "[]byte",
		},
		{
			In:           bq.JSON,
			Convert:      bq.ToBasicType,
			TypeName:     "string",
			NullTypeName: "bigquery.NullString",
		},
		{
			In:           bq.Timestamp,
			Convert:      bq.ToBasicType,
			TypeName:     "time.Time",
			NullTypeName: "bigquery.NullTimestamp",
		},
		{
			In:           bq.Decimal(10, 2),
			Convert:      bq.ToBasicType,
			TypeName:     "*big.Rat",
			NullTypeName: "*big.Rat",
		},
		{
			In:           spanner.String,
			Convert:      spanner.ToBasicType,
			TypeName:     "string",
			NullTypeName: "spanner.NullString",
		},
		{
			In:           spanner.Bool,
			Convert:      spanner.ToBasicType,
			TypeName:     "bool",
			NullTypeName: "spanner.NullBool",
		},
	} {
		in := each.Convert(each.In)
		gt := golang.FromBasicType(in)
		if got, want := gt.Name, each.TypeName; got != want {
			t.Errorf("%d:%v got [%v]:%T want [%v]:%T", i, each.In, got, got, want, want)
		}
		nullin := each.In.WithNullable()
		in = each.Convert(nullin)
		gt = golang.FromBasicType(in)
		if got, want := gt.Name, each.NullTypeName; got != want {
			t.Errorf("%d:%v got [%v]:%T want nullable [%v]:%T", i, each.In, got, got, want, want)
		}
	}
}

func TestBigQueryTableToGolangStruct(t *testing.T) {
	ds := bq.NewDataset("test")
	tab := ds.Table("test")
	tab.C("name", bq.String, "test")
	tab.C("nullname", bq.String, "test").Nullable()
	e := tab.ToEntity()
	if got, want := e.Attributes[1].AttributeType.Name, "string"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := e.Attributes[1].AttributeType.IsNullable, true; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	s := golang.ToStruct(e)
	t.Log(s.ToGo())
	if got, want := s.Name, "Test"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := s.Fields[0].Name, "Name"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := s.Fields[0].FieldType.Name, "string"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	if got, want := s.Fields[1].FieldType.Name, "bigquery.NullString"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}
