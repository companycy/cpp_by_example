/**********************************************************************************
*
* Divide two integers without using multiplication, division and mod operator.
*
* If it is overflow, return MAX_INT.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func doDivision(dividend, divisor int) int {
    // fmt.Println(dividend, divisor)
    if divisor == 0 {
        // int_max := 2147483647
        return 2147483647
    } else if dividend < divisor {
        return 0
    } else if dividend == 0 {
        return 0
    }

    result := 1
    var p, pr int
    for i := divisor; i <= dividend; i = i << 1 {
        pr = result
        result = result << 1
        p = i
    }
    // fmt.Println(p, result)
    return pr + doDivision(dividend-p, divisor)
}

func divide(dividend, divisor int) int {
    if divisor == 0 {
        // int_max := 2147483647
        return 2147483647
    } else if dividend == 0 {
        return 0
    } else if dividend > 0 && divisor > 0 && (dividend < divisor) {
        return 0
    }
    // fmt.Println(dividend, divisor)

    var positive bool
    if (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0) {
        positive = false
    } else if (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0) {
        positive = true
    }

    newdividend, newdivisor := dividend, divisor
    if dividend < 0 {
        newdividend = 0 - dividend
    }
    if divisor < 0 {
        newdivisor = 0 - divisor
    }

    // fmt.Println(newdividend, newdivisor)
    result := doDivision(newdividend, newdivisor)
    if positive == false {
        result = 0 - result
    }
    return result
}

func main() {
    dividend := []int{
        -2147483648,
        -2147483648,
        2147483647,
        2147483647,
        -2147483647,
        2147483647,
        -2147483647,
        2147483647,
        -1,
        1,
        -1,
        // 10,
        // 10,
        // 10,
        // 10,
        // 10,
        // 10,
        // 129, 0, 2,
    }
    divisor := []int{
        1,
        -1,
        10,
        2,
        -1,
        -1,
        1,
        1,
        -1,
        -1,
        1,
        // 11,
        // 10,
        // 7,
        // 5,
        // 3,
        // 2,
        // 7, 2, 0,
    }

    if len(dividend) == len(divisor) {
        for i := range dividend {
            result := divide(dividend[i], divisor[i])
            fmt.Println(dividend[i], divisor[i], result)
        }
    }
}
