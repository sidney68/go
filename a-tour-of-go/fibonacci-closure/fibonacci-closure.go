package main

import (
	"fmt"
)

func fibonacci() func() int {
	var x, y = 0, 1
	return func() int {
		z := x
		x, y = y, x+y
		return z
	}
}

func main() {
	f := fibonacci()
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = f()
	}
	fmt.Println(arr)
}
