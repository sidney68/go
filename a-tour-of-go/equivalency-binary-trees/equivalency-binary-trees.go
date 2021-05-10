package main

import (
    "fmt"
    "golang.org/x/tour/tree"
)

func walkTree(t *tree.Tree, c chan int) {
    if t.Left != nil {
        walkTree(t.Left, c)
    }
    c <- t.Value
    if t.Right != nil {
        walkTree(t.Right, c)
    }
}

func Walk(t *tree.Tree, c chan int) {
    defer close(c)
    walkTree(t, c)
}

func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := make(chan int), make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for v := range c1 {
        if v != <-c2 {
            return false
        }
    }
    if _, ok := <-c2; ok == true {
        return false
    }
    return true
}

func main() {
    c := make(chan int)
    go Walk(tree.New(1), c)
    for i := 0; i < 10; i++ {
        fmt.Println(<-c)
    }
    println("Same(tree.New(1), tree.New(1))", Same(tree.New(1), tree.New(1)))
    println("Same(tree.New(1), tree.New(2))", Same(tree.New(1), tree.New(2)))
    println("Same(tree.New(2), tree.New(1))", Same(tree.New(2), tree.New(1)))
    println("Same(tree.New(2), tree.New(2))", Same(tree.New(2), tree.New(2)))
}
