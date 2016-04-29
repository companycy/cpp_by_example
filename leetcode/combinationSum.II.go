/**********************************************************************************
*
* Given a collection of candidate numbers (C) and a target number (T), find all
* unique combinations in C where the candidate numbers sums to T.
*
* Each number in C may only be used once in the combination.
*
* Note:
*
* > All numbers (including target) will be positive integers.
* > Elements in a combination (a1, a2, … , ak) must be in non-descending order.
*   (ie, a1 ≤ a2 ≤ … ≤ ak).
* > The solution set must not contain duplicate combinations.
*
* For example, given candidate set 10,1,2,7,6,1,5 and target 8,
* A solution set is:
* [1, 7]
* [1, 2, 5]
* [2, 6]
* [1, 1, 6]
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

        // news := make([]int, len(*solution))
        // copy(news, *solution)
        // sort.Ints(news)
        // fmt.Println(news) // todo: duplicated cases
        // solutions = append(solutions, news)
        fmt.Println("one solution: ", *solution)

        *solution = (*solution)[:len(*solution)-1]
    }

    for i := 0; i < len(*l); i++ {
        *solution = append(*solution, (*l)[i])

        // todo: be careful!!
        newl := make([]int, len(*l))
        copy(newl, *l)
        newl = append(newl[:i], newl[i+1:]...)
        // fmt.Println((*l)[i], newl)

        newm := make(map[int]int, len(newl))
        for j := 0; j < len(newl); j++ {
            newm[newl[j]] = j
        }
        getSolution(&newl, &newm, solution, target-(*l)[i])

        *solution = (*solution)[:len(*solution)-1]
    }
}

func test(l *[]int) {
    for i := 0; i < len(*l); i++ {
        newl := make([]int, len(*l)-1)
        newl = append((*l)[:i], (*l)[i+1:]...)
        fmt.Println(newl)
    }
}

func main() {
    target := 8
    l := []int{
        10, 1, 2, 7, 6, 1, 5,
    }
    sort.Ints(l)

    m := make(map[int]int, len(l))
    for i := 0; i < len(l); i++ {
        m[l[i]] = i
    }

    var solution []int
    getSolution(&l, &m, &solution, target)
    fmt.Println(solutions)

    test(&l)
}
