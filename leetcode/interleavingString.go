/**********************************************************************************
*
* Given s1, s2, s3, find whether s3 is formed by the interleaving of s1 and s2.
*
* For example,
* Given:
* s1 = "aabcc",
* s2 = "dbbca",
*
* When s3 = "aadbbcbcac", return true.
* When s3 = "aadbbbaccc", return false.
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "strings"
)

func verifyInterleaving1(s1, s2, s3 string) bool {
    var ppos, pos int
    for i := range s1 {
        pos = strings.Index(s3[ppos:], string(s1[i]))
        if pos == -1 {
            return false
        } else {
            ppos = pos + ppos
        }
    }
    // fmt.Println("s1 exists")

    ppos, pos = 0, 0
    for i := range s2 {
        pos = strings.Index(s3[ppos:], string(s2[i]))
        fmt.Println(string(s2[i]), ppos, pos)
        if pos == -1 {
            return false
        } else {
            ppos = pos + ppos
        }
    }

    return true
}

func verifyInterleaving(s1, s2, s3 string) int {
    dp := make([][]int, len(s2)+1)
    for i := range dp {
        dp[i] = make([]int, len(s1)+1)
    }

    for i := 1; i < len(s2)+1; i++ {
        if s2[:i] == s3[:i] {
            dp[i][0] = 1
        } else {
            dp[i][0] = 0
        }
    }

    for i := 1; i < len(s1)+1; i++ {
        if s1[:i] == s3[:i] {
            dp[0][i] = 1
        } else {
            dp[0][i] = 0
        }
    }

    for i := 1; i < len(s2)+1; i++ {
        for j := 1; j < len(s1)+1; j++ {
            if dp[i-1][j-1] == 1 {
                if s3[i+j-2:i+j] == string(s1[j-1])+string(s2[i-1]) {
                    dp[i][j] = 1
                } else if s3[i+j-2:i+j] == string(s2[i-1])+string(s1[j-1]) {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = 0
                }
            } else {
                if dp[i-1][j] == 1 && s2[i-1] == s3[i+j-1] {
                    dp[i][j] = 1
                } else if dp[i][j-1] == 1 && s1[j-1] == s3[i+j-1] {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = 0
                }
            }
        }
    }
    fmt.Println(dp)
    return dp[len(s2)][len(s1)]
}

func main() {
    s1 := "aabcc"
    s2 := "dbbca"
    s3 := "aadbbcbcac"
    s3 = "aadbbbaccc"

    fmt.Println(verifyInterleaving(s1, s2, s3))
}
