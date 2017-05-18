package main

import "fmt"
import "math"
import "strings"
import "strconv"

func split(testData string) []float64 {
	var resultSlice []float64
	var testDataSlice = strings.Fields(testData)
	for _, distance := range(testDataSlice) {
		var x float64
		if num, err := strconv.ParseFloat(distance, 32); err == nil {
			x = num
		}
		resultSlice = append(resultSlice, x)
	}
	return resultSlice
}

func taxiCount(distance float64) string {
	var dollars int
	if distance < 1.5 {
		dollars = 75
	} else {
		var y = int(math.Ceil((distance - 1.5) * 4))
		dollars = 75 +  y * 5
	}
	return fmt.Sprintf("%d\n" , dollars)
}

func roseAns(testData string) string {
	myAns := ""
	problemSet := split(testData) // 把題目拆成一個 slice, 每個元素都是一個題目(int)
	for _, distance := range(problemSet) {
		ans := taxiCount(distance)
		myAns += ans
	}
	return myAns
}

func main() {
	testData := "0.9\n1.6\n2.2\n3.5\n4.8\n"
	correctAns := "75\n80\n90\n115\n145\n"

	myAns := roseAns(testData)
	fmt.Printf("myAns:\n%s", myAns)
	if myAns == correctAns {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}