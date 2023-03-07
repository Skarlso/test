package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type S struct {
	Name string
}

func main() {
	fmt.Println("vim-go")
	spew.Dump(&S{Name: "bla"})
}
