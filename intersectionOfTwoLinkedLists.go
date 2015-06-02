/**********************************************************************************
 *
 * Write a program to find the node at which the intersection of two singly linked lists begins.
 *
 * For example, the following two linked lists:
 *
 *
 *    A:          a1 → a2
 *                       ↘  *    A:          a1 → a2                                                         |
 *                       ↗                                                                                                                                              *                         c1 → c2 → c3                                          |
 *    B:     b1 → b2 →
 * begin to intersect at node c1.
 *
 * Notes:
 *
 * If the two linked lists have no intersection at all, return null.
 * The linked lists must retain their original structure after the function returns.
 * You may assume there are no cycles anywhere in the entire linked structure.
 * Your code should preferably run in O(n) time and use only O(1) memory.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    val  int
    next *Node
}

func getLen(head *Node) int {
    var len int
    for p := head; p != nil; p = p.next {
        len++
    }
    return len
}

func getIntersection(head1, head2 *Node) *Node {
    if head1 == nil || head2 == nil {
        return nil
    } else if head1.next == nil && head2.next == nil {
        if head1.val == head2.val {
            return head1
        } else {
            return nil
        }
    }

    len1, len2 := getLen(head1), getLen(head2)
    var p1, p2 *Node
    var diff int
    if len1 >= len2 {
        p1, p2 = head1, head2
        diff = len1 - len2
    } else {
        p1, p2 = head2, head1
        diff = len2 - len1
    }

    for i := 0; i < diff; i++ {
        p1 = p1.next
    }

    var pos *Node
    for ; p1 != nil && p2 != nil; p1, p2 = p1.next, p2.next {
        if p1 == p2 && pos == nil {
            pos = p1
            break
        }
    }

    return pos
}

func main() {
    head1 := Node{
        1, &Node{
            6, &Node{
                8, &Node{
                    3, &Node{
                        6, &Node{
                            8, &Node{
                                10, nil,
                            }}}}}}}

    head2 := Node{
        6, &Node{
            8, &Node{
                6, &Node{
                    8, &Node{
                        5, &Node{
                            6, &Node{
                                8, &Node{
                                    10, nil,
                                }}}}}}}}

    for pos := getIntersection(&head1, &head2); pos != nil; pos = pos.next {
        fmt.Print(pos.val, "\t")
    }
}
