package main

import (
	"fmt"
)

func main() {
	fmt.Println("When's Saturday?")
	today := 5
	switch today {
	case 6:
		fmt.Println("Today.")
	case 6 - 1:
		fmt.Println("Tomorrow.")
	case 6 - 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
