package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"Daniel", 27}
	fmt.Println(p)
}
