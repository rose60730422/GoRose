package main

import "fmt"

func main() {
	str := "hello, world!"
	for tmpStr := str[:]; len(tmpStr) > 0; tmpStr = tmpStr[1:] {
		fmt.Println(tmpStr)
	}
}
