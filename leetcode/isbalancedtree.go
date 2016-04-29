/**********************************************************************************
*
* Given a binary tree, determine if it is height-balanced.
*
* For this problem, a height-balanced binary tree is defined as a binary tree in which
* the depth of the two subtrees of every node never differ by more than 1.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val int
    left, right *Node
}

func isBalanced(root *Node, height *int) (bool) {
    if root == nil {
        *height = 0
        return true
    }
    var lHeight, rHeight, diff int
    isLeftBalanced := isBalanced(root.left, &lHeight)
    isRightBalanced := isBalanced(root.right, &rHeight)
    if lHeight > rHeight {
        *height = lHeight
        diff = lHeight - rHeight
    } else {
        *height = rHeight
        diff = rHeight - lHeight
    }
    *height = *height+1
    return isLeftBalanced && isRightBalanced && diff <= 1
}

func main() {
    root := Node{
        1, nil, nil,
    }
    height := 0
    if isBalanced(&root, &height) {
        fmt.Println("is balanced")
    } else {
        fmt.Println("is not balanced with height: ", height)
    }
}
