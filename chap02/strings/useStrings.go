package main

import (
	"fmt"
	"strings"
)

func main() {
	uName := strings.ToUpper("Arun Kr Khattri")
	fmt.Printf("|- To upper: %s\n", uName)

	fmt.Printf("|- %s\n", strings.Title("This will be a title!"))

	// compare two strings ignoring case
	fmt.Printf("|- EqualFold: %v\n", strings.EqualFold("Arun", "aRUN"))
	fmt.Printf("|- EqualFold: %v\n", strings.EqualFold("Arun", "ArunK"))

	// check given string begins with string
	fmt.Printf("|- Prefix: %v\n", strings.HasPrefix("Arun", "Ar"))
	// its case sensitive
	fmt.Printf("|- Prefix: %v\n", strings.HasPrefix("Arun", "ar"))
	// check given string ends with another given string
	fmt.Printf("|- Suffix: %v\n", strings.HasSuffix("Arun", "un"))
	// its also case sensitive
	fmt.Printf("|- Suffix: %v\n", strings.HasSuffix("Arun", "UN"))

	// split the strings with white space character
	// defined by unicode.IsSpace()
	t := strings.Fields("Arun Kumar Khattri")
	fmt.Printf("|- Fields: %v\n", len(t))

	t1 := strings.Fields("ThisIs a\tstring!")
	fmt.Printf("|- Fields: %v\n", len(t1))

	for i := 0; i < len(t1); i++ {
		fmt.Printf("   %d-  %v\n", i+1, t1[i])
	}

	// index of second string if found in first string
	fmt.Printf("|- Index: %v\n", strings.Index("Khattri", "ttr"))

	// split the string with given separator
	fmt.Printf("|- %s\v", strings.Split("arun kr khattri", ""))

}
