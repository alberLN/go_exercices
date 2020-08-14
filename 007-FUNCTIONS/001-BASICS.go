package main

import "fmt"

func main() {
	fmt.Println("Start referred function")
	defer deferred()
	fmt.Println("Baisc funcs")
	a := foo()
	b, c := boo()
	fmt.Println(a, b, c)
	fmt.Println("Variadic func")
	d := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Println(d)
	fmt.Println("Anonymous function")
	func() {
		fmt.Println("I'm an anonymous function")
	}()
	fmt.Println("Assigning functions")
	assignFuction()
}

/*Basic function that return one result*/
func foo() int {
	return 50
}

/*Basic function that return two results*/
func boo() (int, string) {
	return 50, "I love functions"
}

/*Variadic functions*/
func sum(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

/*defer example*/
func deferred() {
	fmt.Println("Referred function executed")
}

/*assignig function*/
func assignFuction() {
	af := func() {
		fmt.Println("I'am an assigned function")
	}

	af()
}
