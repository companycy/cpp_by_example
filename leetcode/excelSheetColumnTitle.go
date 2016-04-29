/**********************************************************************************
 *
 * Given a positive integer, return its corresponding column title as appear in an Excel sheet.
 *
 * For example:
 *
 *     1 -> A
 *     2 -> B
 *     3 -> C
 *     ...
 *     26 -> Z
 *     27 -> AA
 *     28 -> AB
 *
 * Credits:Special thanks to @ifanchu for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

func getColStr(colnum int) string {
    var result string
    var i, m int
    for i = colnum; i > 0; {
        m = i % 26
        result = fmt.Sprintf("%c", 'A'+m-1) + result
        i = i / 26
    }
    return result
}

func main() {
    colnum := 28
    colnum = 27
    s := getColStr(colnum)
    fmt.Println(s)
}
