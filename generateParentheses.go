/**********************************************************************************
*
* Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
*
* For example, given n = 3, a solution set is:
*
* "((()))", "(()())", "(())()", "()(())", "()()()"
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val         int
    left, right *Node
}

func traverse(result string, n, right, left int) {
    sum := right - left
    if sum < 0 || n < 0 {
        return
    } else if n == 0 && sum == 0 {
        fmt.Print(result) //
        for i := 0; i < sum; i++ {
            fmt.Print(")")
        }
        fmt.Println()
        return
    }

    traverse(result+"(", n-1, right+1, left)

    traverse(result+")", n, right, left+1)
}

func main() {
    n := 3
    result := "("
    right := 1
    left := 0
    traverse(result, n-1, right, left)
}
