package main

import "fmt"

func fibonacci() func() int {
	var p1, p2, curr int
	return func() int {
		switch {
		case curr == 0:
			curr = 1
		case p1 == 0:
			p1 = 1
		default:
			curr = p1 + p2
			p2 = p1
			p1 = curr
		}
		return p1
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
