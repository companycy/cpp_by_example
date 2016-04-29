/**********************************************************************************
*
* Given a linked list and a value x, partition it such that all nodes less than x come
* before nodes greater than or equal to x.
*
* You should preserve the original relative order of the nodes in each of the two partitions.
*
* For example,
* Given 1->4->3->2->5->2 and x = 3,
* return 1->2->2->4->3->5.
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

func swapLinkedNode(l, r *LinkedNode) {
    if l == nil || r == nil {
        return
    }

    val := l.val
    l.val = r.val
    r.val = val

    // val, next := l.val, l.next
    // l.val, l.next = r.val, r.next
    // r.val, r.next = val, next
}

func partition2(head *LinkedNode, pivot int) *LinkedNode {
    if head == nil {
        return nil
    }

    pnewhead, newhead := head, head.next
    for pp, p := head, head.next; p != nil; {
        if p.val < pivot {
            if p == newhead {
                pnewhead, newhead = newhead, newhead.next
                pp, p = p, p.next
            } else {
                pp.next = p.next

                p.next = newhead
                pnewhead.next = p

                p = pp.next
                pnewhead = pnewhead.next
            }
        } else {
            pp, p = p, p.next
        }
    }

    return newhead
}

func partition(head *LinkedNode, pivot int) *LinkedNode {
    if head == nil {
        return nil
    }

    var tail, pivotPos *LinkedNode
    for tail = head; tail != nil && tail.next != nil; tail = tail.next {
        if tail.val == pivot {
            pivotPos = tail
        }
    }

    swapLinkedNode(tail, pivotPos)
    newp := head
    for p := head; p != tail; p = p.next {
        if p.val <= pivot {
            swapLinkedNode(newp, p)
            newp = newp.next
        }
    }
    swapLinkedNode(newp, tail)
    return newp
}

func main() {
    head := &LinkedNode{
        1, &LinkedNode{
            4, &LinkedNode{
                3, &LinkedNode{
                    2, &LinkedNode{
                        5, &LinkedNode{
                            2, nil,
                        },
                    },
                },
            },
        },
    }
    pivot := 3

    newhead := partition2(head, pivot)
    for p := newhead; p != nil; p = p.next {
        fmt.Print(p.val, "\t")
    }
    fmt.Println()

    for p := head; p != nil; p = p.next {
        fmt.Print(p.val, "\t")
    }
    fmt.Println()
}
