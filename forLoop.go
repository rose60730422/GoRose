package main

import (
	"fmt"
)

func main() {
	sum := 10
	for i := 1; i < 11; i = i + 1 {
		sum = sum + i
	}
	fmt.Println(sum)

	fmt.Println(foo(8))
}

func foo(x int) int {
	var ans = 1
	for i := 1; i < x+1; i = i + 1 {
		ans = ans * i
	}
	return ans
}
