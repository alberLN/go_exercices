package main

import "fmt"

const (
	a     = 1
	b int = 2
)

func main() {
	fmt.Printf("Const A with value %d is untyped\n", a)
	fmt.Printf("Const B with value %d is typed\n", b)
}
