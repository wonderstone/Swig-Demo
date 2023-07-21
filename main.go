package main

import (
	"fmt"

	"github.com/wonderstone/Swig-demo/lib"
	solib "github.com/wonderstone/Swig-demo/so"
)

func main() {
	fmt.Println("Hello World!")
	res := lib.Minus(1, 2)
	res2 := solib.Add(2, 2)
	fmt.Println(res)
	fmt.Println(res2)
}
