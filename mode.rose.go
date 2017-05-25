package main

import "fmt"
import "strings"
import "strconv"

func countMode(myMap map[int]int) (s []int, count int) {
	var max int
	s = make([]int,0)
	for _, value := range(myMap){
		if max < value{
			max = value
		}  
	}
	for key, value := range(myMap){
		if value == max {
			s = append(s ,key )
		}
	}
	return s, max
}

func myMode(input string) {
	var s_split = strings.Fields(input)
	var box = make(map[int]int)
	for _, v := range(s_split){
		var num int
		if x, err := strconv.Atoi(v) ; err == nil {
			num = x
		} 
		box[num]++
		
		// 把眾數找出來
		mode, times := countMode(box)
		answer := format(mode, times)
		fmt.Println(answer)
	}
}

func format(modeSlice []int, times int) string {
	var str string 
	fmt.Println(modeSlice)

	// 把Slice中的元素回傳成字串
	for _, v := range modeSlice {
		str += fmt.Sprintf("%d ", v)
		fmt.Println(v)
	}

	str += fmt.Sprintf("(%d)", times)
	return str
}

func main() {
	input := "8\n14\n4\n14\n3\n8\n8\n3\n14\n3\n"

	myMode(input)
}