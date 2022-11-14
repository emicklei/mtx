package tests

import (
	"testing"

	"github.com/emicklei/mtx/basic"
	"github.com/emicklei/mtx/gcp/bq"
	"github.com/emicklei/mtx/gcp/spanner"
	"github.com/emicklei/mtx/golang"
)

func TestNullableString(t *testing.T) {
	bs := bq.String.Nullable()
	es := bq.ToBasicType(bs)
	gt := golang.FromBasicType(es)
	t.Log("bq", bs)
	t.Log("bq->basic", es)
	t.Log("bq->golang:", gt)
	if got, want := gt.Name, "bigquery.NullString"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	{
		ss := spanner.String.Nullable()
		es := spanner.ToBasicType(ss)
		gt := golang.FromBasicType(es)
		t.Log("ss", ss)
		t.Log("spanner->basic", es)
		t.Log("spanner->golang:", gt)
		if got, want := gt.Name, "spanner.NullString"; got != want {
			t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
		}

	}
}

func TestDecimal(t *testing.T) {
	in := bq.Decimal(10, 2)
	b := bq.ToBasicType(in)
	gt := golang.FromBasicType(b)
	if got, want := gt.Name, "*big.Rat"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestTime(t *testing.T) {
	in := bq.Timestamp
	b := bq.ToBasicType(in)
	if got, want := b.Name, basic.Timestamp.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	gt := golang.FromBasicType(b)
	if got, want := gt.Name, "time.Time"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestBQ_JSON(t *testing.T) {
	in := bq.JSON
	b := bq.ToBasicType(in)
	if got, want := b.Name, basic.JSON.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	gt := golang.FromBasicType(b)
	if got, want := gt.Name, "string"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
}

func TestBQ_Bytes(t *testing.T) {
	in := bq.Bytes
	b := bq.ToBasicType(in)
	if got, want := b.Name, basic.Bytes.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	gt := golang.FromBasicType(b)
	if got, want := gt.Name, "[]byte"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
	}
	in = in.Nullable()
	b = bq.ToBasicType(in)
	if got, want := b.Name, basic.Bytes.Name; got != want {
		t.Errorf("got [%v]:%T want [%v]:%T", got, got, want, want)
	}
	gt = golang.FromBasicType(b)
	if got, want := gt.Name, "[]byte"; got != want {
		t.Errorf("got [%v:%T] want [%v:%T]", got, got, want, want)
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
