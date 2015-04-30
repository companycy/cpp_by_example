package main

import (
    "fmt"
)

type Node struct {
    val int
    next *Node
}

func main() {
    // Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
    lh := []Node {
        Node{2, nil},
        Node{4, nil},
        Node{3, nil},
    }
    rh := []Node {
        Node{5, nil},
        Node{6, nil},
        Node{4, nil},
    }

    for i := 0; i < len(lh)-1; i++ {
        lh[i].next = &lh[i+1]
    }
    for i := 0; i < len(rh)-1; i++ {
        rh[i].next = &rh[i+1]
    }

    var sum *Node
    carry := 0
    for i := 0; i < len(lh); i++ {
        sumval := lh[i].val + rh[i].val + carry
        carry = 0 // reset to 0
        if sumval >= 10 {
            carry = 1
        }

        sumval = sumval % 10
        sumnode := Node{sumval, sum}
        // fmt.Println(sumval, sumnode)
        sum = &sumnode
    }

    fmt.Println(lh)
    fmt.Println(rh)
    for sum != nil {
        fmt.Println(sum.val)
        sum = sum.next
    }
}

