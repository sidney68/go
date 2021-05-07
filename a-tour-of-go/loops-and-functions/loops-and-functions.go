package main

import (
    "fmt"
    "math"
)

type ErrorSqrt float64

func (e ErrorSqrt) Error() string {
    return fmt.Sprintf("[ERROR] Cannot datermine sqrt of value: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    switch {
    case x < 0:
        return math.NaN(), ErrorSqrt(x)
    case x == 0:
        return 0, nil
    default:
        z := 1.0
        for i := 0; i < 32; i++ {
            y := z
            z -= (z*z - x) / (2 * z)
            if math.Abs(z-y) <= 1e-9 {
                return z, nil
            }
        }
    }
    return math.NaN(), ErrorSqrt(x)
}

func main() {
    eval(2)
    eval(0)
    eval(-2)
}

func eval(x float64) {
    if r, e := Sqrt(x); e == nil {
        fmt.Printf(" => Sqrt(%v) = %f [should be %f]\n", x, r, math.Sqrt(x))
    } else {
        fmt.Println(e)
    }
}
