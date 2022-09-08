package mtx

type FieldMapping struct {
	Type TypeMapping
}

type EntityMapping struct {
	Fields []FieldMapping
}

type TypeMapping struct{}
