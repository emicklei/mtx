package m2m

import (
	"testing"

	"github.com/emicklei/mtx/basic"
)

func TestFieldMap(t *testing.T) {
	from := basic.NewEntity("personapi")
	from.A("id", basic.String, "id")

	mapping := NewEntityMapping(from)
	mapping.SetName("persondb")
	mapping.SetAttributeName("id", "id2")
	mapping.SetAttributeType("id", basic.Integer)

	to := mapping.ToEntity()

	t.Log(to)
	t.Log(to.Attributes[0])
}
