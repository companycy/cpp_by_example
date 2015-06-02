/**********************************************************************************
*
* Given n non-negative integers representing the histogram's bar height where the width of each bar is 1,
* find the area of largest rectangle in the histogram.
*
* Above is a histogram where width of each bar is 1, given height = [2,1,5,6,2,3].
*
* The largest rectangle is shown in the shaded area, which has area = 10 unit.
*
* For example,
* Given height = [2,1,5,6,2,3],
* return 10.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Stack struct {
    l []int
}

func (s *Stack) push(val int) {
    s.l = append([]int{val}, s.l...)
}

func (s *Stack) pop() (int, bool) {
    if len(s.l) > 0 {
        val := s.l[0]
        s.l = s.l[1:]
        return val, true
    } else {
        return 0, false
    }
}

func (s *Stack) back() (int, bool) {
    if len(s.l) > 0 {
        return s.l[0], true
    } else {
        return 0, false
    }
}

func (s *Stack) empty() bool {
    return len(s.l) == 0
}

func getMaxArea2(l []int) int {
    l = append(l, 0)

    var s Stack
    var result int
    for i := range l {
        if v, ok := s.back(); s.empty() || (ok && l[i] >= v) {
            s.push(i)
        } else {
            if j, ok := s.pop(); ok {
                var newresult int
                if s.empty() {
                    newresult = l[j] * i
                } else {
                    newresult = l[j] * (i - j)
                }
                if result < newresult {
                    result = newresult
                }
            }
            i--
        }
    }
    return result
}

func getMaxArea(l *[]int) int {
    if len(*l) == 0 {
        return 0
    } else if len(*l) == 1 {
        return (*l)[0]
    }

    result := (*l)[0]
    min := result
    for i := range *l {
        if min > (*l)[i] {
            min = (*l)[i]
        }

        if result < min*(i+1) {
            result = min * (i + 1)
        }
        // if result < (*l)[i] {
        //     result = (*l)[i]
        // }
        for j, newmin := i, (*l)[i]; (*l)[j] > min; j-- {
            if newmin > (*l)[j] {
                newmin = (*l)[j]
            }

            if result < newmin*(i-j+1) {
                result = newmin * (i - j + 1)
            }
        }
    }
    return result
}

func main() {
    l := []int{
        // 2, 1, 5, 6, 2, 3,
        // 1, 2, 3, 4,
        // 2, 1, 2,
        4, 2, 0, 3, 2, 5,
    }
    max := getMaxArea(&l)
    fmt.Println(max)
}
