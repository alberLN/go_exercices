package main

import "fmt"

const (
	year1 = 2020 + iota
	year2 = 2020 + iota
	year3 = 2020 + iota
	year4 = 2020 + iota
)

func main() {
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)
}
