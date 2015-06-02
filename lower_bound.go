package main

import (
    "fmt"
)

func lower_bound1(l *[]int, low, high, target int) int {
    len := high - low + 1
    for len > 0 {
        mid := low + len/2
        if target > (*l)[mid] {
            low = mid + 1
        }
        len = len / 2
    }
    return low
}

// first pos of num bigger than target
func lower_bound(l *[]int, low, high, target int) int {
    var pos int
    for low < high {
        mid := (low + high) / 2
        // fmt.Println(low, high, mid)
        if target > (*l)[mid] {
            low = mid + 1
            pos = low
        } else {
            high = mid
            pos = high
        }
    }
    return pos
}

func main() {
    l := []int{
        0, 1, 1, 1, 2, 33, 55,
    }

    target := 2
    fmt.Println(lower_bound(&l, 0, len(l)-1, target))
    fmt.Println(lower_bound1(&l, 0, len(l)-1, target))
}
