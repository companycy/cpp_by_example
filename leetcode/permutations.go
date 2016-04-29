/**********************************************************************************
*
* Given a collection of numbers, return all possible permutations.
*
* For example,
* [1,2,3] have the following permutations:
* [1,2,3], [1,3,2], [2,1,3], [2,3,1], [3,1,2], and [3,2,1].
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getNextPermutation(l *[]int) {
    // fmt.Print(*l, "\t")
    i := len(*l) - 1
    for ; i >= 1 && (*l)[i-1] >= (*l)[i]; i-- {
    }

    j := i
    for ; j < len(*l) && i-1 >= 0 && (*l)[j] > (*l)[i-1]; j++ {
    }

    if i != 0 {
        t := (*l)[j-1]
        (*l)[j-1] = (*l)[i-1]
        (*l)[i-1] = t
    }

    for m, n := i, len(*l)-1; m < n; m, n = m+1, n-1 {
        t := (*l)[m]
        (*l)[m] = (*l)[n]
        (*l)[n] = t
    }
    // fmt.Println(*l)
}

func factorial(n int) int {
    if n == 0 {
        return 1
    } else if n == 1 {
        return 1
    }

    pp, p := 1, 1
    for i := 2; i <= n; i++ {
        pp = p
        p = pp * i
    }
    return p
}

func getAllPermutations2(l *[]int, result []int, results *[][]int) {
    if len(*l) == 0 {
        return
    } else if len(*l) == 1 {
        result := []int{
            (*l)[0],
        }
        *results = append(*results, result)
        return
    }

    // *results = append(*results, *l)
    // var newresult *[]int
    for i := 0; i < factorial(len(*l)); i++ {
        getNextPermutation(l)
        newl := make([]int, len(*l))
        copy(newl, *l)
        *results = append(*results, newl)
    }
}

func getAllPermutations(l *[]int, result []int, results *[][]int) {
    if len(*l) == 0 {
        return
    } else if len(*l) == 1 {
        result = append(result, (*l)[0])
        *results = append(*results, result)
        return
    }

    // m := make(map[int]int)
    for i := range *l {
        // if v, ok := m[(*l)[i]]; !ok {
        var newl []int
        newl = append(newl, (*l)[:i]...)
        newl = append(newl, (*l)[i+1:]...)
        result = append(result, (*l)[i])

        getAllPermutations(&newl, result, results)
        result = result[:len(result)-1]
        // fmt.Println(newl, result)
        //     m[(*l)[i]] = i
        // }
    }
}

func main() {
    l := &[]int{
        1, 2, 3,
    }

    var result []int
    var results [][]int
    // getAllPermutations(l, result, &results)
    getAllPermutations2(l, result, &results)
    fmt.Println(results)
}
