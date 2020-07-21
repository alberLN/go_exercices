package main

import "fmt"

func main() {
	println("Create Map\n_________________________\n")
	create()
	println("\n\nAdd element\n_________________________\n")
	add()
	println("\n\nRemove element\n_________________________\n")
	remove()
}

func create() {
	myMap := map[string][]string{"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`}}

	for i, v := range myMap {
		fmt.Printf("\n\n%s :\n", i)
		for j, w := range v {
			fmt.Printf("position %d value %s \n", j, w)
		}
	}
}

func add() {
	myMap := map[string][]string{"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`}}

	myMap["new_element"] = []string{`I am`, `a new`, `element`}

	for i, v := range myMap {
		fmt.Printf("\n%s :\n", i)
		for j, w := range v {
			fmt.Printf("position %d value %s \n", j, w)
		}
	}
}

func remove() {
	myMap := map[string][]string{"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`}}

	delete(myMap, "no_dr")

	for i, v := range myMap {
		fmt.Printf("\n%s :\n", i)
		for j, w := range v {
			fmt.Printf("position %d value %s \n", j, w)
		}
	}
}
