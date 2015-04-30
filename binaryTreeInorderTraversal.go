/**********************************************************************************
*
* Given a binary tree, return the inorder traversal of its nodes' values.
*
* For example:
* Given binary tree {1,#,2,3},
*
*    1
*     \
*      2
*     /
*    3
*
* return [1,3,2].
*
* Note: Recursive solution is trivial, could you do it iteratively?
*/

package main

import (
    "fmt"
)


type Stack struct {
    l []*Node
}

func (s *Stack) push(val *Node) {
    s.l = append([]*Node{val}, s.l...)
}

func (s *Stack) pop() (*Node, bool) {
    if len(s.l) > 0 {
        val := s.l[0]
        s.l = s.l[1:]
        return val, true
    } else {
        return nil, false
    }
}

type Node struct {
    val int
    left, right *Node
}

func inOrderTraverse(root *Node) {
    var st Stack
    for root != nil {
        // fmt.Println("root", root)
        if root.left != nil {
            st.push(root)
            root = root.left
        } else {
            fmt.Println(root.val) // left
            if v, ok := st.pop(); ok {
                fmt.Println(v.val) // mid
                root = v.right
                // fmt.Println("right", root)
            } else {
                break; // 
            }
        }
    }
}

func main() {
    root := Node{
        10, nil, nil,
    }
    left := Node {
        6, nil, &Node{8, nil, nil,},
    }
    right := Node {
        15, &Node{12, nil, nil, }, nil,
    }
    root.left = &left
    root.right = &right

    // fmt.Println(root)
    inOrderTraverse(&root)
}
