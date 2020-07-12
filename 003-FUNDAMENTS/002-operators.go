package main

import "fmt"

func main() {
	a := 1
	b := 2

	fmt.Println("Operator: ==")
	fmt.Printf("%d == %d ? %t\n", a, b, (a == b))
	fmt.Printf("%d == %d ? %t\n", a, a, (a == a))
	fmt.Println("\nOperator: <=")
	fmt.Printf("%d <= %d ? %t\n", a, b, (a <= b))
	fmt.Printf("%d <= %d ? %t\n", b, a, (b <= a))
	fmt.Printf("%d <= %d ? %t\n", a, a, (a <= a))
	fmt.Println("\nOperator: >=")
	fmt.Printf("%d >= %d ? %t\n", a, b, (a >= b))
	fmt.Printf("%d >= %d ? %t\n", b, a, (b >= a))
	fmt.Printf("%d >= %d ? %t\n", a, a, (a >= a))
	fmt.Println("\nOperator: !=")
	fmt.Printf("%d != %d ? %t\n", a, b, (a != b))
	fmt.Printf("%d != %d ? %t\n", a, a, (a != a))
	fmt.Println("\nOperator: <")
	fmt.Printf("%d < %d ? %t\n", a, b, (a < b))
	fmt.Printf("%d < %d ? %t\n", b, a, (b < a))
	fmt.Printf("%d < %d ? %t\n", a, a, (a < a))
	fmt.Println("\nOperator: >")
	fmt.Printf("%d > %d ? %t\n", a, b, (a > b))
	fmt.Printf("%d > %d ? %t\n", b, a, (b > a))
	fmt.Printf("%d > %d ? %t\n", a, a, (a > a))
}
