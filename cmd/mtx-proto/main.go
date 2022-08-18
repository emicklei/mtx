package main

import (
	"flag"
	"os"

	protom "github.com/emicklei/mtx/proto"
	"github.com/emicklei/proto"
)

var oFile = flag.String("f", "", "proto file to read")

func main() {
	flag.Parse()
	reader, _ := os.Open(*oFile)
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, _ := parser.Parse()

	v := new(visitor)
	proto.Walk(definition,
		proto.WithPackage(v.handlePackage),
		proto.WithMessage(v.handleMessage),
		proto.WithNormalField(v.handleNormalField))

	v.pkg.SourceOn(os.Stdout)
}

type visitor struct {
	proto.NoopVisitor
	pkg *protom.Package
	msg *protom.Message
}

func (v *visitor) handlePackage(p *proto.Package) {
	v.pkg = protom.NewPackage(p.Name)
}

func (v *visitor) handleMessage(m *proto.Message) {
	v.msg = v.pkg.Message(m.Name)
}

func (v *visitor) handleNormalField(m *proto.NormalField) {
	v.msg.F(m.Name, m.Sequence, protom.Type(m.Type), m.Comment.Message())
}
