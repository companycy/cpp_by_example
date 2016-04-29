/**********************************************************************************
*
* Find the contiguous subarray within an array (containing at least one number)
* which has the largest product.
*
* For example, given the array [2,3,-2,4],
* the contiguous subarray [2,3] has the largest product = 6.
*
* More examples:
*
*   Input: arr[] = {6, -3, -10, 0, 2}
*   Output:   180  // The subarray is {6, -3, -10}
*
*   Input: arr[] = {-1, -3, -10, 0, 60}
*   Output:   60  // The subarray is {60}
*
*   Input: arr[] = {-2, -3, 0, -2, -40}
*   Output:   80  // The subarray is {-2, -40}
*
**********************************************************************************/

package main

import (
    "fmt"
    "math"
)

type Node struct {
    min, max float64
}

func getMaxProduct(l *[]int) float64 {
    if len(*l) == 0 {
        return 0
    } else if len(*l) == 1 {
        return float64((*l)[0])
    }

    dp := make([]Node, len(*l))
    dp[0].max, dp[0].min = float64((*l)[0]), float64((*l)[0])

    // low, high := 0, 0
    product := float64((*l)[0])
    for i := 1; i < len(*l); i++ {
        t1 := float64((*l)[i])
        t2, t3 := float64(dp[i-1].max*t1), float64(dp[i-1].min*t1)
        dp[i].max = math.Max(t1, math.Max(t2, t3))
        dp[i].min = math.Min(t1, math.Min(t2, t3))
        product = math.Max(product, dp[i].max)

        // if (*l)[i] > dp[i-1].max*(*l)[i] && (*l)[i] > product {
        //     product = (*l)[i]
        //     low, high = i, i
        // } else if dp[i-1].max*(*l)[i] > product {
        //     product = dp[i-1].max * (*l)[i]
        //     high = i
        //     if (*l)[i] == 0 {
        //         low = i
        //     } else if zeroPos < i {
        //         low = zeroPos + 1
        //     }
        // }

        // if (*l)[i] == 0 && i+1 < len(*l) {
        //     dp[i].max = 1
        //     zeroPos = i
        // } else {
        //     dp[i].max = dp[i-1].max * (*l)[i]
        // }
    }

    return product
}

func main() {
    l := [][]int{
        {2, 3, -2, 4},
        {6, -3, -10, 0, 2},
        {-6, -3, -10, 0, 2},
        {-1, -3, -10, 0, 60},
        {-2, -3, 0, -2, -40},
        {-2, -3, 0, -2, -40, 0, -3, -30},
        {-4, -3},
        {-1, -1},
        {-1, 0, -2},
    }
    for i := range l {
        fmt.Println(getMaxProduct(&l[i]))
    }
}
