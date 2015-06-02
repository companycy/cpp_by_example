/**********************************************************************************
*
* Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.
*
* For example:
* Given the below binary tree and sum = 22,
*
*               5
*              / \
*             4   8
*            /   / \
*           11  13  4
*          /  \    / \
*         7    2  5   1
*
* return
*
* [
*    [5,4,11,2],
*    [5,8,4,5]
* ]
*
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

func getPath(root *Node, target int, result []int, results *[][]int) {
    if root == nil {
        return
    } else if root.left == nil && root.right == nil {
        if target == root.val {
            result = append(result, root.val)
            fmt.Println(result)
            *results = append(*results, result)
            return
        } else {
            return
        }
    }

    newtarget := target - root.val
    result = append(result, root.val)
    if root.left != nil {
        getPath(root.left, newtarget, result, results)
    }

    if root.right != nil {
        getPath(root.right, newtarget, result, results)
    }
}

func main() {
    root := &Node{
        5,
        &Node{4, &Node{11, &Node{7, nil, nil}, &Node{2, nil, nil}}, nil},
        &Node{8, &Node{13, nil, nil}, &Node{4, &Node{5, nil, nil}, &Node{1, nil, nil}}},
    }
    target := 22
    var result []int
    var results [][]int
    getPath(root, target, result, &results)
}
