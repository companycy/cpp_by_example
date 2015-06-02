/**********************************************************************************
*
* Given a binary tree, flatten it to a linked list in-place.
*
* For example,
* Given
*
*          1
*         / \
*        2   5
*       / \   \
*      3   4   6
*
* The flattened tree should look like:
*
*    1
*     \
*      2
*       \
*        3
*         \
*          4
*           \
*            5
*             \
*              6
*
*
* Hints:
* If you notice carefully in the flattened tree, each node's right child points to
* the next node of a pre-order traversal.
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

type LinkedNode struct {
    val  int
    next *LinkedNode
}

func flattenBinaryTree(root *Node, head *LinkedNode) *LinkedNode {
    var tail *LinkedNode
    if root != nil {
        node := LinkedNode{
            root.val, nil,
        }
        head.next = &node
        tail = head.next
        // fmt.Println(head.val, node.val, "\t")
        if root.left != nil {
            tail = flattenBinaryTree(root.left, tail)
        }
        if root.right != nil {
            tail = flattenBinaryTree(root.right, tail)
        }
    }
    return tail
}

func main() {
    root := Node{
        1, nil, nil,
    }
    left := Node{
        2, &Node{3, nil, nil}, &Node{4, nil, nil},
    }
    right := Node{
        5, nil, &Node{6, nil, nil},
    }
    root.left = &left
    root.right = &right

    head := &LinkedNode{
        0, nil,
    }
    flattenBinaryTree(&root, head)
    for p := head.next; p != nil; p = p.next {
        fmt.Print(p.val, "\t")
    }
}
