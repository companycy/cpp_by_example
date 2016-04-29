/**********************************************************************************
*
* Given inorder and postorder traversal of a tree, construct the binary tree.
*
* Note:
* You may assume that duplicates do not exist in the tree.
*
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

func constructBinaryTree(inorder, postorder []int) *Node {
    if len(postorder) == 0 || len(inorder) != len(postorder) {
        return nil
    }

    rootval := (postorder)[len(postorder)-1]
    var rootpos int
    for p, v := range inorder {
        if rootval == v {
            rootpos = p
            break
        }
    }
    left := constructBinaryTree((inorder)[:rootpos], (postorder)[:rootpos])
    right :=
        constructBinaryTree((inorder)[rootpos+1:], (postorder)[rootpos:len(postorder)-1])
    root := &Node{
        rootval,
        left, right,
    }
    return root
}

func preorderTraverse(root *Node) {
    if root != nil {
        fmt.Print(root.val)
        if root.left != nil {
            preorderTraverse(root.left)
        }
        if root.right != nil {
            preorderTraverse(root.right)
        }
    }
}

func main() {
    // expected := Node {
    //     1, nil, nil,
    // }
    // left := Node {
    //     4, &Node{5, nil, nil,}, &Node{6, nil, nil,},
    // }
    // right := Node {
    //     2, &Node{3, nil, nil,}, nil,
    // }
    // expected.left = &left
    // expected.right = &right

    inorder := []int{
        5, 4, 6, 1, 3, 2,
    }
    postorder := []int{
        5, 6, 4, 3, 2, 1,
    }

    root := constructBinaryTree(inorder, postorder)
    preorderTraverse(root)
}
