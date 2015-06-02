/**********************************************************************************
*
* Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.
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

func getHeads(ls *[]LinkedNode) *[]*LinkedNode {
    var heads []*LinkedNode
    for i := range *ls {
        heads = append(heads, &(*ls)[i])
    }
    return &heads
}

func hasNil(heads *[]*LinkedNode) (bool, int) {
    for i := range *heads {
        if (*heads)[i] == nil {
            return true, i
        }
    }
    return false, -1
}

func notAllNil(heads *[]*LinkedNode) bool {
    for hasnil, i := hasNil(heads); hasnil; hasnil, i = hasNil(heads) {
        *heads = append((*heads)[:i], (*heads)[i+1:]...)
    }
    if len(*heads) > 0 {
        return true
    } else {
        return false
    }
}

func getCurNodeAndNext(heads *[]*LinkedNode) *LinkedNode {
    curNode, min := (*heads)[0], (*heads)[0].val
    for i := range *heads {
        if min > (*heads)[i].val {
            min = (*heads)[i].val
            curNode = (*heads)[i]
        }
    }
    for i := range *heads {
        if (*heads)[i] == curNode {
            (*heads)[i] = curNode.next
            break
        }
    }
    return curNode
}

func mergeLinkedLists(ls *[]LinkedNode) *LinkedNode {
    head := LinkedNode{
        0, nil,
    }
    heads := getHeads(ls)
    for p := &head; notAllNil(heads); {
        curNode := getCurNodeAndNext(heads)
        if p != nil {
            p.next = curNode
            p = curNode
        }
    }
    return head.next
}

func main() {
    ls := []LinkedNode{
        LinkedNode{
            1, &LinkedNode{5, &LinkedNode{8, &LinkedNode{10, nil}}},
        },
        LinkedNode{
            3, &LinkedNode{12, &LinkedNode{18, &LinkedNode{30, nil}}},
        },
        LinkedNode{
            7, &LinkedNode{15, &LinkedNode{20, &LinkedNode{40, nil}}},
        },
    }

    ret := mergeLinkedLists(&ls)

    fmt.Print("result:\t")
    for ; ret != nil; ret = ret.next {
        fmt.Print(ret.val, "\t")
    }

}
