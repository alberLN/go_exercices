package main

import "fmt"

func main() {
	fmt.Println("If else: ")
	if_else(true)
	if_else(false)
	fmt.Println("\nSwitch case: ")
	switchCase(1)
	switchCase(2)
	switchCase(3)
}

func if_else(b bool) {
	if b {
		fmt.Println("I am true")
	} else {
		fmt.Println("I am false")
	}
}

func switchCase(c int) {
	switch c {
	case 1:
		fmt.Println(c)
	case 2:
		fmt.Println(c)
	default:
		fmt.Println("Default")
	}
}
