/**********************************************************************************
*
* A linked list is given such that each node contains an additional random pointer
* which could point to any node in the list or null.
*
* Return a deep copy of the list.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type RandomNode struct {
    val              int
    next, randomnext *RandomNode
}

func copyRandomList(head *RandomNode) *RandomNode {
    if head == nil {
        return nil
    }

    for p := head; p != nil; {
        newnode := &RandomNode{
            p.val, p.next, p.randomnext,
        }
        p.next = newnode
        p = newnode.next
    }

    for p := head; p != nil; p = p.next.next {
        fmt.Print(p.val, "\t")
        if p.next != nil {
            fmt.Print(p.next.val, "\t")
        }
        if p.next.randomnext != nil {
            fmt.Print(p.next.randomnext.val, "\t")
        }
        if p.next.next != nil {
            fmt.Print(p.next.next.val, "\t")
        }
        fmt.Println()
    }

    newhead := head.next
    for p := head; p != nil; p = p.next.next {
        newnode := p.next
        if newnode == nil {
            break
        } else {
            if newnode.next != nil {
                p.next = newnode.next
                newnode.next = newnode.next.next
            }
            if newnode.randomnext != nil {
                newnode.randomnext = newnode.randomnext.next
            }
        }
    }

    return newhead
}

func printRandomList(head *RandomNode) {
    for p := head; p != nil; p = p.next {
        fmt.Print(p.val, "\t")
        if p.randomnext != nil {
            fmt.Print(p.randomnext.val, "\t")
        }
        if p.next != nil {
            fmt.Print(p.next.val, "\t")
        }
        fmt.Println()
    }
    fmt.Println()
}

func main() {
    node3 := &RandomNode{
        3, nil, nil,
    }
    node2 := &RandomNode{
        2, node3, nil,
    }
    node1 := &RandomNode{
        1, node2, node2,
    }
    node3.randomnext = node1
    node2.randomnext = node1
    printRandomList(node1)

    newhead := copyRandomList(node1)
    printRandomList(newhead)

}
