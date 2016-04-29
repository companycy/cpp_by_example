/**********************************************************************************
*
* Given a binary tree, return the preorder traversal of its nodes' values.
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
* return [1,2,3].
*
* Note: Recursive solution is trivial, could you do it iteratively?
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Stack struct {
    l []*Node
}

func (s *Stack) empty() bool {
    return len(s.l) == 0
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

func preorderTraverse(root *Node) {
    st := Stack {
        []*Node {
            root,
        },
    }
    var l []*Node
    for !st.empty() {
        v, ok := st.pop()
        if ok {
            l = append(l, v)
            if v.left != nil {
                st.push(v.left)
            }
            if v.right != nil {
                st.push(v.right)
            }
        }
    }
    for i := 0; i < len(l); i++ {
        fmt.Println(l[i].val)
    }
}

func main() {
    root := Node{
        1, nil, nil,
    }
    // left := Node {
    //     , nil, &Node{8, nil, nil,},
    // }
    right := Node {
        2, &Node{3, nil, nil, }, nil, // &Node{7, nil, nil, },
    }
    // root.left = &left
    root.right = &right

    preorderTraverse(&root)

}
