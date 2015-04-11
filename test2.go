package main

import (
	"fmt"
	"strconv"
	"sort"
)

// seem go can't sort map for now
// sort its key arr and refer to value then
func test_sort(l []int) {
	m := make(map[string]int, len(l)) 
	sl := make([]string, len(l))
	for i := 0; i < len(l); i++ {
		t := strconv.Itoa(l[i])
		sl[i] = t
		m[t] = l[i]
	}
	sort.Strings(sl)

	var result string
	for i := len(sl)-1; i > 0; i-- {
		t := sl[i]
		result += t
		fmt.Println(t, m[t])
	}
	fmt.Println(result)
}

func main() {
	l := []int{3, 30, 34, 5, 9}
	test_sort(l)

	// m := make(map[int]string)
	// m[1] = "a"
	// m[2] = "c"
	// m[0] = "b"

	// // To store the keys in slice in sorted order
	// var keys []int
	// for k := range m {
	// 	keys = append(keys, k)
	// }
	// sort.Ints(keys)

	// // To perform the opertion you want
	// for _, k := range keys {
	// 	fmt.Println("Key:", k, "Value:", m[k])
	// }
}
