/**********************************************************************************
*
* Implement next permutation, which rearranges numbers into the lexicographically next
* greater permutation of numbers.
*
* If such arrangement is not possible, it must rearrange it as the lowest possible order
* (ie, sorted in ascending order).
*
* The replacement must be in-place, do not allocate extra memory.
*
* Here are some examples. Inputs are in the left-hand column and its corresponding outputs
* are in the right-hand column.
*
*   1,2,3 → 1,3,2
*   3,2,1 → 1,2,3
*   1,1,5 → 1,5,1
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getNextPermutation(l *[]int) *[]int {
    i := len(*l) - 1
    for ; i >= 1 && (*l)[i-1] >= (*l)[i]; i-- {
    }

    if i != 0 {
        min, k := 65535, 0
        for j := len(*l) - 1; j >= i; j-- {
            if (*l)[j] > (*l)[i-1] && min > (*l)[j] {
                min, k = (*l)[j], j
            }
        }

        t := (*l)[i-1]
        (*l)[i-1] = (*l)[k]
        (*l)[k] = t
    }

    for n, m := i, len(*l)-1; n < m; n, m = n+1, m-1 {
        t := (*l)[n]
        (*l)[n] = (*l)[m]
        (*l)[m] = t
    }
    return l
}

func main() {
    s := [][]int{
        []int{1, 2, 3},
        []int{3, 2, 1},
        []int{1, 1, 5},
    }
    for i := range s {
        fmt.Println(getNextPermutation(&s[i]))
    }
}
