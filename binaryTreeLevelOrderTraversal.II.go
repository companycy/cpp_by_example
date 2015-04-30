/**********************************************************************************
*
* Given a binary tree, return the bottom-up level order traversal of its nodes' values.
* (ie, from left to right, level by level from leaf to root).
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
* return its bottom-up level order traversal as:
*
* [
*   [15,7],
*   [9,20],
*   [3]
* ]
*
*/

package main

import (
    "fmt"
)

type Node struct {
    val int
    left, right *Node
    // rightFlag bool
}

type LevelNode struct {
    depth int
    node *Node
}

var result [3][]int

func traverse1(root *Node) {
    depth := 0
    st := []*LevelNode {
        &LevelNode{
            depth, root,
        },
    }
    for len(st) > 0 {
        v := st[0]
        depth = v.depth
        result[depth] = append(result[depth], v.node.val)
        st = st[1:]
        if v.node.left != nil {
            st = append(st, &LevelNode{depth+1, v.node.left})
        }
        if v.node.right != nil {
            st = append(st, &LevelNode{depth+1, v.node.right})
        }
    }

    for i := len(result)-1; i >= 0; i-- {
        fmt.Println(result[i])
    }
}

func traverse(l []*Node, depth int) {
    var newl []*Node
    for i := 0; i < len(l) && len(l) > 0; i++ {
        // fmt.Println(l[i])
        result[depth] = append(result[depth], l[i].val)
        if l[i].left != nil {
            newl = append(newl, l[i].left)
        }
        if l[i].right != nil {
            newl = append(newl, l[i].right)
        }
    }
    fmt.Println("one level is finished")
    if len(newl) > 0 {
        traverse(newl, depth+1)
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

    depth := 0
    traverse(l, depth)
    for i := len(result)-1; i >= 0; i-- {
        fmt.Println(result[i])
    }
}
