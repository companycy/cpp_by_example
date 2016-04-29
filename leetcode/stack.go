
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
        return -1, false
    }
}

func main() {
    var st Stack
    st.push(1)
    fmt.Println(st.pop())

    st.push(3)
    fmt.Println(st.pop())
    fmt.Println(st.pop())

}
