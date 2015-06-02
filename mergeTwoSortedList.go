/**********************************************************************************
*
* Merge two sorted linked lists and return it as a new list. The new list should be
* made by splicing together the nodes of the first two lists.
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

func mergeLinkedLists(l, r *LinkedNode) *LinkedNode {
    newl := &LinkedNode{
        0, nil,
    }
    i, j, k := (l), (r), newl
    for i != nil && j != nil {
        if i.val <= j.val {
            k.next = i
            i = i.next
        } else { // if i.val > j.val
            k.next = j
            j = j.next
        }
        k = k.next
    }
    for i != nil {
        k.next = i
        i, k = i.next, k.next
    }
    for j != nil {
        k.next = j
        j, k = j.next, k.next
    }
    return newl.next
}

func main() {
    l := LinkedNode{
        1, &LinkedNode{5, &LinkedNode{8, &LinkedNode{10, nil}}},
    }
    r := LinkedNode{
        3, &LinkedNode{12, &LinkedNode{18, &LinkedNode{30, nil}}},
    }

    newl := mergeLinkedLists(&l, &r)
    for p := newl; p != nil; p = p.next {
        fmt.Println(p.val)
    }
}
