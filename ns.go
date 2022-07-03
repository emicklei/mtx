package mtx

type Namespace struct {
	Name     string
	elements map[string]Named
}

func NewNamespace(name string) *Namespace {
	return &Namespace{
		Name:     name,
		elements: map[string]Named{},
	}
}
