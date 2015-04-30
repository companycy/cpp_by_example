/**********************************************************************************
*
* Evaluate the value of an arithmetic expression in Reverse Polish Notation.
*
* Valid operators are +, -, *, /. Each operand may be an integer or another expression.
*
* Some examples:
*
*   ["2", "1", "+", "3", "*"] -> ((2 + 1) * 3) -> 9
*   ["4", "13", "5", "/", "+"] -> (4 + (13 / 5)) -> 6
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "strconv"
    "strings"
)

type Stack struct {
    l []string
}

func (s *Stack) push(val string) {
    s.l = append([]string{val}, s.l...)
}

func (s *Stack) pop() (string, bool) {
    if len(s.l) > 0 {
        val := s.l[0]
        s.l = s.l[1:]
        return val, true
    } else {
        return "", false
    }
}

func (s *Stack) empty() bool {
    return len(s.l) == 0
}

func (s *Stack) len() int {
    return len(s.l)
}

func doOperation(lh, rh, operator string) int {
    lhval, _ := strconv.Atoi(lh)
    rhval, _ := strconv.Atoi(rh)
    switch operator {
    case "+":
        return lhval + rhval
    case "-":
        return lhval - rhval
    case "*":
        return lhval * rhval
    case "/":
        return lhval / rhval
    default:
        return 0
    }
    return 0
}

func evaluate(l *[]string) (string, bool) {
    if len(*l) == 0 {
        return "", false
    }

    // operators := []string {
    //     "+", "-", "*", "/",
    // }
    operators := "+-*/"

    var st Stack
    for i := range *l {
        if strings.Index(operators, (*l)[i]) == -1 {
            st.push((*l)[i])
        } else {
            if st.len() >= 2 {
                rh, _ := st.pop()
                lh, _ := st.pop()
                result := doOperation(lh, rh, (*l)[i])
                st.push(strconv.Itoa(result))
            }
        }
    }
    return st.pop()
}

func main() {
    l := []string{
        // "2", "1", "+", "3", "*",
        "4", "13", "5", "/", "+",
    }

    result, _ := evaluate(&l)
    fmt.Println(result)
}
