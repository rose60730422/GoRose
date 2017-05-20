package main

import (
	"strings"
	"fmt"
	"reflect"
)

func roseWordCount(s string) map[string]int {
	s_split := strings.Fields(s)
	ans := make(map[string]int)
	for _, v := range(s_split) {
		ans[v]++
	}
	return ans
}

func WordCount(s string) map[string]int {
	s_split := strings.Fields(s)
	// fmt.Println(s_split)
	ans := make(map[string]int)
	for _,v := range(s_split) {
		// fmt.Println(v)
		ans[v]++
		// fmt.Println(ans)
	}
	return ans
}

func myTest(myWC func(s string) map[string]int) {
	testData := "I ate a donut. Then I ate another donut."
	correctAns := map[string]int{"I":2, "ate":2, "a":1, "donut.":2, "Then":1, "another":1}
	myAns := myWC(testData)
	compareResult := reflect.DeepEqual(correctAns, myAns)
	if (compareResult == true) {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}

func main() {
	//myTest(WordCount)
	myTest(roseWordCount)
}
