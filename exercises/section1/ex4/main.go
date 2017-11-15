package main

import "fmt"

type mytype int

var x mytype

func main() {
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)
	x = 42
	fmt.Printf("%v\n", x)
}
