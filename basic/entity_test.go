package basic

import (
	"testing"

	"github.com/emicklei/mtx"
)

func TestCreateEntity(t *testing.T) {
	m := NewEntity("Person")
	m.Attribute("id").Type(String)
	m.Attribute("age").Type(Integer)
	m.Attribute("birthdate").Type(Date)
	t.Log(mtx.ToJSON(m))
}

// func TestDatatypeMapping(t *testing.T) {
// 	dt := mtx.Datatype{
// 		Named: N("mtx.Datatype", "string"),
// 	}
// 	dt.Set("bigquery", "STRING")
// 	dt.Set("pg", "STRING")
// }
