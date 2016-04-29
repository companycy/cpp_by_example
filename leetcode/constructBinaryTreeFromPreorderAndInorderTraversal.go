/**********************************************************************************
*
* Given preorder and inorder traversal of a tree, construct the binary tree.
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

func constructBinaryTree(preorder []int, inorder []int) *Node {
    if len(preorder) != len(inorder) || len(inorder) == 0 {
        return nil
    }

    rootval := preorder[0]
    var rootpos int
    for p, v := range inorder {
        if v == rootval {
            rootpos = p
            break
        }
    }

    left := constructBinaryTree(preorder[1:rootpos+1], inorder[0:rootpos])
    right := constructBinaryTree(preorder[rootpos+1:], inorder[rootpos+1:])
    root := &Node{
        rootval,
        left, right,
    }
    return root
}

func postorderTraverse(root *Node) {
    if root != nil {
        if root.left != nil {
            postorderTraverse(root.left)
        }
        if root.right != nil {
            postorderTraverse(root.right)
        }
        fmt.Print(root.val, "\t")
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

    preorder := []int{
        1, 4, 5, 6, 2, 3,
    }
    inorder := []int{
        5, 4, 6, 1, 3, 2,
    }
    // postorder := []int{
    //     5, 6, 4, 3, 2, 1,
    // }

    root := constructBinaryTree(preorder, inorder)

    postorderTraverse(root)
}
