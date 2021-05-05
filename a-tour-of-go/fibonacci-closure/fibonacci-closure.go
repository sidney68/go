package main

import "fmt"

func fibonacci() func() int {
    var p1, p2 int
    return func() int {
        curr := p1 + p2
        switch {
        case p1 == 0 && p2 == 0:
            p1 = 1
        default:
            p1 = p2
            p2 = curr
        }
        return curr
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
