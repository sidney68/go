package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for x := range arr {
		arr[x] = make([]uint8, dx)
		for y := range arr[x] {
			arr[x][y] = uint8(x ^ y)
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
}
