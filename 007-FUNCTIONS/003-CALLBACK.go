package main

import "fmt"

func main() {
	/*Returned func*/
	rf := returnfunc()
	rf("I'm a returned func")
	callback("I'm a call back", rf)
}

func returnfunc() func(text string) {
	return func(text string) {
		fmt.Println(text)
	}
}

func callback(text string, cb func(string)) {
	cb(text)
}
