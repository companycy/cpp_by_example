/**********************************************************************************
*
* There are two sorted arrays A and B of size m and n respectively.
* Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getMedian(l, r *[]int, k int) int {
    if len(*l) == 0 && len(*r) == 0 {
        return 0
    } else if len(*l) > len(*r) {
        return getMedian(r, l, k)
    } else if len(*l) == 0 {
        return r[k-1]
    } else if k == 1 {
        median := (*l)[0]
        if median < (*r)[0] {
            median = (*r)[0]
        }
        return median
    }

    var median int
    lk := k / 2
    if lk < len(*l) {
        lk = len(*l)
    }
    rk := k - lk
    if (*l)[lk-1] < (*r)[rk-1] {
        return getMedian((*l)[lk:], r, k-lk)
    } else if (*l)[lk-1] > (*r)[rk-1] {
        return getMedian(l, r[rk:], k-rk)
    } else {
        return (*l)[lk-1]
    }

    return median
}

func main() {
    l := [][]int{
        // {1, 3, 6, 8,},
        {1, 12, 15, 26, 38},
        {1, 12, 15, 26, 38},
    }
    r := [][]int{
        // {2, 5, 8, 9, 15,},
        {2, 13, 17, 30, 45, 50},
        {2, 13, 17, 30, 45},
    }
    for i := range l {
        mn := len(l[i]) + len(r[i])
        if mn%2 == 0 {
            median := (getMedian(&l[i], &r[i], mn/2-1) + getMedian(&l[i], &r[i], mn/2)) / 2
            fmt.Println(median)
        } else {
            median := getMedian(&l[i], &r[i], mn/2)
            fmt.Println(median)
        }
        // var sum int
        // for j := range l[i] {
        //     sum += l[i][j]
        // }
        // for j := range r[i] {
        //     sum += r[i][j]
        // }
        // fmt.Println(sum, sum/(len(r[i])+len(l[i])))
    }
}
