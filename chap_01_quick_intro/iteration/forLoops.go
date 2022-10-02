package main

import "fmt"

func main() {
	// traditional
	fmt.Println("Traditional...")
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// idiomatic go
	fmt.Println("Idiomatic go...")
	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// for loop used as while
	fmt.Println("for loop used as while...")
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// This is a slice but range also works with arrays
	fmt.Println("slice ...")
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index: ", i, "value: ", v)
	}
}
