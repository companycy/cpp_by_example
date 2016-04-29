package main

import (
    "fmt"
)

func main() {
    // l := make([]int)
    var l []int
    l = append(l, 1)
    l = append(l, 2)
    for i := 0; i < len(l); i++ {
        fmt.Println(l[i])
    }

    var result = float32(6 * 1.0 / 5)
    // result := 6 / 5
    fmt.Println(result)

    m := make(map[int]int)
    m[1] = 11
    if v, ok := m[1]; ok {
        v++
        fmt.Println(v)
    }
    fmt.Println(m)
}
