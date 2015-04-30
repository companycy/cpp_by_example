/**********************************************************************************
*
* Given a set of candidate numbers (C) and a target number (T), find all unique combinations
* in C where the candidate numbers sums to T.
*
* The same repeated number may be chosen from C unlimited number of times.
*
* Note:
*
* All numbers (including target) will be positive integers.
* Elements in a combination (a1, a2, … , ak) must be in non-descending order. (ie, a1 ≤ a2 ≤ … ≤ ak).
* The solution set must not contain duplicate combinations.
*
* For example, given candidate set 2,3,6,7 and target 7,
* A solution set is:
* [7]
* [2, 2, 3]
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "sort"
)

var solutions [][]int

func getSolution(l *[]int, m *map[int]int, solution *[]int, target int) {
    if target < (*l)[0] {
        return
    }

    if _, ok := (*m)[target]; ok {
        *solution = append(*solution, target)

        news := make([]int, len(*solution))
        copy(news, *solution)
        // sort.Ints(news)
        fmt.Println(news) // todo: duplicated cases
        solutions = append(solutions, news)

        *solution = (*solution)[:len(*solution)-1]
    }

    for i := 0; i < len(*l); i++ {
        *solution = append(*solution, (*l)[i])
        getSolution(l, m, solution, target-(*l)[i])
        *solution = (*solution)[:len(*solution)-1]
    }
}

func main() {
    target := 8
    l := []int{
        2, 3, 6, 7,
        // 10, 1, 2, 7, 6, 1, 5,
    }
    sort.Ints(l)

    m := make(map[int]int, len(l))
    for i := 0; i < len(l); i++ {
        m[l[i]] = i
    }

    var solution []int
    getSolution(&l, &m, &solution, target)

    fmt.Println(solutions)
}
