/**********************************************************************************
*
* A peak element is an element that is greater than its neighbors.
*
* Given an input array where num[i] ≠ num[i+1], find a peak element and return its index.
*
* You may imagine that num[-1] = num[n] = -∞.
*
* For example, in array [1, 2, 3, 1], 3 is a peak element and your function should return the index number 2.
*
* click to show spoilers.
*
* Note:
* Your solution should be in logarithmic complexity.
*
* Credits:Special thanks to @ts for adding this problem and creating all test cases.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getPeak(l []int, low, high int) (int, int) {
    if len(l) < 3 {
        return -1, -1
    } else if len(l) == 3 {
        if (l[0] < l[1]) && (l[1] > l[2]) {
            return 1, l[1]
        } else {
            return -1, -1
        }
    }

    mid := low + (high-low)/2
    if l[mid-1] < l[mid] && l[mid] > l[mid+1] {
        return mid, l[mid]
    } else if l[mid-1] <= l[mid] && l[mid] <= l[mid+1] {
        return getPeak(l, mid, high)
    } else if l[mid-1] >= l[mid] && l[mid] >= l[mid+1] {
        return getPeak(l, low, mid)
    }
    return -1, -1
}

func main() {
    l := []int{
        4, 5, 6, 7, 1, 2,
        // 4, 5, 6, 7, 0, 1, 2,
        // 1, 2, 3, 1,
    }

    index, peak := getPeak(l, 0, len(l)-1)
    fmt.Println(index, peak)
}
