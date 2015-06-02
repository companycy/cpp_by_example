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

type Value struct {
    val, times int
}

type LRU struct {
    m        map[int]Value
    capacity int
}

func (lru *LRU) get(key int) int {
    if v, ok := lru.m[key]; ok {
        lru.m[key].times = lru.m[key].times + 1
        return v.val
    } else {
        return -1
    }
}

func (lru *LRU) set(key int, value int) {
    if _, ok := lru.m[key]; ok {
        lru.m[key] = Value{
            value, lru.m[key].times + 1,
        }
    } else {
        if len(lru.m) == lru.capacity {
            mintimes := 0 // todo:
            var targetk int
            for k, v := range lru.m {
                if mintimes == 0 || mintimes > v.times {
                    targetk = k
                    mintimes = v.times
                }
            }
            delete(lru.m, targetk)
        }
        lru.m[key] = Value{value, 1}
    }
}

func main() {
    lru := LRU{
        make(map[int]Value), 5,
    }

    fmt.Println(lru)
}
