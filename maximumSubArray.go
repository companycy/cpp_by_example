/**********************************************************************************
*
* Find the contiguous subarray within an array (containing at least one number)
* which has the largest sum.
*
* For example, given the array [−2,1,−3,4,−1,2,1,−5,4],
* the contiguous subarray [4,−1,2,1] has the largest sum = 6.
*
* More practice:
*
* If you have figured out the O(n) solution, try coding another solution using
* the divide and conquer approach, which is more subtle.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getMaxSum(l *[]int) int {
    if len(*l) == 0 {
        return 0
    } else if len(*l) == 1 {
        return (*l)[0]
    }

    prevMaxSum, curMaxSum, maxSum := (*l)[0], 0, 0
    for i := 1; i < len(*l); i++ {
        if prevMaxSum+(*l)[i] > 0 {
            curMaxSum = prevMaxSum + (*l)[i]
        } else {
            curMaxSum = 0
        }
        // fmt.Println(maxSum, prevMaxSum, curMaxSum)
        prevMaxSum = curMaxSum
        if maxSum < curMaxSum {
            maxSum = curMaxSum
        }
    }

    return maxSum
}

func main() {
    l := [][]int{
        {-2, 1, -3, 4, -1, 2, 1, -5, 4},
    }
    for i := range l {
        fmt.Println(getMaxSum(&l[i]))
    }
}
