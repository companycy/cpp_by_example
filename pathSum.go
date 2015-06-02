/**********************************************************************************
*
* Given a binary tree and a sum, determine if the tree has a root-to-leaf path
* such that adding up all the values along the path equals the given sum.
*
* For example:
* Given the below binary tree and sum = 22,
*
*               5
*              / \
*             4   8
*            /   / \
*           11  13  4
*          /  \      \
*         7    2      1
*
* return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
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

func getPath(root *Node, target int, result []int) bool {
    if root == nil {
        return false
    } else if root.left == nil && root.right == nil {
        if target == root.val {
            result = append(result, root.val)
            fmt.Println(result)
            return true
        } else {
            return false
        }
    }

    newtarget := target - root.val
    result = append(result, root.val)
    if root.left != nil && getPath(root.left, newtarget, result) {
        return true
    } else if root.right != nil && getPath(root.right, newtarget, result) {
        return true
    } else {
        return false
    }
}

func main() {
    root := &Node{
        5,
        &Node{4, &Node{11, &Node{7, nil, nil}, &Node{2, nil, nil}}, nil},
        &Node{8, &Node{13, nil, nil}, &Node{4, nil, &Node{1, nil, nil}}},
    }
    target := 22
    var result []int
    fmt.Println(getPath(root, target, result))
}
