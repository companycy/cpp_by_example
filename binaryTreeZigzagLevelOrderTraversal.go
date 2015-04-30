/**********************************************************************************
*
* Given a binary tree, return the zigzag level order traversal of its nodes' values.
* (ie, from left to right, then right to left for the next level and alternate between).
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
* return its zigzag level order traversal as:
*
* [
*   [3],
*   [20,9],
*   [15,7]
* ]
**/

package main

import (
    "fmt"
)

type Node struct {
    val         int
    left, right *Node
}

func reverse(s *[]*Node) {
    for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
        (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
    }
}

func zzTraverse(l []*Node, fromLeftToRight bool) {
    var newl []*Node
    if fromLeftToRight {
        for i := 0; i < len(l); i++ {
            fmt.Print(l[i].val, "\t")
            if l[i].left != nil {
                newl = append(newl, l[i].left)
            }
            if l[i].right != nil {
                newl = append(newl, l[i].right)
            }
        }
    } else {
        for i := len(l) - 1; i >= 0; i-- {
            fmt.Print(l[i].val, "\t")
            if l[i].left != nil {
                newl = append(newl, l[i].left)
            }
            if l[i].right != nil {
                newl = append(newl, l[i].right)
            }
        }
        // reverse(&newl)
    }
    fmt.Println("")
    if len(newl) > 0 {
        zzTraverse(newl, !fromLeftToRight)
    }
}

func main() {
    root := Node{
        3, nil, nil,
    }
    left := Node{
        9, nil, nil, // &Node{8, nil, nil,},
    }
    right := Node{
        20, &Node{15, nil, nil}, &Node{7, nil, nil},
    }
    root.left = &left
    root.right = &right

    l := []*Node{
        &root,
    }
    fromLeftToRight := true
    zzTraverse(l, fromLeftToRight)
}
