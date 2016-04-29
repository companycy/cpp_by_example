/**********************************************************************************
*
* Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
* Find all unique triplets in the array which gives the sum of zero.
*
* Note:
*
* Elements in a triplet (a,b,c) must be in non-descending order. (ie, a ≤ b ≤ c)
* The solution set must not contain duplicate triplets.
*
*     For example, given array S = {-1 0 1 2 -1 -4},
*
*     A solution set is:
*     (-1, 0, 1)
*     (-1, -1, 2)
*
*
**********************************************************************************/

package main
import (
    "fmt"
)

func main() {
    l := []int {-1, 0, 1, 2, -1, -4}
    m := make(map[int]int, len(l))
    for i := 0; i < len(l); i++ {
        m[l[i]] = i
    }
    sum := 0
    for i := 0; i < len(l); i++ {
        sum1 := sum - l[i]
        for j := i+1; j < len(l); j++ {
            sum2 := sum1 - l[j]
            if v, ok := m[sum2]; ok && v > j {
                fmt.Println(l[i], l[j], sum2)
            }
        }
    }
}

