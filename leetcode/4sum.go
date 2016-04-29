/**********************************************************************************
*
* Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target?
* Find all unique quadruplets in the array which gives the sum of target.
*
* Note:
*
* Elements in a quadruplet (a,b,c,d) must be in non-descending order. (ie, a ≤ b ≤ c ≤ d)
* The solution set must not contain duplicate quadruplets.
*
*     For example, given array S = {1 0 -1 0 -2 2}, and target = 0.
*
*     A solution set is:
*     (-1,  0, 0, 1)
*     (-2, -1, 1, 2)
*     (-2,  0, 0, 2)
*
*
**********************************************************************************/


// todo: when there is duplicated, need to customize map

package main

import (
    "fmt"
    "sort"
)


func sum4(l []int, sum int) {
    for i := 0; i < len(l); i++ {
        // new l except l[i]
        // m := make([]int), len(l)-1)
        var m []int
        for j := 0; j < i; j++ {
            m = append(m, l[j])
        }
        for j := i+1; j < len(l); j++ {
            m = append(m, l[i])
        }
        newsum := sum - l[i]
        sum3(m, newsum)
    }
}

func sum3(l []int, sum int) {
    for i := 0; i < len(l); i++ {
        // new l except l[i]
        // m := make([]int), len(l)-1)
        var m []int
        for j := 0; j <= i-1; j++ {
            m = append(m, l[j])
        }
        for j := i+1; j < len(l); j++ {
            m = append(m, l[j])
        }
        newsum := sum - l[i]
        sum2(m, newsum)
    }
}

func sum2(l []int, sum int) {
    for i := 0; i < len(l); i++ {
        m := make(map[int]int, len(l)-1)
        // to avoid duplication, skip l[i]
        for j := 0; j < i-1; j ++ {
            m[l[j]] = j
        }
        for j := i+1; j < len(l); j ++ {
            m[l[j]] = j
        }
        newsum := sum - l[i]
        if _, ok := m[newsum]; ok {
            fmt.Println(l[i], newsum, )
        }
    }
}

func main() {
    l := []int {1, -1, -2, 2, 0, 0}
    sort.Ints(l)
    fmt.Println(l)

    target := 0
    sum3(l, target)
}

