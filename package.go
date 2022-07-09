package mtx

import (
	"fmt"
	"strings"
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

func (p *Package) OneToMany(one, many *Entity) *OneToMany {
	rel := &OneToMany{
		Named:   N("mtx.OneToMany", fmt.Sprintf("%s_to_%s", strings.ToLower(one.Name), strings.ToLower(many.Name))),
		pkg:     p,
		OneRef:  EntityRef{PackageName: p.Name, EntityName: one.Name},
		ManyRef: EntityRef{PackageName: p.Name, EntityName: many.Name},
	}
	p.Relations = append(p.Relations, rel)
	return rel
}

func (p *Package) Fullname(e *Entity) string {
	return fmt.Sprintf("%s.%s", p.Name, e.Name)
}

type EntityRef struct {
	PackageName string `json:"package"`
	EntityName  string `json:"name"`
}

type OneToMany struct {
	*Named
	pkg      *Package
	OneRef   EntityRef `json:"one"`
	OneName  string    `json:"one_name"`
	ManyRef  EntityRef `json:"many"`
	ManyName string    `json:"many_name"`
}

func (r *OneToMany) One(name string) *OneToMany {
	r.OneName = name
	return r
}

func (r *OneToMany) Many(name string) *OneToMany {
	r.ManyName = name
	return r
}
