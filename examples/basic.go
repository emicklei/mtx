package main

import (
	"fmt"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/basic"
)

func main() {
	p := basic.NewPackage("mypackage")
	c := p.Entity("UserStory")
	c.A("Priority", basic.Integer, "the importance of the story")
	fmt.Println(mtx.ToSource(p))
}
