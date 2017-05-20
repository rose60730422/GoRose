package main

import "fmt"
import "strings"
import "strconv"

type Question struct {
	number int
}

func countFive(number int) int {
	i := 0
	for number % 5 == 0 {
		number /= 5
		i++
	}
	return i
}

func roseAns(testData string) string {
	myAns := ""
	testDataSlice := parseTestData(testData)
	for _, question := range(testDataSlice) {
		ans := countFactorial(question)
		myAns += ans
	}
	return myAns
}

func parseTestData (s string) []int {
	var input = strings.Fields(s)
	var resultSlice []int
	for _ , Question := range(input) {
		var x int		
		if number , err := strconv.Atoi(Question) ; err == nil {
			x = number 
		}
		resultSlice = append(resultSlice , x)
	}	
	return resultSlice
}

func countFactorial(queNum int) string {
	var sum int 
	for  i:= 1 ; i < queNum + 1 ; i++ {
		sum = sum + countFive(i)
	}
	// sum += countFive(1)
	// sum += countFive(2)
	// sum += countFive(3)
	// sum += countFive(4)
	// ...

	return fmt.Sprintf("%d\n" , sum)
}



func main() {
	// fmt.Println(countFive(1))
	// fmt.Println(countFive(5))
	// fmt.Println(countFive(25))
	// fmt.Println(countFive(26))
	// fmt.Println(countFive(625))

	
	// fmt.Println(recursiveFactorial(5))

	// testData := "3\n8\n12\n25\n50\n80\n100\n200\n500\n1000\n"
	// correctAns := "0\n1\n2\n6\n12\n19\n24\n49\n124\n249\n"

	// myAns := roseAns(testData)
	// fmt.Printf("myAns:\n%s", myAns)
	// if myAns == correctAns {
	// 	fmt.Println("PASS")
	// } else {
	// 	fmt.Println("FAIL")
	// }
}