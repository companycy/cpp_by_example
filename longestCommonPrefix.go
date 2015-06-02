/**********************************************************************************
*
* Write a function to find the longest common prefix string amongst an array of strings.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getLongestCommonPrefix(l *[]string) string {
    if len(*l) == 0 {
        return ""
    } else if len(*l) == 1 {
        return (*l)[0]
    }

    for i := range (*l)[0] {
        for j := 1; j < len(*l); j++ {
            if len((*l)[j]) < i || (*l)[j][i] != (*l)[0][i] {
                return (*l)[0][:i]
            }
        }
    }
    return (*l)[0]
}

func main() {
    l := []string{
        // "abab", "aba", "abc",
        "abc", "ab", "ac",
    }
    fmt.Println(getLongestCommonPrefix(&l))
}
