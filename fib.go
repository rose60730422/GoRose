package main

import (
	"fmt"
)


	var last2 = 0
	var last = 1


func fib() {

	fmt.Println(last2)
	// update last2, last
	previous_last2 := last2
	previous_last := last
	last2 = previous_last
	last = previous_last + previous_last2
}

func fibHard() {

	fmt.Println(last2)
	// update last2, last
	last = last2 + last
	last2 = last - last2
}

var a0 = 1

func sn() {
	fmt.Println(a0)
	a0 = a0 * 2
}

func main() {
	for i := 0; i < 10; i++{
		fibHard()
	}
}

