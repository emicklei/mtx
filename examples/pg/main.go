package main

import (
	"fmt"

	"github.com/emicklei/mtx"
	"github.com/emicklei/mtx/db"
	"github.com/emicklei/mtx/golang"
	"github.com/emicklei/mtx/pg"
)

func main() {
	mydb := pg.NewDatabase("mydb")
	tab := mydb.Table("mytable")
	tab.Column("mytext").Type(pg.Text)

	ent := db.ToEntity(tab)
	str := golang.ToStruct(ent)
	src := mtx.ToSource(str)
	fmt.Println(src)
}
