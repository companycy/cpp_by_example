/**********************************************************************************
*
* Given a linked list, determine if it has a cycle in it.
*
* Follow up:
* Can you solve it without using extra space?
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type LinkedNode struct {
    val  int
    next *LinkedNode
}

func hasCycle(head *LinkedNode) bool {
    if head == nil || head.next == nil || head.next.next == nil {
        return false
    }

    pp, p := head, head.next
    for ; pp != nil && p != nil && p.next != nil; pp, p = pp.next, p.next.next {
        if pp == p {
            return true
        }
    }
    return false
}

func main() {
    head := LinkedNode{
        1, nil,
    }
    fmt.Println(hasCycle(head))
}
