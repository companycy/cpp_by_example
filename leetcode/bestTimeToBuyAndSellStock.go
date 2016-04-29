/**********************************************************************************
*
* Say you have an array for which the ith element is the price of a given stock on day i.
*
* If you were only permitted to complete at most one transaction (ie, buy one and sell one share of the stock),
* design an algorithm to find the maximum profit.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func solution2() {
    l := []int {
        1, 10, 25, 9, 16,
    }
    var profit, cursum int
    for i := 1; i < len(l); i++ {
        if cursum <= 0 {
            cursum = l[i] - l[i-1]
        } else {
            cursum += l[i] - l[i-1]
        }
        if cursum > profit {
            profit = cursum
        }
    }
    fmt.Println(profit)
}

func solution1() {
    l := []int {
        1, 10, 25, 9, 16,
    }
    var profit int
    low := 65535
    for i := 0; i < len(l); i++ {
        if l[i] < low {
            low = l[i]
        }
        if profit < l[i] - low {
            profit = l[i] - low
        }
    }
    fmt.Println(profit)
}

func main() {
    solution1()
    solution2()
}
