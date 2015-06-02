/**********************************************************************************
*
* Given a string s, partition s such that every substring of the partition is a palindrome.
*
* Return all possible palindrome partitioning of s.
*
* For example, given s = "aab",
*
* Return
*
*   [
*     ["aa","b"],
*     ["a","a","b"]
*   ]
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

func partition(s string, result []string, results *[][]string) {
    if len(s) == 0 || len(s) == 1 {
        result = append(result, s)
        *results = append(*results, result)
        return
    } else if isPalindromes(s) {
        newresult := result
        newresult = append(newresult, s)
        *results = append(*results, newresult)
    }

    for i := range s {
        if isPalindromes(s[:i+1]) {
            newresult := append(result, (s[:i+1]))
            partition(s[i+1:], newresult, results)
        }
    }
}

func getAllPalindromes(s string) *[][]string {
    if len(s) == 0 || len(s) == 1 {
        results := make([][]string, 1)
        results[0] = []string{
            s,
        }
        return &results
    }

    var result []string
    var results [][]string
    partition(s, result, &results)
    return &results
}

func main() {
    s := []string{
        "aab",
        "aaba",
    }

    for i := range s {
        fmt.Println(getAllPalindromes(s[i]))

    }
}
