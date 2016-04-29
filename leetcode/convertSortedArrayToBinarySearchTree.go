/**********************************************************************************
*
* Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val         int
    left, right *Node
}

func getBST(l []int) *Node {
    if len(l) == 0 {
        return nil
    }

    var rootpos int
    rootpos = len(l) / 2

    left := getBST((l)[:rootpos])
    right := getBST((l)[rootpos+1:])
    root := &Node{
        (l)[rootpos],
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
    }
    root := getBST(l)
    preorderTraverse(root)
}
