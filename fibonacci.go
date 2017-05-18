package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var last2 = 0
	var last = 1
	return func ()int {
		previous_last := last
		previous_last2 := last2
		last2 = previous_last
		last = previous_last + previous_last2
		return previous_last2
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}