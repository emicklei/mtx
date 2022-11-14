package golang

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
	"github.com/iancoleman/strcase"
)

type StructBuilder struct {
	entity         *basic.Entity
	result         *Struct
	fieldTaggers   []FieldTagger
	methodBuilders []MethodBuilder
}

func NewStructBuilder(e *basic.Entity) *StructBuilder {
	return &StructBuilder{
		entity: e,
		result: new(Struct)}
}

func (b *StructBuilder) Build() *Struct {
	// set name
	n := strcase.ToCamel(b.entity.Name)
	if v, ok := b.entity.Get(GoTypeName); ok {
		n = v.(string)
	}
	b.result.Named = mtx.N("golang.Struct", n)
	// set doc
	b.result.Documentation = b.entity.Documentation
	// set fields
	for _, each := range b.entity.Attributes {
		var fieldType mtx.Datatype
		// see if type is overridden
		if n, ok := each.Get(GoTypeName); ok {
			fieldType = mtx.NewBasicType(n.(string))
		} else {
			fieldType = FromBasicType(each.AttributeType)
		}
		f := &Field{
			Named:     mtx.N("golang.Field", b.goFieldName(each)),
			FieldType: fieldType,
		}
		f.CopyPropertiesFrom(each.Named)
		// add tags
		for _, tagger := range b.fieldTaggers {
			tagger(each, f)
		}
		f.Documentation = each.Documentation
		b.result.Fields = append(b.result.Fields, f)
	}
	for _, each := range b.methodBuilders {
		each(b.result)
	}
	return b.result
}

// TODO create a FieldNamer interface/func
func (b *StructBuilder) goFieldName(a *basic.Attribute) string {
	// TODO check override
	return strcase.ToCamel(a.Name)
}

type Option func(b *StructBuilder) *StructBuilder
