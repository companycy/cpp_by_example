/**********************************************************************************
*
* The n-queens puzzle is the problem of placing n queens on an n×n chessboard
* such that no two queens attack each other.
*
* Given an integer n, return all distinct solutions to the n-queens puzzle.
*
* Each solution contains a distinct board configuration of the n-queens' placement,
* where 'Q' and '.' both indicate a queen and an empty space respectively.
*
* For example,
* There exist two distinct solutions to the 4-queens puzzle:
*
* [
*  [".Q..",  // Solution 1
*   "...Q",
*   "Q...",
*   "..Q."],
*
*  ["..Q.",  // Solution 2
*   "Q...",
*   "...Q",
*   ".Q.."]
* ]
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func isValid(ret *[][]int, i, k int) bool {
    for m := 0; m < i; m++ {
        if (*ret)[m][k] == 1 {
            return false
        }
    }
    if 0 <= i-1 && i-1 < len(*ret) && 0 <= k-1 && k-1 < len(*ret) && (*ret)[i-1][k-1] == 1 {
        return false
    } else if 0 <= i-1 && i-1 < len(*ret) && k+1 < len(*ret) && (*ret)[i-1][k+1] == 1 {
        return false
    }
    return true
}

func getSolutions(ret [][]int, i, k int) {
    if i > len(ret) || k > len(ret) {
        return
    } else if (i != 0 || k != 0) && !isValid(&ret, i-1, k) {
        // fmt.Println("invalid", ret, i)
        return
    } else if i == len(ret) {
        if isValid(&ret, i-1, k) {
            for i := 0; i < len(ret); i++ {
                fmt.Println((ret)[i])
            }
        }
        fmt.Println()
        return
    }

    for j := 0; j < len(ret); j++ {
        if j == k {
            continue
        } else {
            ret[i][j] = 1
            getSolutions(ret, i+1, j)
            ret[i][j] = 0
        }
    }
}

func getAllSolutions(n int) {
    if n == 1 {
        fmt.Println(1)
    } else if n == 2 {
        return
    }

    ret := make([][]int, n)
    for i := 0; i < n; i++ {
        ret[i] = make([]int, n)
    }

    getSolutions(ret, 0, 0)
}

func main() {
    n := 4
    getAllSolutions(n)
}
