/**********************************************************************************
 *
 * Related to question Excel Sheet Column Title
 * Given a column title as appear in an Excel sheet, return its corresponding column number.
 *
 * For example:
 *     A -> 1
 *     B -> 2
 *     C -> 3
 *     ...
 *     Z -> 26
 *     AA -> 27
 *     AB -> 28
 *
 * Credits:Special thanks to @ts for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

func getColumnNum(s string) int {
    if len(s) == 0 {
        return 0
    }

    result := 0
    for i := range s {
        // fmt.Println(result, int(s[i]-'A'))
        result = result*26 + int(s[i]-'A') + 1
    }

    return result
}

func main() {
    s := "AB"
    s = "ABCDEFG"
    colnum := getColumnNum(s)
    fmt.Println(colnum)
}
