package main

/**********************************************************************************
*
* Given an array of non-negative integers, you are initially positioned at the first index of the array.
*
* Each element in the array represents your maximum jump length at that position.
*
* Your goal is to reach the last index in the minimum number of jumps.
*
* For example:
* Given array A = [2,3,1,1,4]
*
* The minimum number of jumps to reach the last index is 2.
* (Jump 1 step from index 0 to 1, then 3 steps to the last index.)
*
*
**********************************************************************************/


import (
        "fmt"
)

func main() {
    l := []int {2,3,1,1,4}
    steps := make([]int, len(l))
    for i := 0; i < len(l); i++ {
        max := l[i] + i
        for j := i+1; j <= max && j < len(l); j++ {
            if j-i+l[j] > tmp {
                steps[i] = j
            }
        }
    }

    for i := 0; i < len(l); i++ {
        fmt.Println(steps[i])
    }
}
