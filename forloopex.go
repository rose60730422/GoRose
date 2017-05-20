
package main

import "fmt"

func factorial(x int) int {
	var sum int
	for i := 0 ; i < x + 1 ; i++ {
		if i == 0 {
			sum = 1
		} else {
			sum = sum * i
		}
	}
	return sum
}

func recursiveFactorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * recursiveFactorial(x - 1)
}


func main(){
	fmt.Println(factorial(5))
}