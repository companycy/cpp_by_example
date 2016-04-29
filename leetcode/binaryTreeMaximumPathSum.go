/**********************************************************************************
*
* Given a binary tree, find the maximum path sum.
*
* The path may start and end at any node in the tree.
*
* For example:
* Given the below binary tree,
*
*        1
*       / \
*      2   3
*
* Return 6.
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

// maxVal := -65535

func max(root *Node) int { // , maxVal *int
    rootVal := root.val
    var leftMax, rightMax int
    leftValid := false
    rightValid := false
    values := []int {
        rootVal,
    }
    if root.left != nil {
        leftMax = max(root.left)
        // fmt.Println("left", leftMax, rootVal+leftMax)
        values = append(values, leftMax, rootVal+leftMax)
        leftValid = true
    }
    if root.right != nil {
        rightMax = max(root.right)
        // fmt.Println("right", rightMax, rootVal+rightMax)
        values = append(values, rightMax, rootVal+rightMax)
        rightValid = true
    }
    
    if leftValid == true && rightValid == true {
        values = append(values, leftMax+rootVal+rightMax)
    }
    // fmt.Println(values)

    maxVal := -65536
    // root, left, right, root+left, root+right, left+root+right
    for i := 0; i < len(values); i++ {
        if (maxVal < values[i]) {
            maxVal = values[i]
        }
    }

    return maxVal
}

func main() {
    root := Node{
        1, nil, nil,
    }
    left := Node {
        2, nil, &Node{-8, nil, nil,},
    }
    right := Node {
        3, &Node{15, nil, nil, }, &Node{7, nil, nil, },
    }
    root.left = &left
    root.right = &right

    // maxVal := -65535
    maxVal := max(&root) // , maxVal
    fmt.Println(maxVal)

}
