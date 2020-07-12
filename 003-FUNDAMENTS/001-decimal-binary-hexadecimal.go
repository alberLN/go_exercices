package main

import "fmt"

func main() {
	a := 5
	/*
		%d -> deciaml
		%b -> binary
		%#x -> hexadecimal
	*/
	fmt.Printf("%d as decimal %d\n", a, a)
	fmt.Printf("%d as binary %b\n", a, a)
	fmt.Printf("%d as hexadecimal %#x\n", a, a)
}
