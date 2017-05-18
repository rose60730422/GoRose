package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dx)
	for x := 0; x < dx; x++ {
		pic[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			pic[x][y] = uint8((x + y) / 2)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}

func Pic(dx, dy int) [][]uint8 {
	// 宣告畫布
	pic := make([][]uint8, dx)
	// 填 x 方向的 slice
	for x := 0; x < dx; x++ {
		// 填 y 方向的 slice
		pic[x] = make([]uint8, dy)
		for y := 0; y < dy; y++ {
			pic[x][y] = uint8(x ^ (2 * x * y) + y ^ y)
		}
	}
	return pic
}

func main() {
	pic.Show(Pic)
}
