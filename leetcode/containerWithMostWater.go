/**********************************************************************************
*
* Given n non-negative integers a1, a2, ..., an, where each represents a point at coordinate (i, ai).
* n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
*
* Find two lines, which together with x-axis forms a container, such that the container contains the most water.
*
* Note: You may not slant the container.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

// if max(i) is biggest container at i
// max(i+1) = max{max(i), min{l[i], l[j]} * (i-j) where j = 0->i}
func getMaxContainer(l *[]int) (int, int, int) {
    if len(*l) < 2 {
        return -1, -1, 0
    }

    var max int
    if (*l)[0] < (*l)[1] {
        max = (*l)[0] * 1
    } else {
        max = (*l)[1] * 1
    }

    if len(*l) == 2 {
        return 0, 1, max
    }

    top, bottom := 0, 1
    for i := 3; i < len(*l); i++ {
        for j := 0; j < i; j++ {
            var newmax int
            if (*l)[j] < (*l)[i] {
                newmax = (*l)[j] * (i - j)
            } else {
                newmax = (*l)[i] * (i - j)
            }
            if max < newmax {
                top, bottom = j, i
                max = newmax
            }
        }
    }
    return top, bottom, max
}

func main() {
    l := &[]int{
        1, 2, 4, 1, 5,
    }
    fmt.Println(getMaxContainer(l))
}
