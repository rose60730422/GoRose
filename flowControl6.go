package main

import (
	"fmt"
	"math"
)

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func pow2(x, n, lim float64) float64 {
	var ans float64
	if v := math.Pow(x, n); v < lim {
		ans = v
	} else {
		ans = lim
	}
	return ans
}

func pow3(x, n, lim float64) float64 {
	var ans float64
	ans = lim
	if v := math.Pow(x, n); v < lim {
		ans = v
	}

	return ans
}

func main() {
	fmt.Println(
		pow1(2, 5, 300),
		pow1(3, 3, 20),
	)
}
