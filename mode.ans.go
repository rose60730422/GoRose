package main

import "fmt"
import "strings"
import "strconv"

func countMode(myMap map[int]int) (s []int, count int) {
	var max int
	s = make([]int, 0)

	// 選出最大數
	for _, value := range myMap {
		if max < value {
			max = value
		}
	}

	// 選出出現次數為最大數的數字
	for key, value := range myMap {
		if value == max {
			s = append(s, key)
		}
	}

	return s, max
}


func myMode(input string) {

	s_split := strings.Fields(input)

	box := make(map[int]int) // 用來存放出現的數字

	var modeNumber []int
	var modeMax  int
	
	for _, v := range(s_split) {
		var num int
		// 把輸入的數字從字串轉成int
		if x, err := strconv.Atoi(v); err == nil {
			num = x
		}

		// 數字 num 的出現次數 ++
		box[num]++

		// 算出眾數以及眾數出現的次數
		modeNumber, modeMax = countMode(box)


		// 組出答案字串
		var ansStr string
		for _, v := range modeNumber {
			ansStr += fmt.Sprintf("%d ", v)
		}
		ansStr += fmt.Sprintf("(%d)\n", modeMax)
		fmt.Print(ansStr)
	}
	return
}

func main() {
	input := "8\n14\n4\n14\n3\n8\n8\n3\n14\n3\n"

	myMode(input)
}