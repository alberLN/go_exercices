package main

import "fmt"

type vehicle struct {
	doors  int
	colors []string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	fmt.Println("test")
	v1 := truck{
		vehicle{
			5,
			[]string{"red", "black"},
		},
		true,
	}

	v2 := sedan{
		vehicle{
			3,
			[]string{"white", "silver"},
		},
		true,
	}

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v1.doors)
	fmt.Println(v2.doors)
}
