/**********************************************************************************
 *
 * Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
 *
 * Calling next() will return the next smallest number in the BST.
 *
 * Note: next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val int
    left, right *Node
}

type BstIterator struct {
    l []int
    pos int
}

func push(root *Node, li *BstIterator) {
    if root != nil { // in-order loop
        push(root.left, li)
        li.l = append(li.l, root.val)
        push(root.right, li)
    }
}

func (li *BstIterator) next() int {
    val := li.l[li.pos]
    li.pos++
    return val
}

func (li *BstIterator) hasNext() bool {
    return li.pos < len(li.l)-1
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

    var li BstIterator
    push(&root, &li)
    fmt.Println(li)

    next := li.next()
    hasNext := li.hasNext()
    fmt.Println(next, hasNext)

}
