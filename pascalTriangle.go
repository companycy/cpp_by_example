/**********************************************************************************
*
* Given numRows, generate the first numRows of Pascal's triangle.
*
* For example, given numRows = 5,
* Return
*
* [
*      [1],
*     [1,1],
*    [1,2,1],
*   [1,3,3,1],
*  [1,4,6,4,1]
* ]
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getFirstNRows(n int, l *[]int, result *[][]int) {
    // fmt.Println(n, l, result)
    curlvl := []int{
        (*l)[0],
    }
    for i := 1; i < n; i++ {
        curlvl = append(curlvl, (*l)[i-1]+(*l)[i])
    }
    if n > 0 {
        curlvl = append(curlvl, (*l)[len(*l)-1])
    }

    (*result)[n] = curlvl
    if n < len(*result)-1 {
        getFirstNRows(n+1, &curlvl, result)
    }
}

func getFirstNRows2(n int, result *[][]int) {
    (*result)[0] = []int{1}

    if n == 1 {
        return
    }

    for i := 1; i < n; i++ {
        curlvl := []int{(*result)[i-1][0]}
        for j := 1; j < i; j++ {
            // result[i] =
            curlvl = append(curlvl, (*result)[i-1][j-1]+(*result)[i-1][j])
        }
        curlvl = append(curlvl, (*result)[i-1][len((*result)[i-1])-1])
        (*result)[i] = curlvl
    }
}

func main() {
    n := 5
    // l := []int{
    //     1,
    // }
    result := make([][]int, n)
    // getFirstNRows(0, &l, &result)
    getFirstNRows2(n, &result)
    fmt.Println(result)

}
