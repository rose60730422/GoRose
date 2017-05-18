package main

import (
	"fmt"
)
func printStars(number int){
	for i := 0 ; i < number ; i++{
		fmt.Printf("%s", "*")
	}
	fmt.Printf("\n")
}
func triangleOne(number int){
	for i := 1 ; i < number + 1 ; i++{
		printStars(i)
	}
}
func main() {
	//printStars(5)
	triangleOne(10)
}