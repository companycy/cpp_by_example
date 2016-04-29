/**********************************************************************************
*
* Given a string containing just the characters '(' and ')',
* find the length of the longest valid (well-formed) parentheses substring.
*
* For "(()", the longest valid parentheses substring is "()", which has length = 2.
*
* Another example is ")()())", where the longest valid parentheses substring is "()()",
* which has length = 4.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Stack struct {
    l []uint8
}

func (s *Stack) push(val uint8) {
    s.l = append([]uint8{val}, s.l...)
}

func (s *Stack) pop() (uint8, bool) {
    if len(s.l) > 0 {
        val := s.l[0]
        s.l = s.l[1:]
        return val, true
    } else {
        return 0, false
    }
}

func (s *Stack) back() (uint8, bool) {
    if len(s.l) > 0 {
        return s.l[0], true
    } else {
        return '0', false
    }
}

func (s *Stack) empty() bool {
    return len(s.l) == 0
}

func getLenOfLongestValidParenthese(s string) int {
    if len(s) < 2 {
        return 0
    } else if len(s) == 2 {
        if s[0] == '(' && s[1] == ')' {
            return 2
        } else {
            return 0
        }
    }

    var maxlen int
    var st Stack
    for i := range s {
        if st.empty() {
            st.push(s[i])
        } else {
            if back, ok := st.back(); ok && back == '(' && s[i] == ')' {
                maxlen = maxlen + 2
                st.pop()
            } else {
                st.push(s[i])
            }
        }
    }
    return maxlen
}

func main() {
    l := []string{
        "(()",
        ")()())",
    }
    for i := range l {
        fmt.Println(getLenOfLongestValidParenthese(l[i]))
    }
}
