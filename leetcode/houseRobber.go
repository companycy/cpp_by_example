/**********************************************************************************
 *
 * You are a professional robber planning to rob houses along a street. Each house has
 * a certain amount of money stashed, the only constraint stopping you from robbing
 * each of them is that adjacent houses have security system connected and it will
 * automatically contact the police if two adjacent houses were broken into on the same night.
 *
 * Given a list of non-negative integers representing the amount of money of each house,
 * determine the maximum amount of money you can rob tonight without alerting the police.
 *
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

func getMax(l *[]int) int {
    if len(*l) == 0 {
        return 0
    } else if len(*l) == 1 {
        return (*l)[0]
    }

    var p int
    if (*l)[0] > (*l)[1] {
        p = (*l)[0]
    } else {
        p = (*l)[1]
    }

    if len(*l) == 2 {
        return p
    }

    pp, result := (*l)[0], 0
    for i := 2; i < len(*l); i++ {
        result = p
        if result < pp+(*l)[i] {
            result = pp + (*l)[i]
        }

        pp = p
        p = result
    }

    return result
}

func main() {
    l := []int{
        // 1, 5, 9,
        2, 1, 3, 4,
    }

    max := getMax(&l)
    fmt.Println(max)
}
