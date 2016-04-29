/**********************************************************************************
*
* Given a linked list, return the node where the cycle begins. If there is no cycle, return null.
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

// refer to
// http://www.cnblogs.com/hiddenfox/p/3408931.html
// when detect if there is cycle, there is equation:
//
func hasCycle(head *LinkedNode) (*LinkedNode, bool) {
    if head == nil || head.next == nil || head.next.next == nil {
        return false
    }

    var len int

    pp, p := head, head.next
    for ; pp != nil && p != nil && p.next != nil; pp, p = pp.next, p.next.next {
        len++
        if pp == p {
            fmt.Println(len) // length of cycle
            return p, true
        }
    }
    return nil, false
}

func getCycleStartNode(head, p *LinkedNode) *LinkedNode {
    for pp := head; p != pp; p, pp = p.next, pp.next {
    }
    return p
}

func main() {
    head := LinkedNode{
        1, nil,
    }

    if p, ok := hasCycle(head); ok {
        getCycleStartNode(head, p)
    }
}
