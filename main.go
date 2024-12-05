package main

import (
	"fmt"

	"github.com/AVM-tiket/AVM-LAB-GOMOD/a"
	"github.com/AVM-tiket/AVM-LAB-GOMOD/b"
	"github.com/AVM-tiket/AVM-LAB-GOMOD/pkg/c"
	"github.com/AVM-tiket/AVM-LAB-GOMOD/pkg/d"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(a.Echo("World"))
	fmt.Println(b.Echo("World"))
	fmt.Println(c.Echo("World"))
	fmt.Println(d.Echo("World"))
}
