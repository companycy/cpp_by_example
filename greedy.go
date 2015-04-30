package main

import (
	"fmt"
	"sort"
)

type pair struct {
	s, f int
}

func main() {
	m := map[int]int {
		1: 1, // pair{1, 4}},
		3: 3, // pair{3, 5}},
	}

	l := make([]int, len(m))
	i := 0
	for k, _ := range m { // i := 0; i < len(l); i++ {
		l[i] = k
		i++
	}
	sort.Ints(l)
	for i = 0; i < len(l); i++ {
		fmt.Println(l[i])
	}
}


