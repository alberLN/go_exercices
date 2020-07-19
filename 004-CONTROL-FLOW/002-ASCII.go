package main

import "fmt"

func main() {
	for i := 32; i < 127; i++ {
		fmt.Printf("%#U\n", i)
	}
}
