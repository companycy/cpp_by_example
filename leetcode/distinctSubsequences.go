/**********************************************************************************
*
* Given a string S and a string T, count the number of distinct subsequences of T in S.
*
* A subsequence of a string is a new string which is formed from the original string
* by deleting some (can be none) of the characters without disturbing the relative positions
* of the remaining characters. (ie, "ACE" is a subsequence of "ABCDE" while "AEC" is not).
*
* Here is an example:
* S = "rabbbit", T = "rabbit"
*
* Return 3.
*
*
**********************************************************************************/

// http://blog.csdn.net/yangliuy/article/details/43526751
// 这题的题意是我们只能进行删除操作，有多少做方法可以把S变成T
// 当S[i]!=T[j]时 dp[i][j] = dp[i-1][j]，
// 因为此时我们只能把S前i-1个变成T的前j个，然后删除掉S[i]既可达到目的
// 当S[i]=T[j]时，显然上面的方法仍然可行，同时我们多了一种方法，就是把S的前i-1个变成T的前j-1个，
// 然后加上后面相等的字符自然可以匹配上
package main

import (
    "fmt"
)

func main() {
    s := "rabbbit"
    t := "rabbit"
    dp := make([][]int, len(s)+1)
    for i := range dp {
        dp[i] = make([]int, len(t)+1)
        dp[i][0] = 1
    }
    for i := range dp[0] {
        dp[0][i] = 0
    }
    dp[0][0] = 1
    // for i := range dp {
    //     fmt.Println(dp[i])
    // }

    for i := 1; i < len(s)+1; i++ {
        for j := 1; j < len(t)+1; j++ {
            if s[i-1] != t[j-1] {
                dp[i][j] = dp[i-1][j]
            } else {
                dp[i][j] = dp[i-1][j] + dp[i-1][j-1]
            }
            // fmt.Println(s[i-1], t[j-1], dp[i][j])
        }
        fmt.Println(dp[i])
    }

    fmt.Println(dp[len(s)][len(t)])
}

