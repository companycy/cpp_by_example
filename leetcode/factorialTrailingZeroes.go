/**********************************************************************************
 *
 * Given an integer n, return the number of trailing zeroes in n!.
 *
 * Note: Your solution should be in polynomial time complexity.
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

func getZeroCnt(n int) int {
    var fivecnt int
    for i, j := 5.0, int(n/5.0); (j) >= 1; {
        fivecnt += (j)
        i = i * 5
        j = int(float64(n) / i)
    }
    return fivecnt
}

func getZeroCnt1(n int) int {
    var fivecnt, twocnt int
    for i := 1; i < n+1; i++ { // 1->n
        previ := i
        for j, k := i%2, i; j == 0 && k != 0; {
            twocnt++
            previ = k
            k = k / 2
            j = k % 2
        }
        for j, k := previ%5, previ; j == 0 && k != 0; {
            fivecnt++
            k = k / 5
            j = k % 5
        }
    }
    if fivecnt > twocnt {
        return twocnt
    } else {
        return fivecnt
    }
}

func main() {
    n := 5
    n = 10
    n = 15
    n = 25
    n = 26
    n = 125
    n = 4617
    zerocnt := getZeroCnt(n)
    // zerocnt = n / 5
    fmt.Println(zerocnt)
}
