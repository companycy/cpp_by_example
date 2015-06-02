/**********************************************************************************
*
* Given a binary tree, find its maximum depth.
*
* The maximum depth is the number of nodes along the longest path from the root node
* down to the farthest leaf node.
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

func getMaxDepth1(root *Node) int {
    if root == nil {
        return 0
    } else if root.left == nil && root.right == nil {
        return 1
    }
    var maxdepth int
    return maxdepth
}

func getMaxDepth(root *Node) int {
    if root == nil {
        return 0
    } else if root.left == nil && root.right == nil {
        return 1
    }

    var maxdepth int
    if root.left != nil {
        maxdepth = 1 + getMaxDepth(root.left)
    }
    if root.right != nil {
        t := 1 + getMaxDepth(root.right)
        if maxdepth < t {
            maxdepth = t
        }
    }
    return maxdepth
}

func main() {
    root := Node{
        10, &Node{5, nil, nil}, &Node{15, &Node{25, nil, nil}, nil},
    }
    fmt.Println(getMaxDepth(&root))
}
