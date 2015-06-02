/*
 *    Given two strings S and T, determine if they are both one edit distance apart.
 */

package main

import (
    "fmt"
)

func isOneEditDistance2(s, t string) bool {
    if len(s) > len(t) {
        s, t = t, s
    }
    var equalLen bool
    if len(s) == len(t) {
        equalLen = true
    } else if len(s) == len(t)-1 {
        equalLen = false
    } else {
        return false
    }

    for i, j := 0, 0; i < len(s) && j < len(t); i, j = i+1, j+1 {
        if s[i] != t[j] {
            if !equalLen {
                i--
            }

            for i, j = i+1, j+1; i < len(s) && j < len(t); i, j = i+1, j+1 {
                if s[i] != t[j] {
                    return false
                }
            }
        }
    }
    return true
}

func isOneEditDistance1(s, t string) bool {
    if len(s) > len(t) {
        s, t = t, s
    }
    if len(s) < len(t)-1 {
        return false
    }

    i, j := 0, 0
    var editForL, editForR int
    for i < len(s) && j < len(t) {
        if s[i] == t[j] {
            i, j = i+1, j+1
        } else {
            if i < len(s)-1 && s[i+1] == t[j] {
                editForL++
                i++
            } else if j < len(t)-1 && s[i] == t[j+1] {
                editForR++
                j++
            } else if i < len(s)-1 && j < len(t)-1 && s[i+1] == t[j+1] {
                editForL, editForR = editForL+1, editForR+1
                i, j = i+1, j+1
            } else {
                return false
            }

            if editForL > 1 || editForR > 1 {
                return false
            }
        }
    }
    if i != len(s) {
        editForR += len(s) - i
    }
    if j != len(t) {
        editForL += len(t) - j
    }
    // fmt.Println(editForL, editForR)
    if editForL > 1 || editForR > 1 {
        return false
    } else {
        return true
    }
}

func isOneEditDistance(s, t string) bool {
    if len(s) > len(t) {
        s, t = t, s
    }
    if len(s) < len(t)-1 {
        return false
    }

    dp := make([][]int, len(s))
    for i := range s {
        dp[i] = make([]int, len(t))
    }

    for i := range t {
        if s[0] == t[i] {
            dp[0][i] = 0
        } else {
            dp[0][i] = 1
        }
    }

    for i := 1; i < len(s); i++ {
        if t[0] == s[i] {
            dp[i][0] = 0
        } else {
            dp[i][0] = 1
        }
    }
    // fmt.Println(s, t, dp)

    for i := 1; i < len(s); i++ {
        for j := 1; j < len(t); j++ {
            if s[i] == t[j] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = dp[i-1][j-1] + 1
                if dp[i][j] > dp[i-1][j]+1 {
                    dp[i][j] = dp[i-1][j] + 1
                }
                if dp[i][j] > dp[i][j-1]+1 {
                    dp[i][j] = dp[i][j-1] + 1
                }
                // if dp[i][j] > 1 {
                // fmt.Println(dp)
                //     return false
                // }
            }
        }
    }
    return dp[len(s)-1][len(t)-1] <= 1
}

func main() {
    s := []string{
        "abc",
        "a",
        "abc",
        "abc",
        "abcd",
    }
    t := []string{
        "ab",
        "abc",
        "acb",
        "adc",
        "aecf",
    }

    for i := 0; i < len(s); i++ {
        fmt.Println(s[i], t[i], isOneEditDistance(s[i], t[i]))
        fmt.Println(s[i], t[i], isOneEditDistance1(s[i], t[i]))
        fmt.Println(s[i], t[i], isOneEditDistance2(s[i], t[i]))
    }

}
