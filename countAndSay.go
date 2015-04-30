/**********************************************************************************
*
* The count-and-say sequence is the sequence of integers beginning as follows:
* 1, 11, 21, 1211, 111221, ...
*
* 1 is read off as "one 1" or 11.
* 11 is read off as "two 1s" or 21.
* 21 is read off as "one 2, then one 1" or 1211.
*
* Given an integer n, generate the nth sequence.
*
* Note: The sequence of integers will be represented as a string.
*
*
**********************************************************************************/

package main

import (
    "fmt"
    "strconv"
)

type Node struct {
    val string
    cnt int
}

func getNodes(s string) *[]Node {
    var nodes []Node
    var p uint8
    for i, cur := 0, s[0]; i < len(s); i++ {
        cur = s[i]
        if p != cur {
            nodes = append(nodes, Node{string(cur), 1,})
            p = cur
        } else {
            nodes[len(nodes)-1].cnt++
        }
    }
    return &nodes
}

func buildSeq(nodes *[]Node) string {
    var seq string
    for i := 0; i < len(*nodes); i++ {
        cntstr := strconv.Itoa((*nodes)[i].cnt)
        seq += cntstr + (*nodes)[i].val
        // fmt.Println((*nodes)[i].cnt, "\t", (*nodes)[i].val)
    }
    return seq
}

func getNthSeq(s string, n int) string {
    seq := s
    for i := 0; i < n+1; i++ {
        nodes := getNodes(seq)
        seq = buildSeq(nodes)
    }
    return seq
}

func main() {
    s := "1"
    n := 3
    seq := getNthSeq(s, n)
    fmt.Println(seq)
    // buildSeq(getNodes(s))
    // s = "11"
    // buildSeq(getNodes(s))
    // s = "21"
    // seq := buildSeq(getNodes(s))
    // fmt.Println(seq)
}
