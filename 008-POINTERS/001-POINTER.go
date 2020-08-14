package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{name: "David Jow"}
	fmt.Println(p.name)
	modify(&p)
	fmt.Println(p.name)
}

func modify(p *person) {
	p.name = "Jowi not david"
}
