/*
 * Given a string, find the length of the longest substring T that contains at most 2 distinct characters.
 *
 * For example, Given s = “eceba”,
 *
 * T is "ece" which its length is 3.
 */

package main

import (
    "fmt"
)

func getSubstrWith2DistinctChar(s string) string {
    dp := make([][]int, len(s))
    for i := range dp {
        dp[i] = make([]int, len(s))
    }

    var result string
    for i := range dp {
        m := make(map[uint8]int)
        for j := i; j < len(dp[i]); j++ {
            if i == j {
                m[s[j]] = j
                dp[i][j] = 1
            } else if _, ok := m[s[j]]; ok {
                dp[i][j] = dp[i][j-1]
            } else if !ok {
                m[s[j]] = j
                dp[i][j] = dp[i][j-1] + 1
            }
            if dp[i][j] == 2 && j-i+1 > len(result) {
                result = s[i : j+1]
            }
        }
        // fmt.Println(dp[i], m)
    }

    return result
}

func main() {
    s := []string{
        "eceba",
    }
    for i := range s {
        fmt.Println(getSubstrWith2DistinctChar(s[i]))
    }
}
