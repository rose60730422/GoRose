package main

import (
	"fmt"
)

func main() {
	var slice = make([][]int, 10)
	var x int
	for x = 0; x < 10; x++ {
		slice[x] = make([]int, 10)
	}

	// TODO
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			slice[x][y] = 1
		}
	}

	// for y := 0; y < 10; y++ {
	// 	slice[5][y] = 1
	// }
	// for x = 0; x < 10; x++ {
	// 	slice[x][5] = 1
	// }
	// fmt.Println(slice)
	for x = 0; x < 10; x++ {
		fmt.Println(slice[x])
	}
}
