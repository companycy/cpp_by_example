/**********************************************************************************
 *
 * Given a binary tree, imagine yourself standing on the right side of it, return
 * the values of the nodes you can see ordered from top to bottom.
 *
 * For example:
 * Given the following binary tree,
 *
 *      1            <---
 *    /   \
 *   2     3         <---
 *    \     \
 *     5     4       <---
 *
 * You should return [1, 3, 4].
 *
 * Credits:Special thanks to @amrsaqr for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

// type Stack struct {
//     l []*LevelNode
// }

// func (s *Stack) empty() bool {
//     return len(s.l) == 0
// }

// func (s *Stack) back() *LevelNode {
//     return s.l[len(s.l)-1]
// }

// func (s *Stack) push(val *LevelNode) {
//     s.l = append([]*LevelNode{val}, s.l...)
// }

// func (s *Stack) pop() (*LevelNode, bool) {
//     if len(s.l) > 0 {
//         val := s.l[0]
//         s.l = s.l[1:]
//         return val, true
//     } else {
//         return nil, false
//     }
// }

type Node struct {
    val int
    left, right *Node
}

type LevelNode struct {
    depth int
    node *Node
}

var result [3][]int

// refer to solution of level view
func rightSideView(root *Node) {
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
            st = append(st, &LevelNode{depth+1, v.node.left,})
        }
        if v.node.right != nil {
            st = append(st, &LevelNode{depth+1, v.node.right})
        }
    }

    for i := 0; i < len(result) && len(result[i]) > 0; i++ {
        fmt.Println(result[i][len(result[i])-1])
    }
}

func main() {
    root := Node{
        1, nil, nil,
    }
    left := Node {
        2, nil, &Node{5, nil, nil,},
    }
    right := Node {
        3, nil, &Node{4, nil, nil, },
    }
    root.left = &left
    root.right = &right

    // root := Node{
    //     3, nil, nil,
    // }
    // left := Node {
    //     9, nil, nil, // &Node{8, nil, nil,},
    // }
    // right := Node {
    //     20, &Node{15, nil, nil, }, &Node{7, nil, nil, },
    // }
    // root.left = &left
    // root.right = &right


    rightSideView(&root)
}