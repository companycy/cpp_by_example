/**********************************************************************************
*
* Given a string S, find the longest palindromic substring in S.
* You may assume that the maximum length of S is 1000,
* and there exists one unique longest palindromic substring.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getLongestPalindromincSubstr(s string) string {
    if len(s) == 0 || len(s) == 1 {
        return s
    } else if len(s) == 2 {
        if s[0] == s[1] {
            return s
        } else {
            return s[1:] // or s[1]
        }
    }

    newbegin, newend, begin, end := 0, 0, 0, 0
    for i := 1; i < len(s); i++ {
        if s[i] == s[i-1] { // && end == i-1
            repeated := true
            for j := begin; j < end; j++ {
                if s[j+1] != s[j] {
                    repeated = false
                    break
                }
            }
            if repeated {
                end = end + 1 // i
                if newend-newbegin+1 < end-begin+1 {
                    newbegin, newend = begin, end
                }
            } else {
                begin, end = i-1, i
            }
        } else { // if s[i] != s[i-1] && end == i-1
            if begin-1 >= 0 && s[begin-1] == s[i] {
                begin = begin - 1
                end = end + 1

                if newend-newbegin < end-begin {
                    newbegin, newend = begin, end
                }
                // fmt.Println(begin, end, s[begin:end+1])
            } else {
                begin, end = i, i
            }
        }
    }

    return s[newbegin : newend+1]
}

func getLongestPalindromincSubstr2(low, high int, s string) string {
    for ; low >= 0 && high < len(s); low, high = low-1, high+1 {
        if s[low] != s[high] {
            return s[low+1 : high]
        }
    }

    return s[low+1 : high]
}

func getLongestPalindromincSubstr1(s string) string {
    var result string
    for i := range s {
        // fmt.Println(i)
        news := getLongestPalindromincSubstr2(i, i, s)
        // fmt.Println(news)
        if len(result) < len(news) {
            result = news
        }
        news = getLongestPalindromincSubstr2(i, i+1, s)
        // fmt.Println(news)
        if len(result) < len(news) {
            result = news
        }
    }
    return result
}

func getLongestPalindromincSubstr3(s string) string {
    dp := make([][]bool, len(s))
    for i := range dp {
        dp[i] = make([]bool, len(s))
    }

    var result string
    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < len(s); j++ {
            if (j-i <= 2 && s[i] == s[j]) || (s[i] == s[j] && dp[i+1][j-1]) {
                dp[i][j] = true
                if j-i+1 > len(result) {
                    result = s[i : j+1]
                }
            }
        }
    }
    return result
}

func main() {
    s := []string{
        "abccba", "abcba", "fabbac",
        "fffaaaf",
    }
    for i := range s {
        fmt.Println(getLongestPalindromincSubstr(s[i]))
        fmt.Println(getLongestPalindromincSubstr1(s[i]))
        fmt.Println(getLongestPalindromincSubstr3(s[i]))
    }
}
