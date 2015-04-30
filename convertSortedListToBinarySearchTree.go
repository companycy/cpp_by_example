/**********************************************************************************
*
* Given a singly linked list where elements are sorted in ascending order,
* convert it to a height balanced BST.
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

type Node struct {
    val         int
    left, right *Node
}

// func getBST1(head *LinkedNode, low, high int) *Node {
//     if head == nil || low > high {
//         return nil
//     }

//     mid := low + (high-low)/2
//     fmt.Println("left")
//     left := getBST1(head, low, mid-1)

//     rootval := head.val
//     head = head.next

//     fmt.Println("right")
//     right := getBST1(head, mid+1, high)
//     // root.right = right

//     fmt.Println(rootval, "\t", head.val)

//     root := &Node{
//         rootval,
//         left, right,
//     }
//     // fmt.Println(rootval)
//     return root
// }

func getBST(head **LinkedNode, low, high int) *Node {
    if head == nil || (*head) == nil || low > high {
        return nil
    }

    mid := low + (high-low)/2
    // fmt.Println("left")
    left := getBST(head, low, mid-1)

    rootval := (*head).val
    (*head) = (*head).next

    // fmt.Println("right")
    right := getBST(head, mid+1, high)

    // fmt.Println(rootval)

    root := &Node{
        rootval,
        left, right,
    }
    return root
}

func preorderTraverse(root *Node) {
    if root != nil {
        fmt.Print(root.val, "\t")
        if root.left != nil {
            preorderTraverse(root.left)
        }
        if root.right != nil {
            preorderTraverse(root.right)
        }
    }
}

func main() {
    l := []int{
        1, 3, 4, 8, 11, 15,
        // 1, 2, 3,
    }
    head := &LinkedNode{
        l[len(l)-1], nil,
    }
    linklen := 1
    for i := len(l) - 2; i >= 0; i-- {
        head = &LinkedNode{
            l[i], head,
        }
        linklen++
    }

    for root := head; root != nil; root = root.next {
        fmt.Print(root.val, "\t")
    }
    fmt.Println("")

    root := getBST(&head, 0, linklen-1)
    preorderTraverse(root)
}
