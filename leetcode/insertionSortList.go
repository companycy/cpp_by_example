/**********************************************************************************
*
* Sort a linked list using insertion sort.
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val  int
    next *Node
}

func insertionSort(low, high, val int, l *[]int) {
    if low > high || high >= len(*l) {
        return
    } else if len(*l) == 1 {
        if val < (*l)[0] {
            (*l) = []int{val, (*l)[0]}
        } else {
            (*l) = []int{(*l)[0], val}
        }
        return
    } else if val >= (*l)[high] {
        (*l) = append((*l), val)
        return
    } else if val <= (*l)[low] {
        (*l) = append([]int{val}, (*l)...)
        return
    } else if low+1 == high && val >= (*l)[low] && val <= (*l)[high] {
        newl := make([]int, len(*l)-high)
        copy(newl, (*l)[high:])
        (*l) = (*l)[:len(*l)+1]
        (*l)[high] = val
        copy((*l)[high+1:], newl)
        return
    }

    mid := low + (high-low)/2
    if (*l)[low] <= val && val <= (*l)[mid] {
        insertionSort(low, mid, val, l)
    } else if (*l)[mid] <= val && val <= (*l)[high] {
        insertionSort(mid, high, val, l)
    }
}

func main() {
    head := Node{
        3, &Node{
            1, &Node{
                5, &Node{
                    2, nil}}},
    }

    var result []int
    for p := &head; p != nil; p = p.next {
        if len(result) == 0 {
            result = append(result, p.val)
        } else {
            low, high := 0, len(result)-1
            insertionSort(low, high, p.val, &result)
        }
    }

    fmt.Println(result)
}
