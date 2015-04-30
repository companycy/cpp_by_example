/**********************************************************************************
*
* A message containing letters from A-Z is being encoded to numbers using the following mapping:
*
* 'A' -> 1
* 'B' -> 2
* ...
* 'Z' -> 26
*
* Given an encoded message containing digits, determine the total number of ways to decode it.
*
* For example,
* Given encoded message "12",
* it could be decoded as "AB" (1 2) or "L" (12).
*
* The number of ways decoding "12" is 2.
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "strconv"
)

// assume cnt[i-1]
// then one more l[i], cnt[i] = cnt[i-1] // l[i-1] can't work with l[i], so l[i] is a single one
// or cnt[i] = cnt[i-1] + cnt[i-2] // l[i] can work with l[i-1], then it's cnt[i-2]
//    or l[i] can work as single one, then it's cnt[i-1]

func getDecodeWaysCnt(s string) int {
    pp, p := 1, 0
    if len(s) == 0 {
        return 0
    } else if len(s) == 1 {
        return 1
    }

    if preval, preok := strconv.Atoi(string(s[0])); preok == nil {
        if curval, curok := strconv.Atoi(string(s[1])); curok == nil {
            if curval == 0 {
                p = 1
            }
            val := preval *10 + curval
            if val > 26 {
                p = 1
            } else {
                p = 2
            }
        }
    } else {
        p = 0
    }

    if len(s) == 2 {
        return p
    }

    var cur int
    for i := 2; i < len(s); i++ {
        if preval, preok := strconv.Atoi(string(s[i-1])); preok == nil {
            if preval == 0 {
                cur = p
            } else {
                if curval, curok := strconv.Atoi(string(s[i])); curok == nil {
                    val := preval*10 + curval
                    if val > 26 {
                        cur = p
                    } else {
                        cur = pp + p
                        pp = p
                        p = cur
                    }          
                }
            }
        }
    }
    return cur
}

func main() {
    s := "12"
    s = "121"
    s = "611"
    s = "1211"
    cnt := getDecodeWaysCnt(s)
    fmt.Println(cnt)
}


