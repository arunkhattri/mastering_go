package main

import "fmt"

func main() {
	r := '₹'
	fmt.Println("As an int32 value:", r)
	// convert runes to text
	fmt.Printf("As a string: %x and as a character: %c\n", r, r)

	aStr := "hello world! ₹"
	fmt.Println("First character:", string(aStr[0]))

	// print an existing string as runes
	for _, v := range aStr {
		fmt.Printf("%x ", v)
	}
	fmt.Println()
}
