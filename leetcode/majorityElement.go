/**********************************************************************************
 *
 * Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
 *
 * You may assume that the array is non-empty and the majority element always exist in the array.
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
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

func getMajorElement(l *[]int) int {
    if len(*l) == 0 {
        return -1
    } else if len(*l) == 1 {
        return (*l)[0]
    }

    var st Stack
    for i := range *l {
        if st.empty() {
            st.push((*l)[i])
        } else if val, ok := st.back(); ok && val != (*l)[i] {
            st.pop()
        }
    }
    if val, ok := st.back(); ok {
        return val
    } else {
        return -1
    }
}

func main() {
    l := []int{
        // 1, 2, 3, 5, 2, 2, 2,
        1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
    }
    fmt.Println(getMajorElement(&l))
}
