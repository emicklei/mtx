package basic

import (
	"fmt"
	"strings"

	"github.com/emicklei/mtx"
)

type EntityRef struct {
	PackageName string `json:"package"`
	EntityName  string `json:"name"`
}

func (p *Package) OneToMany(one, many *Entity) *Relation {
	rel := &Relation{
		Named:    mtx.N("mtx.Relation", fmt.Sprintf("%s_link_%s", strings.ToLower(one.Name), strings.ToLower(many.Name))),
		pkg:      p,
		LeftRef:  EntityRef{PackageName: p.Name, EntityName: one.Name},
		RightRef: EntityRef{PackageName: p.Name, EntityName: many.Name},
	}
	p.Relations = append(p.Relations, rel)
	return rel
}

func (p *Package) ManyToMany(left, right *Entity) *Relation {
	rel := &Relation{
		Named: mtx.N("mtx.Relation", fmt.Sprintf("%s_link_%s", strings.ToLower(left.Name), strings.ToLower(right.Name))),
		pkg:   p,
		// assume all in same package
		LeftRef:  EntityRef{PackageName: p.Name, EntityName: left.Name},
		RightRef: EntityRef{PackageName: p.Name, EntityName: right.Name},
	}
	p.Relations = append(p.Relations, rel)
	return rel
}

type Relation struct {
	mtx.Named
	pkg              *Package
	LeftRef          EntityRef    `json:"left"`
	LeftRole         string       `json:"left_role_name"`
	LeftCardinality  string       `json:"left_cardinality"`
	RightRef         EntityRef    `json:"right"`
	RightRole        string       `json:"right_role_name"`
	RightCardinality string       `json:"right_cardinality"`
	Attributes       []*Attribute `json:"attributes"`
}

func (r *Relation) Left(leftName string) *Relation {
	r.LeftRole = leftName
	return r
}

func (r *Relation) Right(rightName string) *Relation {
	r.RightRole = rightName
	return r
}

func (r *Relation) A(name string, typ mtx.Datatype, doc string) *Attribute {
	return r.Attribute(name).Type(typ).Doc(doc)
}

func (r *Relation) Attribute(name string) *Attribute {
	attr, ok := mtx.FindByName(r.Attributes, name)
	if ok {
		return attr
	}
	attr = &Attribute{}
	attr.Named = mtx.N(EntityAttributeClass, name)
	r.Attributes = append(r.Attributes, attr)
	return attr
}
