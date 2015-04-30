/**********************************************************************************
*
* Say you have an array for which the ith element is the price of a given stock on day i.
*
* Design an algorithm to find the maximum profit. You may complete at most two transactions.
*
* Note:
* You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
*
**********************************************************************************/

package main

import (
    "fmt"
)

func main() {
    l := []int {
        1, 10, 25, 9, 16,
    }
    // max[0 - i] + max[i+1 - n]
    low := l[0]
    result1 := make([]int, len(l))
    for i := 1; i < len(l)-1; i++ {
        if low > l[i] {
            low = l[i]
        }
        result1[i] = result1[i-1]
        if result1[i] < l[i] - low {
            result1[i] = l[i] - low
        }
    }
    fmt.Println(result1)

    high := l[len(l)-1]
    result2 := make([]int, len(l))
    for i := len(l)-2; i >= 1; i-- {
        if high < l[i] {
            high = l[i]
        }
        if result2[i] < high - l[i] {
            result2[i] = high - l[i]
        }
    }
    fmt.Println(result2)
    
    var profit int
    for i := 0; i < len(l)-1; i++ {
        cur := result1[i] + result2[i]
        if profit < cur {
            profit = cur
        }
    }
    fmt.Println(profit)
}


