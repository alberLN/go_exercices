package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	foo()
	loop()
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
