/**********************************************************************************
*
* Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle
* containing all ones and return its area.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getMaxSubRectArea(l *[][]int) int {
    ret := make([][]int, len(*l))
    for i := range ret {
        ret[i] = make([]int, len((*l)[i]))
        ret[i][0] = (*l)[i][0]
    }
    for i := range ret[0] {
        ret[0][i] = (*l)[0][i]
    }

    var maxlen int
    for i := 1; i < len(*l); i++ {
        for j := 1; j < len((*l)[i]); j++ {
            if (*l)[i][j] == 0 {
                ret[i][j] = 0
            } else if (*l)[i][j] == 1 {
                ret[i][j] = ret[i-1][j-1] + 1
                if ret[i][j] > ret[i][j-1]+1 {
                    ret[i][j] = ret[i][j-1] + 1
                }
                if ret[i][j] > ret[i-1][j]+1 {
                    ret[i][j] = ret[i-1][j] + 1
                }
                if maxlen < ret[i][j] {
                    maxlen = ret[i][j]
                }
            }
        }
    }
    fmt.Println(ret)
    return maxlen * maxlen
}

func main() {
    l := [][]int{
        {1, 0, 1, 1, 1},
        {0, 1, 0, 1, 1},
        {1, 1, 1, 0, 1},
        {1, 1, 1, 1, 0},
        {1, 1, 1, 1, 0},
    }
    fmt.Println(getMaxSubRectArea(&l))
}
