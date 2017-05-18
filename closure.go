package main

import "fmt"


var x int
func foo(y int)int {

	x += y
	return x
}

func main() {
	fmt.Println(foo(5))
	fmt.Println(foo(5))
}