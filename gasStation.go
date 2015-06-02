/**********************************************************************************
*
* There are N gas stations along a circular route, where the amount of gas at station i is gas[i].
*
* You have a car with an unlimited gas tank and it costs cost[i] of gas to travel from station i to
* its next station (i+1). You begin the journey with an empty tank at one of the gas stations.
*
* Return the begining gas station's index if you can travel around the circuit once, otherwise return -1.
*
* Note:
* The solution is guaranteed to be unique.
*
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getbeginPos(gases, costs []int) int {
    if len(gases) != len(costs) {
        return -1
    }

    begin, end := 0, 1
    result := gases[begin] - costs[begin]
    for begin != end && end < len(gases) {
        if result >= 0 {
            result += gases[end] - costs[end]
            end++
        } else { // result < 0
            begin = (begin - 1 + len(gases)) % len(gases)
            result += gases[begin] - costs[begin]
        }
    }

    if end == len(gases) {
        return begin
    } else { // begin == end
        return -1
    }
}

func main() {
    costs := []int{
        1, 3, 5,
    }
    gases := []int{
        0, 2, 4,
    }

    beginpos := getbeginPos(gases, costs)
    fmt.Println(beginpos)
}
