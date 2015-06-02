/**********************************************************************************
*
* Given a string, find the length of the longest substring without repeating characters.
* For example, the longest substring without repeating letters for "abcabcbb" is "abc",
* which the length is 3. For "bbbbb" the longest substring is "b", with the length of 1.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getLenOfLongestSub1(s string) int {
    var maxlen, lastRepeatedPos int
    m := make(map[uint8]int)
    for i := range s {
        if pos, ok := m[s[i]]; ok && lastRepeatedPos < pos {
            lastRepeatedPos = pos
        }
        if maxlen < i-lastRepeatedPos {
            maxlen = i - lastRepeatedPos
        }
        m[s[i]] = i
    }
    return maxlen
}

func getLenOfLongestSub(s string) int {
    var maxlen int
    m := make(map[uint8]int)
    for i := range s {
        if pos, ok := m[s[i]]; ok {
            if maxlen < len(m) {
                maxlen = len(m)
            }
            for k, v := range m {
                if v < pos+1 {
                    delete(m, k)
                }
            }
        } else {
            // maxlen++
            m[s[i]] = i
        }
    }
    return maxlen
}

func main() {
    s := []string{
        "abcabcbb",
    }

    for i := range s {
        fmt.Println(getLenOfLongestSub1(s[i]))
        fmt.Println(getLenOfLongestSub(s[i]))
    }
}
