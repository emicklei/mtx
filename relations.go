package mtx

import (
	"fmt"
	"strings"
)

type EntityRef struct {
	PackageName string `json:"package"`
	EntityName  string `json:"name"`
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

func (p *Package) ManyToMany(left, right *Entity) *ManyToMany {
	rel := &ManyToMany{
		Named:    N("mtx.ManyToMany", fmt.Sprintf("%s_link_%s", strings.ToLower(left.Name), strings.ToLower(right.Name))),
		pkg:      p,
		LeftRef:  EntityRef{PackageName: p.Name, EntityName: left.Name},
		RightRef: EntityRef{PackageName: p.Name, EntityName: right.Name},
	}
	p.Relations = append(p.Relations, rel)
	return rel
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

type ManyToMany struct {
	*Named
	pkg        *Package
	LeftRef    EntityRef    `json:"left"`
	RightRef   EntityRef    `json:"right"`
	Attributes []*Attribute `json:"attributes"`
}

func (r *ManyToMany) A(name string, typ AttributeType, doc string) *Attribute {
	return r.Attribute(name).Type(typ).Doc(doc)
}

func (r *ManyToMany) Attribute(name string) *Attribute {
	attr, ok := FindByName(r.Attributes, name)
	if ok {
		return attr
	}
	attr = &Attribute{
		IsRequired: true, // required by default
	}
	attr.Named = N(EntityAttributeClass, name)
	r.Attributes = append(r.Attributes, attr)
	return attr
}
