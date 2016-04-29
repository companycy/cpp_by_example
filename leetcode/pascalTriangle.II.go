/**********************************************************************************
*
* Given an index k, return the kth row of the Pascal's triangle.
*
* For example, given k = 3,
* Return [1,3,3,1].
*
* Note:
* Could you optimize your algorithm to use only O(k) extra space?
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getNthRow2(n int) *[]int {
    l := make([]int, n+1)
    l[0] = 1

    for i := 0; i < len(l)-1; i++ {
        for j := i + 1; j > 0; j-- {
            l[j] += l[j-1]
        }
    }
    return &l
}

func getNthRow(n int) *[]int {
    prevlvl := []int{1}

    if n == 1 {
        return &prevlvl
    }

    var curlvl []int
    for i := 1; i <= n; i++ {
        curlvl = []int{prevlvl[0]}
        for j := 1; j < i; j++ {
            curlvl = append(curlvl, prevlvl[j-1]+prevlvl[j])
        }
        curlvl = append(curlvl, prevlvl[len(prevlvl)-1])
        prevlvl = curlvl
    }
    return &curlvl
}

func main() {
    k := 3
    fmt.Println(getNthRow(k))
    fmt.Println(getNthRow2(k))
}
