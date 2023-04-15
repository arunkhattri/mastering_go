package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	const shortForm = "02 January 2006"

	start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("usage: dates parse_string")
		return
	}

	dateString := os.Args[1]

	// is this a date only?
	// d, err := time.Parse(shortForm, dateString)

	// if err == nil {
	// 	fmt.Println("Full:", d)
	// 	fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	// }
	// is this a date and time?
	d, err := time.Parse(time.RFC3339, dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// convert epoch time to time.Time
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
	duration := time.Since(start)
	fmt.Println("Execution time:", duration)
}
