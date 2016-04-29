/**********************************************************************************
*
* Given two integers n and k, return all possible combinations of k numbers out of 1 ... n.
*
* For example,
* If n = 4 and k = 2, a solution is:
*
* [
*   [2,4],
*   [3,4],
*   [2,3],
*   [1,2],
*   [1,3],
*   [1,4],
* ]
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

var result [][]int

func printCombination(n, k int, solution *[]int) {
    // fmt.Println(n, k)
    if k == 0 {
        newsolution := make([]int, len(*solution))
        copy(newsolution, *solution)
        result = append(result, newsolution)
        // fmt.Println(*solution, result)
        return
    }
    for i := n; i >= 1; i-- {
        *solution = append(*solution, i)
        printCombination(i-1, k-1, solution)
        *solution = (*solution)[:len(*solution)-1]
    }
}

func printCombination2(s, n, k int, solution *[]int) {
    if k == 0 {
        newsolution := make([]int, len(*solution))
        copy(newsolution, *solution)
        result = append(result, newsolution)
        // fmt.Println(*solution, result)
        return
    }
    for i := s + 1; i <= n; i++ {
        *solution = append(*solution, i)
        printCombination2(i, n, k-1, solution)
        *solution = (*solution)[:len(*solution)-1]
    }
}

func main() {
    n, k := 4, 2
    var solution []int
    // printCombination(n, k, &solution)
    printCombination2(0, n, k, &solution)
    fmt.Println("")
    fmt.Println(result)
}
