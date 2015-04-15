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
    step := make([]int, len(l))
    // f(n, m) = min(m, f(n, m-1)+1, max) X
    for i := 0; i < len(l); i++ {
        step[i] = i
    }

    for i := 0; i < len(l); i++ {
        max := l[i] + i
        currentStep := step[i] + 1
        for j := i+1; j <= max && j < len(l); j++ {
            if currentStep < step[j] {
                step[j] = currentStep
            }
        }
	}
	for i := 0; i < len(step); i++ {
    	fmt.Println(step[i])
	}
}
