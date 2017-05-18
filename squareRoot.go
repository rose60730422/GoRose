package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	var zn1 float64
	zn := 1.0
	for n := 1; n < 50; n++ {
		zn1 = zn - (zn*zn-x)/(2*zn)
		if (zn1-zn)*(zn1-zn) < 0.000000000000000000000000000001 {
			break
		}
		zn = zn1
	}
	return zn
}

func main() {
	z := 3.0
	fmt.Println(sqrt(z))
	fmt.Println(math.Sqrt(z))
}
