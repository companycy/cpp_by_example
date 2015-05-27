/**********************************************************************************
*
* Given an unsorted integer array, find the first missing positive integer.
*
* For example,
* Given [1,2,0] return 3,
* and [3,4,-1,1] return 2.
*
* Your algorithm should run in O(n) time and uses constant space.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getMissingPositiveInt(l []int) int {
    // var t int
    for i := 0; i < len(l); i++ {
        if l[i] <= 0 || i == l[i]-1 {

        } else {
            for l[i] > 0 && i != l[i]-1 {
                if l[i]-1 >= len(l) {
                    continue
                }
                t := l[l[i]-1]
                l[l[i]-1] = l[i]
                l[i] = t
            }
            // i++
        }
    }

    for i := 0; i < len(l); i++ {
        if l[i] != i+1 {
            return i + 1
        }
    }
    return len(l) + 1
}

func main() {
    l := []int{
        // 1, 2, 0,
        3, 4, -1, 1,
    }

    i := getMissingPositiveInt(l)
    fmt.Println(i)
}
