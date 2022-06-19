package bq

import "github.com/emicklei/mtx/core"

type bqSpace core.Namespace

func NewNamespace(name string) *bqSpace {
	return (*bqSpace)(core.NewNamespace(name))
}

func (s bqSpace) Dataset(n string) *Dataset {
	return &Dataset{Named: core.N("bq.Dataset", n)}
}
