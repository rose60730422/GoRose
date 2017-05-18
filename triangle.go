package main

import (
	"fmt"
)

func printStars(number int) {
	var line = ""
	for i :=0 ; i < number ; i++{
		line += "*"
	}
	fmt.Println(line)
}

func triangleOne(){
	for i := 1; i < 6; i++ {
		printStars(i)
	}
}

func triangleTwo() {
	for i := 1; i < 6; i++ {
		printInt(i) // print 1234...i
	}
}

func printInt(number int) {
	for i := 1; i < number + 1; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println("")
}

func triangleThree() {
	for i := 1; i < 6; i++ {
		printInt(i)
	}
	for i := 0; i < 4; i++ {
		printInt(4 - i)
	}
}

func main() {
}