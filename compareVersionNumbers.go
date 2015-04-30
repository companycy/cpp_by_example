/**********************************************************************************
 *
 * Compare two version numbers version1 and version1.
 * If version1 > version2 return 1, if version1 < version2 return -1, otherwise return 0.
 *
 * You may assume that the version strings are non-empty and contain only digits and the . character.
 * The . character does not represent a decimal point and is used to separate number sequences.
 * For instance, 2.5 is not "two and a half" or "half way to version three", it is the fifth second-level revision of the second first-level revision.
 *
 * Here is an example of version numbers ordering:
 * 0.1 < 1.1 < 1.2 < 13.37
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/
package main

import (
    "fmt"
    "strings"
)

func compareVersion(lh string, rh string) int {
    pi, pj := 0, 0
    i, j := strings.Index(lh, "."), strings.Index(rh, ".")
    for ; i != -1 && j != -1; i, j = strings.Index(lh[pi+1:], "."), strings.Index(rh[pj+1:], ".") {
        if lh[pi:i] > rh[pj:j] {
            return 1
        } else if lh[pi:i] < rh[pj:j] {
            return -1
        }
        pi, pj = i, j
    }
    if i == -1 && j == -1 {
        if lh[pi+1:] > rh[pj+1:] {
            return 1
        } else if lh[pi+1:] < rh[pj+1:] {
            return -1
        } else {
            return 0
        }
    } else if j == -1 {
        return 1
    } else {
        return -1
    }
}

func main() {
    v1 := "1.2"   // "0.1"
    v2 := "13.37" // "1.1"

    if result := compareVersion(v1, v2); result == 1 {
        fmt.Println(v1, " > ", v2)
    } else if result == -1 {
        fmt.Println(v2, " > ", v1)
    } else if result == 0 {
        fmt.Println(v1, " = ", v2)
    }
}
