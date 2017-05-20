package main

import "fmt"

func roseCount(input []int) string {
	var max int
	var index int
	max = input[0]
	index = 0
	for i := 1; i < len(input) ; i++{
		if input[i] > max {
			max = input[i]
			index = i 
		}
	}
	return fmt.Sprintf("%d, %d\n", max , index)
}

func danielCount(input []int) string {
	for i, v := range(input) {
		var i2 int
		var v2 int
		for i2, v2 = range(input) {
			if i != i2 {
				if v < v2 {
					i2 = len(input)
					break
				}
			}
		}
		if (i2 == len(input) - 1) {
			return fmt.Sprintf("%d, %d\n", v, i)
		}
	}
	return ""
}

func main() {
	input := []int{10, 2, 3, 0, 5, 6, 7, 999}
	correctAns := "999, 7\n"

	roseAns := danielCount(input)
	fmt.Println(roseAns)
	if (roseAns == correctAns) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}

}
