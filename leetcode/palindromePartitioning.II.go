/**********************************************************************************
*
* Given a string s, partition s such that every substring of the partition is a palindrome.
*
* Return the minimum cuts needed for a palindrome partitioning of s.
*
* For example, given s = "aab",
* Return 1 since the palindrome partitioning ["aa","b"] could be produced using 1 cut.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func isPalindromes(s string) bool {
    if len(s) == 0 {
        return false
    } else if len(s) == 1 {
        return true
    }

    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

// results *[][]string
func partition(s string, result []string, min *int) {
    if len(s) == 0 || len(s) == 1 {
        result = append(result, s)
        // fmt.Println(result, *min)
        if *min > len(result)-1 {
            *min = len(result) - 1
        }
        // *results = append(*results, result)
        return
    } else if isPalindromes(s) {
        newresult := result
        newresult = append(newresult, s)
        // fmt.Println(newresult, *min)
        if *min > len(newresult)-1 {
            *min = len(newresult) - 1
        }
        return
        // *results = append(*results, newresult)
    }

    for i := range s {
        if isPalindromes(s[:i+1]) {
            newresult := append(result, (s[:i+1]))
            partition(s[i+1:], newresult, min)
            // partition(s[i+1:], newresult, results)
        }
    }
}

func getMinCut(s string) int {
    if isPalindromes(s) {
        return 0
    }

    minCut := len(s)
    var result []string
    partition(s, result, &minCut)
    return minCut
}

func getMinCut2(s string) int {
    if isPalindromes(s) {
        return 0
    }

    palindromeMatrix := make([][]int, len(s))
    for i := range palindromeMatrix {
        palindromeMatrix[i] = make([]int, len(s))
    }

    for i := len(s) - 1; i >= 0; i-- {
        for j := i; j < len(s); j++ {
            if i == j {
                palindromeMatrix[i][j] = 1
            } else if j == i+1 && s[i] == s[j] {
                palindromeMatrix[i][j] = 1
            } else if s[i] == s[j] && palindromeMatrix[i+1][j-1] == 1 {
                palindromeMatrix[i][j] = 1
            }
        }
    }
    // fmt.Println(s, palindromeMatrix)

    minCut := make([]int, len(s))
    for i := 0; i < len(s); i++ {
        if palindromeMatrix[0][i] == 1 {
            minCut[i] = 0
        } else {
            minCut[i] = minCut[i-1] + 1
            for j := 0; j < i; j++ {
                if palindromeMatrix[j+1][i] == 1 && minCut[i] > minCut[j]+1 {
                    minCut[i] = minCut[j] + 1
                }
            }
        }
    }
    return minCut[len(s)-1]
}

func main() {
    s := []string{
        "aab",
        "aaba",
        "ababababababababababababcbabababababababababababa",
        "fifgbeajcacehiicccfecbfhhgfiiecdcjjffbghdidbhbdbfbfjccgbbdcjheccfbhafehieabbdfeigbiaggchaeghaijfbjhi",
        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
    }

    for i := range s {
        // fmt.Println(getMinCut(s[i]))
        fmt.Println(getMinCut2(s[i]))
    }
}
