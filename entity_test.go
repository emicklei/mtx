package mtx

import (
	"testing"
)

func TestCreateEntity(t *testing.T) {
	m := NewEntity("Person")
	m.Attribute("id").Type(String)
	m.Attribute("age").Type(Integer)
	m.Attribute("birthdate").Type(Date)
	t.Log(ToJSON(m))
}

// func TestDatatypeMapping(t *testing.T) {
// 	dt := mtx.Datatype{
// 		Named: N("mtx.Datatype", "string"),
// 	}
// 	dt.Set("bigquery", "STRING")
// 	dt.Set("pg", "STRING")
// }
