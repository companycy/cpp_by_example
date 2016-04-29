/**********************************************************************************
*
* Determine whether an integer is a palindrome. Do this without extra space.
*
*
* Some hints:
*
* Could negative integers be palindromes? (ie, -1)
*
* If you are thinking of converting the integer to string, note the restriction of using extra space.
*
* You could also try reversing an integer. However, if you have solved the problem "Reverse Integer",
* you know that the reversed integer might overflow. How would you handle such case?
*
* There is a more generic way of solving this problem.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func isPalindromes(n int) bool {
    if n < 0 {
        return false
    }

    max, digits, p := 10, 1, n/10
    for ; (p / 10) != 0; p, max, digits = p/10, max*10, digits+1 {
    }
    if p != 0 {
        digits++
    }
    // fmt.Println(max, digits)

    mid, min := digits/2, 1
    var high, low, i int
    for high, low, i = (n/max)%10, (n/min)%10, 0; high == low && i < mid; {
        max, min, i = max/10, min*10, i+1
        high = (n / max) % 10
        low = (n / min) % 10
        // fmt.Println(high, low, max, min, i)
    }
    if high != low {
        return false
    } else {
        return true
    }
}

func main() {
    l := []int{
        0,
        -101,
        1001,
        1234321,
        2147447412,
        2142,
    }

    for i := range l {
        fmt.Println(l[i], isPalindromes(l[i]))
    }
}
