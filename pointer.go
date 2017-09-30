package main

import "fmt"

func foo(number *int) {
	*number += 10
	fmt.Println(*number)
}

func main() {
	var x int
	x = 10
	foo(&x)
	fmt.Println(x)
}