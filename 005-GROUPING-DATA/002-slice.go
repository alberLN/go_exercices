package main

import "fmt"

func main() {
	basicSlice()
	slicingSlice()
	appendSlice()
	deleteSlice()
	makeSlice()
	dimensionalSlice()
}

func basicSlice() {
	fmt.Println("\nBasic Slice")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for index, value := range slice {
		fmt.Println(index, value)
	}
	fmt.Printf("%T\n", slice)
}

func slicingSlice() {
	fmt.Println("\nSlicing Slice")
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(slice[:5])
	fmt.Println(slice[:4])
	fmt.Println(slice[5:])
	fmt.Println(slice[4:5])
	fmt.Println(slice[1:3])
	fmt.Println(slice[0:1])
}

func appendSlice() {
	fmt.Println("\nAppend Slice")
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice = append(slice, 52)
	fmt.Println(slice)
	slice = append(slice, 53, 54, 55)
	fmt.Println(slice)
	newSlice := []int{56, 57, 58, 59, 60}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}

func deleteSlice() {
	fmt.Println("\nDelete Slice")
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func makeSlice() {
	fmt.Println("\nMake Slice")
	x := make([]string, 50, 500)
	fmt.Println(len(x)) //tells you the capacity of the underlying array
	fmt.Println(cap(x)) //ells you how many items are in the array

	//Initialize 50 values
	for i := 0; i < 50; i++ {
		x[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//Append states
	x = append(x, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func dimensionalSlice() {
	fmt.Println("\nDimensional Slice")
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{a, b}
	fmt.Println(x)
	for index, value := range x {
		fmt.Println("record: ", index)
		for j, val := range value {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
