/**********************************************************************************
*
* There are N children standing in a line. Each child is assigned a rating value.
*
* You are giving candies to these children subjected to the following requirements:
*
* Each child must have at least one candy.
* Children with a higher rating get more candies than their neighbors.
*
* What is the minimum candies you must give?
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func giveCandy(l *[]int) {
    newl := make([]int, len(*l))
    for i := 0; i < len(*l); i++ {
        newl[i] = 1
    }
    for i := 0; i < len(*l)-1; i++ {
        if (*l)[i] < (*l)[i+1] {
            newl[i+1] = newl[i] + 1
        }
    }
    fmt.Println(newl)
    for i := len(*l) - 1; i >= 1; i-- {
        if (*l)[i] < (*l)[i-1] {
            newl[i-1] = newl[i] + 1
        }
    }
    fmt.Println(newl)

    sum := 0
    for i := 0; i < len(newl); i++ {
        sum += newl[i]
    }
    fmt.Print(sum)
}

func main() {
    l := []int{
        5, 6, 7, 4, 1, 2, 3, 2, 1, 7,
    }

    giveCandy(&l)
}
