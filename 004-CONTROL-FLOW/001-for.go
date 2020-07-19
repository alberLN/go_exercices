package main

import "fmt"

func main() {
	basicFor()
	forCondition()
	forCrud()
}

// for val; condition; sum {}
func basicFor() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

//for condition {}
func forCondition() {
	i := 1
	for i <= 29 {
		fmt.Println(i)
		i++
	}
}

//for {}
func forCrud() {
	i := 1
	for {
		if i >= 29 {
			break
		}
		i++
	}
}
