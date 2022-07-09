package mtx

import (
	"fmt"
)

type Package struct {
	*Named
	Entities  []*Entity
	Relations []any
}

func NewPackage(s string) *Package {
	return &Package{
		Named: N("mtx.Package", s),
	}
}

func (p *Package) Entity(name string) *Entity {
	for _, each := range p.Entities {
		if each.Name == name {
			return each
		}
	}
	e := new(Entity)
	e.pkg = p
	e.Named = N("mtx.Entity", name)
	p.Entities = append(p.Entities, e)
	return e
}

func (p *Package) Relation(rel any) any {
	p.Relations = append(p.Relations, rel)
	return rel
}

func (p *Package) Fullname(e *Entity) string {
	return fmt.Sprintf("%s.%s", p.Name, e.Name)
}
