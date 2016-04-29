/**********************************************************************************
*
* Given a collection of intervals, merge all overlapping intervals.
*
* For example,
* Given [1,3],[2,6],[8,10],[15,18],
* return [1,6],[8,10],[15,18].
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "sort"
)

type Interval struct {
    begin, end int
}

type ByInterval []Interval

func (l ByInterval) Len() int {
    return len(l)
}

func (l ByInterval) Swap(i, j int) {
    t := (l)[i]
    (l)[i] = (l)[j]
    (l)[j] = t
}

func (l ByInterval) Less(i, j int) bool {
    return (l)[i].begin < (l)[j].begin
}

func mergeOverlappingIntervals(l *ByInterval) *ByInterval {
    if len(*l) == 0 {
        return nil
    } else if len(*l) == 1 {
        return l
    }

    // sort by begin
    sort.Sort(*l)
    fmt.Println(*l)

    newl := ByInterval{}
    for i := 0; i < len(*l); i++ {
        newinterval := (*l)[i]
        for ; i+1 < len(*l) && (*l)[i].end >= (*l)[i+1].begin; i++ {
        }
        newinterval.end = (*l)[i].end
        newl = append(newl, newinterval)
    }
    return &newl
}

func main() {
    l := ByInterval{
        {1, 3}, {2, 6}, {8, 10}, {15, 18},
    }
    fmt.Println(*mergeOverlappingIntervals(&l))
}
