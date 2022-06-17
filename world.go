package mtx

import (
	"github.com/emicklei/mtx/namespace"
	"github.com/emicklei/mtx/proto"
)

type World struct {
	namespaces    map[string]*namespace.Namespace
	protoPackages map[string]*proto.Package
}

func NewWorld() *World {
	return &World{
		namespaces:    map[string]*namespace.Namespace{},
		protoPackages: map[string]*proto.Package{},
	}
}

func (w *World) Namespace(name string) *namespace.Namespace {
	n, ok := w.namespaces[name]
	if ok {
		return n
	}
	n = namespace.New(name)
	w.namespaces[name] = n
	return n
}

func (w *World) ProtoPackage(name string) *proto.Package {
	p, ok := w.protoPackages[name]
	if ok {
		return p
	}
	p = proto.NewPackage(name)
	w.protoPackages[name] = p
	return p
}
