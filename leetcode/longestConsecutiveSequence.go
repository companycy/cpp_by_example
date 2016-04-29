/**********************************************************************************
*
* Given an unsorted array of integers, find the length of the longest consecutive elements sequence.
*
* For example,
* Given [100, 4, 200, 1, 3, 2],
* The longest consecutive elements sequence is [1, 2, 3, 4]. Return its length: 4.
*
* Your algorithm should run in O(n) complexity.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getLongestConsecutiveLength(l *[]int) int {
    m := make(map[int]int)
    for i := range *l {
        m[(*l)[i]] = i
    }

    var maxlen int
    for i := 0; i < len(*l) && len(m) > 0; i++ {
        begin := (*l)[i] - 1
        for ; begin > (*l)[i]-len(*l); begin-- {
            if _, ok := m[begin]; !ok {
                begin++
                break
            } else {
                delete(m, begin)
            }
        }
        end := (*l)[i] + 1
        for ; end < (*l)[i]+len(*l); end++ {
            if _, ok := m[end]; !ok {
                end--
                break
            } else {
                delete(m, end)
            }
        }
        if maxlen < (end - begin) {
            maxlen = end - begin + 1
        }
        delete(m, (*l)[i])
    }
    return maxlen
}

func main() {
    l := []int{
        100, 4, 200, 1, 3, 2,
    }
    fmt.Println(getLongestConsecutiveLength(&l))
}
