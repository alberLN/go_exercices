package main

import "fmt"

//vars used in vars fuction
var y = 42
var z = "This is a String"

//custom type
type test int

var typeTest test = 99

func main() {
	fmt.Println("Hello world")
	foo()
	loop()
	vars()
}

func foo() {
	fmt.Println("I'm in foo")
}

func loop() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func vars() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(typeTest)
	fmt.Printf("%T\n", typeTest)
}
