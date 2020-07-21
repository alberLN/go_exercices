package main

import "fmt"

type person struct {
	name     string
	lastName string
	favorite []string
}

func main() {
	fmt.Println("Basic Struct usage\n")
	basicStruct()
	fmt.Println("\nStruct with map\n")
	mapStruct()
	fmt.Println("\nAnonymous struct\n")
	anonymousStruct()
}

func basicStruct() {
	p1 := person{
		"Penelope",
		"Waterson",
		[]string{"Chocolat", "Strawberry", "blackberry"},
	}

	p2 := person{
		"Harry",
		"Strange",
		[]string{"Fish", "Car oil", "Glue"},
	}

	fmt.Printf("I am %s %s and I like: \n", p1.name, p1.lastName)
	for i, v := range p1.favorite {
		println(i, v)
	}
	fmt.Println("\n")
	fmt.Printf("I am %s %s and I like: \n", p2.name, p2.lastName)
	for i, v := range p2.favorite {
		println(i, v)
	}
}

func mapStruct() {
	myMap := map[string]person{
		"Waterson": {
			"Penelope",
			"Waterson",
			[]string{"Chocolat", "Strawberry", "blackberry"},
		},
		"Strange": {
			"Harry",
			"Strange",
			[]string{"Fish", "Car oil", "Glue"},
		},
	}

	for pi, p := range myMap {
		fmt.Printf("About %s: I am %s %s and I like: \n", pi, p.name, p.lastName)
		for fi, f := range p.favorite {
			println(fi, f)
		}
	}
}

func anonymousStruct() {
	p := struct {
		name     string
		lastName string
		favorite []string
	}{
		"Julia",
		"Adams",
		[]string{"Chocolat", "Strawberry", "blackberry"},
	}

	fmt.Println(p)
}
