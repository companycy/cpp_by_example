/**********************************************************************************
*
* Given an integer, convert it to a roman numeral.
*
* Input is guaranteed to be within the range from 1 to 3999.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func int2Rome(n int) string {
    symbols := []string{
        "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
    }
    values := []int{
        1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1,
    }

    var s string
    for i, left := 0, n; left != 0; i++ {
        for left >= values[i] {
            s = s + symbols[i]
            left = left - values[i]
        }
    }
    return s
}

func main() {
    n := 1234
    n = 3
    s := int2Rome(n)
    fmt.Println(s)
}
