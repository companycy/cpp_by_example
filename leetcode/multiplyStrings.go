/**********************************************************************************
*
* Given two numbers represented as strings, return multiplication of the numbers as a string.
*
* Note: The numbers can be arbitrarily large and are non-negative.
*
**********************************************************************************/

package main

import (
    "fmt"
    "strconv"
)

func getMultiplication(l, r string) string {
    if len(l) == 0 {
        return r
    } else if len(r) == 0 {
        return l
    } else if i, err := strconv.Atoi(l); err == nil && i == 0 {
        return "0"
    } else if i, err := strconv.Atoi(r); err == nil && i == 0 {
        return "0"
    } else if len(l) > len(r) {
        l, r = r, l
    }

    var result []string
    var maxlen int
    for i := len(l) - 1; i >= 0; i-- {
        if string(l[i]) == "0" {
            continue
        }

        var newrest string
        var adv int
        left, _ := strconv.Atoi(string(l[i]))
        for j := len(r) - 1; j >= 0; j-- {
            right, _ := strconv.Atoi(string(r[j]))
            ret := left*right + adv
            adv = ret / 10
            newrest = strconv.Itoa(ret%10) + newrest
        }
        if adv != 0 {
            newrest = strconv.Itoa(adv) + newrest
        }
        for j := 0; j < len(l)-1-i; j++ {
            newrest += "0"
        }
        if maxlen < len(newrest) {
            maxlen = len(newrest)
        }
        result = append(result, newrest)
    }
    // fmt.Println(result)

    var ret string
    var adv int
    for i := 0; i < maxlen; i++ {
        var acc int
        for j := range result {
            if len(result[j]) >= i+1 {
                s, _ := strconv.Atoi(string(result[j][len(result[j])-i-1]))
                acc += s + adv
                adv = acc / 10
                acc = acc % 10
            }
        }
        ret = strconv.Itoa(acc) + ret
    }
    return ret
}

func main() {
    l := "201"
    r := "251"
    fmt.Println(getMultiplication(l, r))
}
