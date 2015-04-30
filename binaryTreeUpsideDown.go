/**********************************************************************************
* Given a binary tree where all the right nodes are either leaf nodes with
* a sibling (a left node that shares the same parent node) or empty,
*
* Flip it upside down and turn it into a tree where the original right nodes turned into left leaf nodes.
* Return the new root.
*
* For example:
* Given a binary tree {1,2,3,4,5},
*     1
*    / \
*   2   3
*  / \
* 4   5
* return the root of the binary tree [4,5,2,#,#,3,1].
*    4
*   / \
*  5   2
*     / \
*    3   1
* confused what "{1,#,2,3}" means? > read more on how binary tree is serialized on OJ.
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

func reverse2(root *Node) {
    var left, right *Node
    head := &Node{
        0, nil, nil,
    }
    for root != nil {
        left = root.right
        right = root
        root = root.left

        right.left = head.left
        right.right = head.right

        head.left = left
        head.right = right
    }

    root = head.right
    for root != nil && root.left != nil {
        fmt.Println(root.val, root.left.val, root.right.val)
        root = root.right
    }
}

func reverse1(root *Node) {
    var st []*Node
    for root != nil {
        st = append([]*Node{root}, st...)
        root = root.left
    }
    root = st[0]

    var newroot, newleft, newright *Node
    for len(st) >= 2 {
        newroot = st[0]
        newright = st[1]
        newleft = st[1].right
        newroot.left = newleft
        newroot.right = newright
        // fmt.Println(newroot.val, newroot.left.val, newroot.right.val)
        st = st[1:]
    }
    if len(st) == 1 {
        st[0].left = nil
        st[0].right = nil
    }

    for root != nil && root.left != nil {
        fmt.Println(root.val, root.left.val, root.right.val)
        root = root.right
    }
    // fmt.Println(newroot)
}

func main() {
    root := Node{
        1, nil, nil,
    }
    left := Node{
        2, &Node{4, nil, nil}, &Node{5, nil, nil},
    }
    right := Node{
        3, nil, nil, // &Node{8, nil, nil,},
    }
    root.left = &left
    root.right = &right

    reverse1(&root)
    // reverse2(&root)
}
