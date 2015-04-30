/**********************************************************************************
*
* Say you have an array for which the ith element is the price of a given stock on day i.
*
* Design an algorithm to find the maximum profit. You may complete as many transactions
* as you like (ie, buy one and sell one share of the stock multiple times). However,
* you may not engage in multiple transactions at the same time (ie, you must sell the
* stock before you buy again).
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
    var profit int
    for i := 1; i < len(l); i++ {
        if l[i] > l[i-1] {
            profit += l[i] - l[i-1]
        }
    }
    fmt.Println(profit)
}

