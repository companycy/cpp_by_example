/**********************************************************************************
*
* Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.
*
**********************************************************************************/

package main

import (
    "fmt"
)

type Point struct {
    x, y int
}

func getMaxPoints(l *[]Point) int {
    if len(*l) == 0 {
        return 0
    } else if len(*l) == 1 {
        return 1
    }

    maxPoints := 1
    for i := range *l {
        j, max := i+1, 1
        var a float64 // let y = a*x + b
        for ; j < len(*l); j++ {
            if (*l)[i].x == (*l)[j].x && (*l)[i].y == (*l)[j].y {
                max++
            } else if (*l)[i].x == (*l)[j].x {
                max++
                a = 65535
                break
            } else {
                max++
                a = float64(((*l)[j].y - (*l)[i].y)) / float64(((*l)[j].x - (*l)[i].x))
                break
            }
        }

        for j++; j < len(*l); j++ {
            var newa float64
            if (*l)[i].x == (*l)[j].x && (*l)[i].y == (*l)[j].y {
                max++
            } else {
                if (*l)[i].x == (*l)[j].x { // vertical line
                    newa = 65535
                } else {
                    newa = float64(((*l)[j].y - (*l)[i].y)) / float64(((*l)[j].x - (*l)[i].x))
                }
                if a == newa {
                    max++
                }
            }
        }

        if maxPoints < max {
            maxPoints = max
        }
    }
    return maxPoints
}

func main() {
    l := [][]Point{
        {{1, 1}, {2, 2}, {1, 5}, {3, 3}, {2, 6}, {1, 1}, {0, 0}},
    }
    for i := range l {
        fmt.Println(getMaxPoints(&l[i]))
    }
}
