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
    var t int
    for i := range l {
        t = l[l[i]]
        l[l[i]] = l[i]
    }

    for i := 1; i < len(l); i++ {

    }
}

func main() {
    l := []int{
        1, 2, 0,
        // 3, 4, -1, 1,
    }

    i := getMissingPositiveInt(l)
    fmt.Println(i)
}
