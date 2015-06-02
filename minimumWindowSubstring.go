/**********************************************************************************
*
* Given a string S and a string T, find the minimum window in S which will
* contain all the characters in T in complexity O(n).
*
* For example,
* S = "ADOBECODEBANC"
* T = "ABC"
*
* Minimum window is "BANC".
*
* Note:
*
* > If there is no such window in S that covers all characters in T,
*   return the emtpy string "".
*
* > If there are multiple such windows, you are guaranteed that there
*   will always be only one unique minimum window in S.
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "strings"
)

func getMinWindowSubstr2(l, r string) string {
    m := make(map[uint8][]int)
    for i := range l {
        if v, ok := m[l[i]]; ok {
            v = append(i)
        } else {
            m[l[i]] = []int{i}
        }
    }
    return minWin
}

func getMinWindowSubstr(l, r string) string {
    minWin := l
    for i := 0; i < len(l); i++ {
        newl := l[i:]
        j := 0
        for j < len(r) {
            if pos := strings.Index(newl, string(r[j])); pos != -1 {
                j++
            } else {
                break
            }
        }
        if i == 0 && j != len(r) {
            return ""
        } else if j == len(r) && len(minWin) > len(newl) {
            minWin = newl
        }
    }
    return minWin
}

func main() {
    l := "ADOBECODEBANC"
    r := "ABC"
    fmt.Println(getMinWindowSubstr(l, r))
}
