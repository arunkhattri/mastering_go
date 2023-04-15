package main

import "fmt"

func main() {
	// create an empty slice
	aSlice := []float64{}
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// add elements to aSlice
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice))

	// A slice with length 4
	s := make([]int, 4)
	s[0] = -1
	s[1] = -2
	s[2] = -3
	s[3] = -4

	fmt.Printf("%v\n", s)

	s = append(s, -5)
	fmt.Println(s)

	// A 2D slice
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}

	// visiting all elements of a 2D slice
	for _, i := range twoD {
		for _, j := range i {
			fmt.Print(j, " ")
		}
	}

}
