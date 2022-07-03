package mtx

import (
	"testing"
)

func TestCreateEntity(t *testing.T) {
	m := NewEntity("Person")
	m.Attribute("id").Type(STRING)
	m.Attribute("age").Type(INTEGER)
	m.Attribute("birthdate").Type(DATE)
	t.Log(ToJSON(m))
}
