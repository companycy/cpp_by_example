/**********************************************************************************
*
* Given a binary tree, return the level order traversal of its nodes' values.
* (ie, from left to right, level by level).
*
* For example:
* Given binary tree {3,9,20,#,#,15,7},
*
*     3
*    / \
*   9  20
*     /  \
*    15   7
*
* return its level order traversal as:
*
* [
*   [3],
*   [9,20],
*   [15,7]
* ]
*/
package main

import (
    "fmt"
)

type Node struct {
    val int
    left, right *Node
}

// func traverse(root *Node) { // use stack
//     st Stack
//     for root != nil {
//         fmt.Println(root)
//         st.push(root.left)
//     }
// }

func traverse(l []*Node) {
    var newl []*Node
    for i := 0; i < len(l) && len(l) > 0; i++ {
        fmt.Println(l[i])
        if l[i].left != nil {
            newl = append(newl, l[i].left)
        }
        if l[i].right != nil {
            newl = append(newl, l[i].right)
        }
    }
    fmt.Println("one level is finished")
    if len(newl) > 0 {
        traverse(newl)
    }
}

func main() {
    root := Node{
        3, nil, nil,
    }
    left := Node {
        9, nil, nil, // &Node{8, nil, nil,},
    }
    right := Node {
        20, &Node{15, nil, nil, }, &Node{7, nil, nil, },
    }
    root.left = &left
    root.right = &right

    l := []*Node {
        &root,
    }
    traverse(l)
    // l = append(l, root)
}
