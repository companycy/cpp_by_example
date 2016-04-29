/**********************************************************************************
*
* Given a set of non-overlapping intervals, insert a new interval into the intervals (merge if necessary).
*
* You may assume that the intervals were initially sorted according to their start times.
*
* Example 1:
* Given intervals [1,3],[6,9], insert and merge [2,5] in as [1,5],[6,9].
*
* Example 2:
* Given [1,2],[3,5],[6,7],[8,10],[12,16], insert and merge [4,9] in as [1,2],[3,10],[12,16].
*
* This is because the new interval [4,9] overlaps with [3,5],[6,7],[8,10].
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Interval struct {
    begin, end int
}

func insertInterval(x Interval, l *[]Interval) *[]Interval {
    if len(*l) == 0 {
        *l = append(*l, x)
    }

    var newintervals []Interval
    for i := 0; i < len(*l); i++ {
        if (*l)[i].end < x.begin || (*l)[i].begin > x.end {
            newintervals = append(newintervals, (*l)[i])
        } else { // if (*l)[i].end >= x.begin
            newbegin := (*l)[i].begin
            newend := x.end
            for ; (*l)[i+1].end <= newend; i++ {
            }
            if (*l)[i+1].begin <= newend && i+1 < len(*l) {
                newend = (*l)[i+1].end
                i++
            }
            // else if (*l)[i+1].begin > newend {
            // }
            newintervals = append(newintervals, Interval{newbegin, newend})
            if i+1 < len(*l) {
                newintervals = append(newintervals, (*l)[i+1:]...)
            }
            break
        }
    }
    return &newintervals
}

func main() {
    l := []Interval{
        Interval{1, 3},
        Interval{6, 9},
    }
    l = []Interval{
        Interval{1, 2},
        Interval{3, 5},
        Interval{6, 7},
        Interval{8, 10},
        Interval{12, 16},
    }

    i := Interval{2, 5}
    i = Interval{4, 9}

    newl := insertInterval(i, &l)
    fmt.Println(*newl)
}
