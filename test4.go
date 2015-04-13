package main

import (
    "fmt"
)

/**********************************************************************************
*
* Given an array of non-negative integers, you are initially positioned at the first index of the array.
*
* Each element in the array represents your maximum jump length at that position.
*
* Determine if you are able to reach the last index.
*
* For example:
* A = [2,3,1,1,4], return true.
*
* A = [3,2,1,0,4], return false.
*
*
**********************************************************************************/


func can_get_last(l []int) (bool) {
    for i := 0; i < len(l); i++ {
        if l[i] <= 0 {
            return false
        } else if l[i] >= (len(l) - i) {
            return true
        } else {
            // do nothing
        }
    }
    return false
}


func main() {
    // l := []int{2,3,1,1,4}
    l := []int{3,2,1,0,4}

    if can_get_last(l) {
        fmt.Println("success")
    } else {
        fmt.Println("failure")
    }
}

