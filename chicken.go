package main

import (
	"fmt"
	"strings"
	"strconv"
)

func danielAns(input string) string {
	var inputData = strings.Fields(input)
	var myAnswer string
	// fmt.Println(inputData)

	for i := range(inputData) {
		if i % 2 == 0 {
			var legs, amount, amountChicken, amountRabbit int
			
			// fmt.Println(inputData[i], inputData[i + 1])

			// 1. convert inputData[i] from stirng to int
			// 2. assign to legs
			if intFromStr, err := strconv.Atoi(inputData[i]);err != nil{
				return ""
			} else {
				legs = intFromStr
			}

			// 1. convert inputData[i + 1] from stirng to int
			// 2. assign to amount
			if intFromStr, err := strconv.Atoi(inputData[i+1]);err != nil{
				return ""
			} else {
				amount = intFromStr
			}
			

			// fmt.Printf("legs:%d, amount:%d\n", legs, amount)
			amountChicken = (4 * amount - legs) / 2
			amountRabbit = amount - amountChicken
			// fmt.Printf("chichen: %d, rabbit: %d\n", amountChicken, amountRabbit)
			myAnswer = myAnswer  + strconv.Itoa(amountChicken) + " " + strconv.Itoa(amountRabbit) + "\n"		
		}
		
	}
	myAnswer = myAnswer[:len(myAnswer) - 1]

	return myAnswer
}

func roseAns(input string) string {
	var inputData = strings.Fields(input)
	var roseAnswer string

	for i := range(inputData){
		if i % 2 == 0 {
			var legs , amount , chickenamount , rabbitamount int
			if number , err := strconv.Atoi(inputData[i]) ; err != nil{
				return ""
			} else{
				legs = number
			}
			if number , err := strconv.Atoi(inputData[i + 1]) ; err != nil{
				return ""
			} else{
				amount = number
			}
			rabbitamount = ( legs - 2 * amount ) / 2
			chickenamount = amount - rabbitamount	
			roseAnswer = roseAnswer + strconv.Itoa(chickenamount) + " " + strconv.Itoa(rabbitamount) +"\n"
		}
	
	}
	roseAnswer = roseAnswer[:len(roseAnswer) - 1 ]
	return roseAnswer
}
	
	// 把輸入資料切成一格一格
	// 一組一組資料處理
	// {
	// 		把資料轉成int
	// 		運算
	// 		接到答案字串裡面
	// }
	// 回傳答案


func main() {
	fmt.Println("testing input:")
	var input = "36 10\n44 20\n142 40\n268 100\n666 200\n80 25"
	var expected = "2 8\n18 2\n9 31\n66 34\n67 133\n10 15"
	fmt.Println(input)
	fmt.Println("Rose's ans is:")
	ansReturn := roseAns(input)
	fmt.Println(ansReturn)
	fmt.Println("Expeted ans:")
	fmt.Println(expected)
	if (ansReturn == expected) {
		fmt.Println("PASS!!!")
	} else {
		fmt.Println("FAILLL!!!")
	}
}