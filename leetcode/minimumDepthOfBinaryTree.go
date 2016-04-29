/**********************************************************************************
*
* Given a binary tree, find its minimum depth.
*
* The minimum depth is the number of nodes along the shortest path from the root node
* down to the nearest leaf node.
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

func getMinDepth(root *Node) int {
    if root == nil {
        return 0
    } else if root.left == nil && root.right == nil {
        return 1
    } else if root.left == nil || root.right == nil {
        return 1
    }

    var minDepth int
    var left, right int
    if root.left != nil {
        left = 1 + getMinDepth(root.left)
    }
    if root.right != nil {
        right = 1 + getMinDepth(root.right)
    }
    minDepth = left
    if minDepth > right {
        minDepth = right
    }
    return minDepth
}

func main() {
    root := Node{
        10, &Node{5, &Node{3, nil, nil}, &Node{7, nil, nil}}, &Node{20, nil, &Node{25, nil, &Node{30, nil, nil}}},
    }
    fmt.Println(getMinDepth(&root))
}
