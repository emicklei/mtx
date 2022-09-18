package mtx

import "testing"

func TestFieldMap(t *testing.T) {
	from := NewEntity("personapi")
	from.A("id", String, "id")

	mapping := NewEntityMapping(from)
	mapping.SetName("persondb")
	mapping.SetAttributeName("id", "id2")
	mapping.SetAttributeType("id", Integer)

	to := mapping.ToEntity()

	t.Log(to)
	t.Log(to.Attributes[0])
}
