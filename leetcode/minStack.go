/**********************************************************************************
*
* Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
*
* push(x) -- Push element x onto stack.
*
* pop() -- Removes the element on top of the stack.
*
* top() -- Get the top element.
*
* getMin() -- Retrieve the minimum element in the stack.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val        int
    prev, next *Node
}

type Stack struct {
    l    []*Node
    head []*Node
}

func (st *Stack) push(val int) {
    newnode := Node{
        val, nil, nil,
    }
    st.l = append([]*Node{&newnode}, st.l...)
    // for i := range st.l {
    //     fmt.Print(st.l[i].val, "\t")
    // }
    // fmt.Println()

    i := 0
    for ; i < len(st.head) && val >= st.head[i].val; i++ {
    }

    if i == len(st.head) {
        st.head = append(st.head, &newnode)
    } else {
        t := make([]*Node, i)
        copy(t, st.head[:i])
        t = append(t, &newnode)
        st.head = append(t, st.head[i:]...)
    }
    // for i := range st.head {
    //     fmt.Print(st.head[i].val, "\t")
    // }
    // fmt.Println()
}

func (st *Stack) pop() (int, bool) {
    if len(st.l) == 0 {
        return -1, false
    } else {
        top := st.l[0]
        st.l = st.l[1:]

        i := 0
        for ; i < len(st.head) && top != st.head[i]; i++ {
        }
        if i < len(st.head) {
            st.head = append(st.head[:i], st.head[i+1:]...)
        }

        return top.val, true
    }
}

func (st *Stack) top() (int, bool) {
    if len(st.l) == 0 {
        return -1, false
    } else {
        return st.l[0].val, true
    }
}

func (st *Stack) getMin() (int, bool) {
    if len(st.l) == 0 {
        return -1, false
    } else {
        return st.head[0].val, true
    }
}

func main() {
    var st Stack
    st.push(2)
    st.push(0)
    st.push(3)
    st.push(0)

    fmt.Println(st.getMin()) // 0
    fmt.Println(st.pop())    // 0

    fmt.Println(st.getMin()) // 0
    fmt.Println(st.pop())    // 3

    fmt.Println(st.getMin()) // 0
    fmt.Println(st.pop())    // 0

    fmt.Println(st.getMin()) // 2
    fmt.Println(st.pop())    // 2

    st.push(2147483646)
    st.push(2147483646)
    st.push(2147483647)

    fmt.Println(st.top())
    fmt.Println(st.pop())

    fmt.Println(st.getMin())
    fmt.Println(st.pop())

    fmt.Println(st.getMin())
    fmt.Println(st.pop())

    st.push(2147483647)
    fmt.Println(st.top())
    fmt.Println(st.getMin())

    st.push(-2147483648)
    fmt.Println(st.top())
    fmt.Println(st.getMin())
    fmt.Println(st.pop())
    fmt.Println(st.getMin())
}
