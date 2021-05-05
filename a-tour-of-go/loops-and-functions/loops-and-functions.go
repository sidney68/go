package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x <= 0 {
		return math.NaN()
	}
	z := 1.0
	for i := 0; i < 32; i++ {
		prev := z
		z -= (z*z - x) / (2 * z)
		if math.Abs(z-prev) <= 0.0000000001 {
			return z
		}
	}
	return math.NaN()
}

func main() {
	v := 27.0
	fmt.Printf("Sqrt(%v) = %f [should be %f]", v, Sqrt(v), math.Sqrt(v))
}
