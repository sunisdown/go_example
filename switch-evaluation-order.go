package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Print("When is Saturday\n")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Print("Today")
	case today + 1:
		fmt.Print("Tomorrow")
	case today + 2:
		fmt.Print("In two days")
	default:
		fmt.Print("Too far away")
	}
}
