package main

import "fmt"

func max(a, b, c int) int {
    if a >= b && a >= c {
        return a
    } else if b >= a && b >= c {
        return b
    } else {
        return c
    }
}

func main() {
    fmt.Println(max(5, 10, 3))
}

