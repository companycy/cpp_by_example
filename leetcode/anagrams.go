/**********************************************************************************
*
* Given an array of strings, return all groups of strings that are anagrams.
*
* Note: All inputs will be in lower-case.
*
**********************************************************************************/

package main
import (
    "fmt"
    "sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func sortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func main() {
    l := []string {
        "tea","and","ate","eat","den",
    }

    m := make(map[string][]int)
    for i := 0; i < len(l); i++ {
        // sort.Strings(l[i])
        l[i] = sortString(l[i])
        fmt.Println(l[i])
        if v, ok := m[l[i]]; ok {
            v = append(v, i)
            m[l[i]] = v
            // fmt.Println(v)
        } else {
            m[l[i]] = []int {
                i,
            }
        }
    }

    for k, v := range m {
        if (len(v) > 1) {
            fmt.Println(k, v)
        }
    }
}

