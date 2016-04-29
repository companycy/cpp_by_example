/**********************************************************************************
*
* Given a string s consists of upper/lower-case alphabets and empty space characters ' ',
* return the length of last word in the string.
*
* If the last word does not exist, return 0.
*
* Note: A word is defined as a character sequence consists of non-space characters only.
*
* For example,
* Given s = "Hello World",
* return 5.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getLenOfLastWord(s string) int {
    var len int
    for i := range s {
        if s[i] == ' ' {
            len = 0
        } else if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
            len++
        }
    }
    return len
}

func main() {
    s := "hello world"
    len := getLenOfLastWord(s)
    fmt.Println(len)
}
