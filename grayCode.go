/**********************************************************************************
*
* The gray code is a binary numeral system where two successive values differ in only one bit.
*
* Given a non-negative integer n representing the total number of bits in the code,
* print the sequence of gray code. A gray code sequence must begin with 0.
*
* For example, given n = 2, return [0,1,3,2]. Its gray code sequence is:
*
* 00 - 0
* 01 - 1
* 11 - 3
* 10 - 2
*
* Note:
* For a given n, a gray code sequence is not uniquely defined.
*
* For example, [0,2,3,1] is also a valid gray code sequence according to the above definition.
*
* For now, the judge is able to judge based on one instance of gray code sequence. Sorry about that.
*
**********************************************************************************/

package main

import (
    "fmt"
    // "sort"
)

func getGrayCode1(n int, results *[]string) {
    if n == 0 {
        return
    } else if n == 1 {
        *results = append(*results, "0", "1")
        return
    }

    var newresults []string
    getGrayCode1(n-1, &newresults)

    for i := 0; i < len(newresults); i++ {
        *results = append(*results, "0"+newresults[i])
    }

    // Reverse(newresults)
    // for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
    //     s[i], s[j] = s[j], s[i]
    // }

    for i := n - 1; i >= 0; i-- {
        *results = append(*results, "1"+newresults[i])
    }
}

func getGrayCode(n int, results *[]int) {
    if n <= 0 {
        return
    } else if n == 1 {
        *results = append(*results, 0, 1)
        return
    }

    var newresults []int
    getGrayCode(n-1, &newresults)

    for i := 0; i < len(newresults); i++ {
        *results = append(*results, newresults[i])
    }

    bits := uint32(n-1)
    for i := len(newresults)-1; i >= 0; i-- {
        *results = append(*results, 1<<bits+newresults[i])
    }
}

func main() {
    n := 2
    n = 3

    // var results []string
    // getGrayCode(n, &results)
    // fmt.Println(results)


    var results []int
    getGrayCode(n, &results)
    fmt.Println(results)
}
