package main

import "fmt"

func main() {
	fmt.Println(month(3))
}

func month(m int) string {
	if m == 5 {
		return "May"
	} else if m == 4 {
		return "April"
	} else {
		fmt.Println("i dont know")
		return "123"
	}

}
