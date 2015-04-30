/**********************************************************************************
*
* Given two words word1 and word2, find the minimum number of steps required to
* convert word1 to word2. (each operation is counted as 1 step.)
*
* You have the following 3 operations permitted on a word:
*
* a) Insert a character
* b) Delete a character
* c) Replace a character
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func GetMinSteps(from, to string, dp [][]int) int {
    if len(from) == 0 {
        return len(from)
    } else if len(to) == 0 {
        return len(to)
    }

    for i := 1; i < len(from)+1; i++ {
        for j := 1; j < len(to)+1; j++ {
            if from[i-1] == to[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = dp[i-1][j-1] + 1
                if dp[i][j] > dp[i-1][j]+1 {
                    dp[i][j] = dp[i-1][j] + 1
                }
                if dp[i][j] > dp[i][j-1]+1 {
                    dp[i][j] = dp[i][j-1] + 1
                }
            }
        }
    }

    return dp[len(from)][len(to)]
}

func main() {
    from := "abb"
    to := "abccb"
    dp := make([][]int, len(from)+1)
    for i := range dp {
        dp[i] = make([]int, len(to)+1)
        dp[i][0] = i
    }
    for i := range dp[0] {
        dp[0][i] = i
    }

    minsteps := GetMinSteps(from, to, dp)
    fmt.Println(minsteps)
}
