package main

import (
    "fmt"
)

func fibonacci() func() int {
    var x, y int
    return func() int {
        z := x + y
        switch {
        case x == 0 && y == 0:
            x = 1
        default:
            x = y
            y = z
        }
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
