/**********************************************************************************
*
* Given a collection of numbers that might contain duplicates, return all possible unique permutations.
*
* For example,
* [1,1,2] have the following unique permutations:
* [1,1,2], [1,2,1], and [2,1,1].
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

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

func getNextPermutation(l *[]int) {
    i := len(*l) - 1
    for ; i > 0 && (*l)[i-1] >= (*l)[i]; i-- {
    }

    j := i
    for ; j < len(*l) && i-1 >= 0 && (*l)[j] >= (*l)[i-1]; j++ {
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
}

func getUniquePermutations(l *[]int, result []int, results *[][]int) {
    if len(*l) == 0 {
        return
    } else if len(*l) == 1 {
        result = append(result, (*l)[0])
        *results = append(*results, result)
        return
    }

    m := make(map[int]int)
    for i := range *l {
        if _, ok := m[(*l)[i]]; !ok {
            var newl []int
            newl = append(newl, (*l)[:i]...)
            newl = append(newl, (*l)[i+1:]...)
            result = append(result, (*l)[i])

            getUniquePermutations(&newl, result, results)
            result = result[:len(result)-1]
            // fmt.Println(newl, result)
            m[(*l)[i]] = i
        }
    }
}

func main() {
    l := &[]int{
        1, 1, 2,
    }
    var result []int
    var results [][]int
    getUniquePermutations(l, result, &results)
    fmt.Println(results)
}
