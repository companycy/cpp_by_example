/**********************************************************************************
*
* Given two binary strings, return their sum (also a binary string).
*
* For example,
* a = "11"
* b = "1"
* Return "100".
*
*
**********************************************************************************/


package main

import (
    "fmt"
)

func main() {
    m := "11"
    l := "1"
    var result string
    n := 0
    if len(l) < len(m) {
        n = len(m) // bigger one
        for i := 0; i < n - len(l); i++ {
            l = "0" + l
        }
    } else {
        n = len(l)
        for i := 0; i < n - len(m); i++ {
            m = "0" + m
        }
    }
    next := 0
    a := int('0')
    for i := n-1; i >= 0; i-- {
        sum := int(l[i]) + int(m[i]) + next - 2 * a
        next = 0
        if sum > 1 {
            next = 1
        }
        if sum % 2 == 1 {
            result = "1" + result
        } else {
            result = "0" + result
        }
    }

    if next == 1 {
        result = "1" + result
    }
    fmt.Println(result)
}
