package main

import (
	"fmt"
	"time"
)

func main() {
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ := time.Parse(longForm, "Jul 5, 1971 at 5:30am (IST)")
	fmt.Println(t)

	// shortform, no timezone
	const shortForm = "2006-Jan-02"
	d, _ := time.Parse(shortForm, "1971-Jul-05")
	fmt.Println(d)

	dt, _ := time.Parse(time.RFC3339, "1971-07-05T05:30:00+05:30")
	fmt.Println(dt)
}
