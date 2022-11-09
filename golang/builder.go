package golang

import (
	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
	"github.com/iancoleman/strcase"
)

type StructBuilder struct {
	entity         *basic.Entity
	typeMapper     TypeMapper
	result         *Struct
	fieldTaggers   []FieldTagger
	methodBuilders []MethodBuilder
}

func NewStructBuilder(e *basic.Entity) *StructBuilder {
	return &StructBuilder{
		entity:     e,
		typeMapper: StandardTypeMapper,
		result:     new(Struct)}
}

func (b *StructBuilder) WithTypeMapper(m TypeMapper) *StructBuilder {
	b.typeMapper = m
	return b
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
			// use the mapper
			fieldType = b.typeMapper(each.AttributeType, each.IsNullable)
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

func WithTypeMapper(tm TypeMapper) Option {
	return func(b *StructBuilder) *StructBuilder {
		return b.WithTypeMapper(tm)
	}
}

// ToStruct builds a Struct. Option control build strategies.
func ToStruct(ent *basic.Entity, options ...Option) *Struct {
	b := NewStructBuilder(ent)
	for _, each := range options {
		b = each(b)
	}
	return b.Build()
}
