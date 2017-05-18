package main

import "fmt"

func main() {
	beverage := make(map[string]int)
	beverage["珍珠奶茶"] = 40
	beverage["青草茶"] = 90
	t := 2*beverage["珍珠奶茶"] + beverage["青草茶"]
	fmt.Println(t)
}
