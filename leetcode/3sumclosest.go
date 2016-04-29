
package main
import (
    "fmt"
    "sort"
)

func main() {
    l := []int {-4,-1,-1,1,2}
    sort.Ints(l)
    target := 1
    for i := 0; i < len(l); i++ {
        t := target - l[i]
        m := append(l[:i], l[i+1:])
        for k := 0; k < len(m); k++ {
            fmt.Println(m[k])   
        }
        // for j := i+1; j < len(l); j++ {
        //     n := append(m[:j], m[j+1:])
        // }
    }
}

