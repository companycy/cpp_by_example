/**********************************************************************************
*
* Design and implement a data structure for Least Recently Used (LRU) cache.
* It should support the following operations: get and set.
*
*    get(key) - Get the value (will always be positive) of the key if the key exists
*               in the cache, otherwise return -1.
*
*    set(key, value) - Set or insert the value if the key is not already present.
*                      When the cache reached its capacity, it should invalidate
*                      the least recently used item before inserting a new item.
*
**********************************************************************************/

package main

import (
    "fmt"
)

type DoubleLinkedNode struct {
    key, val   int
    prev, next *DoubleLinkedNode
}

type LRU struct {
    m          map[int]*DoubleLinkedNode
    head, tail *DoubleLinkedNode
    capacity   int
}

func (lru *LRU) moveToHead(node *DoubleLinkedNode) {
    if node == nil {
        return
    } else if node == lru.head {
        return
    }

    prev, next := node.prev, node.next
    if prev != nil {
        prev.next = node.next
    }
    if next != nil {
        next.prev = node.prev
    }

    node.prev = nil // it's already head
    node.next = lru.head
    lru.head.prev = node
    lru.head = node

    // fmt.Println(node.key, node.val, node.next.key, node.next.val)
}

func (lru *LRU) removeTail() {
    if lru.tail == nil {
        return
    } else if lru.head == lru.tail {
        lru.head, lru.tail = nil, nil
    }

    delete(lru.m, lru.tail.key)
    prev := lru.tail.prev
    lru.tail = prev
    if prev != nil {
        prev.next = nil
    }
    lru.capacity = lru.capacity - 1
}

func (lru *LRU) get(key int) int {
    if v, ok := lru.m[key]; ok {
        lru.moveToHead(v)
        return v.val
    } else {
        return -1
    }
}

func (lru *LRU) set(key int, value int) {
    if v, ok := lru.m[key]; ok {
        v.val = value
        lru.moveToHead(v)
    } else {
        if len(lru.m) == lru.capacity {
            lru.removeTail()
        }

        node := DoubleLinkedNode{
            key, value, nil, nil,
        }
        lru.m[key] = &node

        if len(lru.m) == 1 { // first node, keep its prev and next nil
            lru.head = &node
            lru.tail = lru.head
        } else {
            lru.moveToHead(&node)
        }
    }
}

func main() {
    lru := LRU{
        make(map[int]*DoubleLinkedNode),
        nil, nil,
        2,
    }

    lru.set(2, 1)
    for p := lru.head; p != nil; p = p.next {
        fmt.Println(p.key, p.val)
    }
    fmt.Println()

    lru.set(2, 2)
    for p := lru.head; p != nil; p = p.next {
        fmt.Println(p.key, p.val)
    }
    fmt.Println()

    val := lru.get(2)
    for p := lru.head; p != nil; p = p.next {
        fmt.Println(p.key, p.val)
    }
    fmt.Println()

    lru.set(1, 1)
    for p := lru.head; p != nil; p = p.next {
        fmt.Println(p.key, p.val)
    }
    fmt.Println()

    lru.set(4, 1)
    for p := lru.head; p != nil; p = p.next {
        fmt.Println(p.key, p.val)
    }
    fmt.Println()

    val = lru.get(2)
    fmt.Println(val)
    for p := lru.head; p != nil; p = p.next {
        fmt.Println(p.key, p.val)
    }
    return
}
