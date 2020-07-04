package main

import (
	"fmt"
)

var  z int

func main() {
	x := 43
	y := 26

	fmt.Println(x + y)
	fmt.Println("Hellow, playground", 26, false)

	foo()
}
func foo() {
	         fmt.Println( "it's zero", z)
}
