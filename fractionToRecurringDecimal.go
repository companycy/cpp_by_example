/**********************************************************************************
*
* Given two integers representing the numerator and denominator of a fraction, return the fraction in string format.
*
* If the fractional part is repeating, enclose the repeating part in parentheses.
*
* For example,
*
* Given numerator = 1, denominator = 2, return "0.5".
* Given numerator = 2, denominator = 1, return "2".
* Given numerator = 2, denominator = 3, return "0.(6)".
*
* Credits:Special thanks to @Shangrila for adding this problem and creating all test cases.
*
**********************************************************************************/

package main

import (
    "fmt"
    "strconv"
)

type Fraction struct {
    num, deno int
}

func getResult(num, deno int) string {
    if deno == 0 {
        return ""
    } else if num == 0 {
        return "0"
    }

    newnum := num
    if newnum < 0 {
        newnum = -newnum
    }
    newdeno := deno
    if newdeno < 0 {
        newdeno = -newdeno
    }

    division := newnum / newdeno
    result := strconv.Itoa(division)

    if (num > 0 && deno < 0) || (num < 0 && deno > 0) {
        result = "-" + result
    }

    remainder := newnum % newdeno
    if remainder == 0 {
        return result
    } else {
        result = result + "."
    }

    // detect repeatating part by remainder
    remaindermap := make(map[int]int)
    for remainder%newdeno != 0 {
        k := remainder % newdeno
        if pos, ok := remaindermap[k]; ok {
            newresult := result[:pos]
            newresult = newresult + "(" + result[pos:len(result)] + ")"
            return newresult
        }
        remaindermap[k] = len(result)
        remainder = k * 10
        result += strconv.Itoa(remainder / newdeno)
    }
    return result
}

func main() {
    l := []Fraction{
        Fraction{1, 2},
        Fraction{10, 2},
        Fraction{100, 2},
        Fraction{1, 3},
        Fraction{100, 3},
        Fraction{1, 6},
        Fraction{100, 6},
        Fraction{-1, 4},
        Fraction{1, -3},
        Fraction{-1, -6},
        Fraction{25, 99},
        Fraction{1, 7},
        Fraction{10, 7},
        Fraction{100, 7},
        Fraction{1, 17},
        Fraction{1, 1024},
        Fraction{-2147483648, -1999},
        Fraction{-1, -2147483648},
    }

    for i := range l {
        s := getResult(l[i].num, l[i].deno)
        fmt.Println(s)
    }
}
