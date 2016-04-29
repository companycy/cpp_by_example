/**********************************************************************************
*
* Given a digit string, return all possible letter combinations that the number could represent.
*
* A mapping of digit to letters (just like on the telephone buttons) is given below.
*
* Input:Digit string "23"
* Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
*
* Note:
* Although the above answer is in lexicographical order, your answer could be in any order you want.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

var results []string

func getAllCombinations(s, result string, m *map[int][]int) {
    // fmt.Println(s, result)
    if len(s) == 0 {
        results = append(results, result)
        return
    } else if len(s) == 1 {
        if v, ok := (*m)[int(s[0]-'0')]; ok {
            for i := range v {
                results = append(results, result+string(v[i]))
            }
        }
        return
    }

    k := int(s[0] - '0')
    if v, ok := (*m)[k]; ok {
        for i := range v {
            getAllCombinations(s[1:], result+string(v[i]), m)
        }
    } else {
        getAllCombinations(s[1:], result, m)
    }
}

func main() {
    s := "23"
    s = "123"
    s = "234"
    m := make(map[int][]int)
    for i, j := 2, int('a'); i < 8; i, j = i+1, j+2 {
        m[i] = []int{j + i - 2, j + i - 1, j + i}
    }
    j := int('p')
    m[7] = []int{
        j, j + 1, j + 2, j + 3,
    }
    m[8] = []int{
        j + 4, j + 5, j + 6,
    }
    m[9] = []int{
        j + 7, j + 8, j + 9, j + 10,
    }

    // for k, v := range m {
    //     fmt.Print(k, ":")
    //     for i := range v {
    //         fmt.Printf("%c\t", rune(v[i]))
    //     }
    //     fmt.Println()
    // }
    // fmt.Println()

    // var results []string
    var result string
    getAllCombinations(s, result, &m)
    for i := range results {
        fmt.Println(results[i])
    }
}
