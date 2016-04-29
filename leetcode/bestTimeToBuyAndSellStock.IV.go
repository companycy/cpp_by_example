
package main

import (
    "fmt"
    "sort"
)

func main() {
    l := []int {
        1, 10, 25, 9, 16,
    }

    // dp[i, j] = max{dp[i, j-1], dp[i-1, t] + l[j] - l[t+1]}
    // result := make([]int, len(l))
    var dp [k][len(l)]int
    for i := 0; i < k; i++ {
        for j := 0; j < len(l); j++ {
            dp[i][j] = dp[i][j-1]
            max := 0
            for t := 0; t < j; t++ {
                tmp := dp[i-1][t]+l[j]-l[t+1]
                if max < tmp {
                    max = tmp
                }
            }
            if dp[i][j] < max {
                dp[i][j] = max
            }
        }
    }


}
