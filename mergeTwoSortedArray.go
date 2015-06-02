/**********************************************************************************
*
* Given two sorted integer arrays A and B, merge B into A as one sorted array.
*
* Note:
*   You may assume that A has enough space (size that is greater or equal to m + n)
*   to hold additional elements from B. The number of elements initialized in A and B
*   are m and n respectively.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func mergeArrays(l, r *[]int) *[]int {
    if len(*r) == 0 {
        return l
    } else if len(*l) == 0 {
        *l = append((*l), (*r)...)
        return l
    }

    var newl []int
    i, j := 0, 0
    for i < len(*l) && j < len(*r) {
        if (*l)[i] < (*r)[j] {
            newl = append(newl, (*l)[i])
            i++
        } else if (*l)[i] > (*r)[j] {
            newl = append(newl, (*r)[j])
            j++
        } else {
            newl = append(newl, (*l)[i], (*r)[i])
            i, j = i+1, j+1
        }
    }
    if i < len(*l) {
        newl = append(newl, (*l)[i:]...)
    }
    if j < len(*r) {
        newl = append(newl, (*r)[j:]...)
    }
    return &newl
}

func main() {
    l := [][]int{
        {1, 5, 8, 10},
        {2, 4, 6, 8, 10, 0, 0, 0},
        {2, 4, 0, 0, 0},
        {12, 14, 16, 18, 20, 0, 0, 0},
        {2, 0},
        {0, 0, 0},
        {2, 4, 6, 8, 10, 0, 0, 0},
        {2, 4, 0, 0, 0, 0, 0, 0},
    }
    r := [][]int{
        {7, 15, 20, 40},
        {1, 3, 5},
        {3, 5, 7},
        {1, 3, 5},
        {3},
        {1, 3, 5},
        {11, 13, 15},
        {1, 3, 5, 7, 9, 11},
    }

    for i := range l {
        fmt.Println(mergeArrays(&l[i], &r[i]))
    }
}
