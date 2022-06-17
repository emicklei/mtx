package namespace

import "github.com/emicklei/mtx/model"

type Namespace struct {
	Name string
}

func New(name string) *Namespace { return &Namespace{Name: name} }

// TODO return Builder for Model, not Model itself
func (n *Namespace) Model(name string) *model.Model { return model.New(name) }
