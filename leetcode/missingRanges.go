/**********************************************************************************
 *
 * Given a sorted integer array where the range of elements are [lower, upper] inclusive,
 * return its missing ranges.
 *
 * For example, given [0, 1, 3, 50, 75], lower = 0 and upper = 99,
 * return ["2", "4->49", "51->74", "76->99"].
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

type Range struct {
    begin, end int
}

func getMissingRanges(l *[]int, lower, upper int) *[]Range {
    if len(*l) == 0 {
        return &[]Range{
            Range{lower, upper},
        }
    }

    var ranges []Range
    for i := 0; i < len(*l)-1; i++ {
        var newrange Range
        if len(ranges) == 0 {
            newrange.begin = lower
            newrange.end = (*l)[i] - 1
        } else {
            newrange.begin = (*l)[i] + 1
            newrange.end = (*l)[i+1] - 1
        }
        ranges = append(ranges, newrange)
    }
    newrange := Range{
        (*l)[len(*l)-1] + 1, upper,
    }
    ranges = append(ranges, newrange)

    var result []Range
    for i := 0; i < len(ranges); i++ {
        if ranges[i].begin <= ranges[i].end {
            result = append(result, ranges[i])
        }
    }

    return &result
}

func main() {
    l := []int{
        0, 1, 3, 50, 75,
    }
    lower, upper := 0, 99
    fmt.Println(getMissingRanges(&l, lower, upper))
}
