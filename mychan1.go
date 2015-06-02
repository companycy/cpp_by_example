package main

import (
    "fmt"
    "math/rand"
    "time"
)

func foo(n int) {
    for i := 0; i < n; i++ {
        fmt.Println(i, n)
        amt := time.Duration(rand.Intn(250))
        time.Sleep(amt * time.Millisecond)
    }
}

func main() {
    for i := 0; i < 10; i++ {
        go foo(10)
    }
    var input string
    fmt.Scanln(&input)
}
