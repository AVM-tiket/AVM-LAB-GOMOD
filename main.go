package main

import (
	"fmt"

	"github.com/AVM-tiket/AVM-GO-COMMON/a"
	"github.com/AVM-tiket/AVM-GO-COMMON/b"
	"github.com/AVM-tiket/AVM-GO-COMMON/pkg/c"
	"github.com/AVM-tiket/AVM-GO-COMMON/pkg/d"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(a.Echo("World"))
	fmt.Println(b.Echo("World"))
	fmt.Println(c.Echo("World"))
	fmt.Println(d.Echo("World"))
}
