package mtx

import "testing"

func TestOneToManyRelation(t *testing.T) {
	p := NewPackage("persons")
	e := p.Entity("Person")
	r := p.OneToMany(e, e)
	r.Left("parent")
	r.Right("children")
}

func TestManyToManyRelation(t *testing.T) {
	p := NewPackage("persons")
	e := p.Entity("Person")
	r := p.ManyToMany(e, e)
	r.Name = "followers"
}
