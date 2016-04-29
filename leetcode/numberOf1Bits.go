/**********************************************************************************
*
* Write a function that takes an unsigned integer and returns the number of ’1' bits it has
* (also known as the Hamming weight).
*
* For example, the 32-bit integer ’11' has binary representation 00000000000000000000000000001011,
* so the function should return 3.
*
* Credits:Special thanks to @ts for adding this problem and creating all test cases.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getOneCnt(n uint) int {
    var cnt int
    for i := n; i != 0; i = i >> 1 {
        if (i & 1) == 0x01 {
            cnt++
        }
    }
    return cnt
}

func main() {
    n := []uint{
        11,
    }
    for i := range n {
        fmt.Println(getOneCnt(n[i]))
    }
}
