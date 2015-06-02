/**********************************************************************************
 *
 * Given an unsorted array, find the maximum difference between the successive elements in its sorted form.
 *
 * Try to solve it in linear time/space.
 *
 * Return 0 if the array contains less than 2 elements.
 *
 * You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
 *
 * Credits:Special thanks to @porker2008 for adding this problem and creating all test cases.
 *
 **********************************************************************************/

package main

import (
    "fmt"
)

type Node struct {
    min, max int
    notEmpty bool
}

func getMaxDiff(l *[]int) int {
    if len(*l) < 2 {
        return 0
    }

    min, max := (*l)[0], (*l)[0]
    for i := range *l {
        if min > (*l)[i] {
            min = (*l)[i]
        }
        if max < (*l)[i] {
            max = (*l)[i]
        }
    }

    bucketsLen := (max-min)/(len(*l)-1) + 1
    buckets := make([]Node, len(*l))

    for i := range *l {
        bucketNum := ((*l)[i] - min) / bucketsLen
        if !buckets[bucketNum].notEmpty {
            buckets[bucketNum].max = (*l)[i]
            buckets[bucketNum].min = (*l)[i]
            buckets[bucketNum].notEmpty = true
            continue
        }
        if buckets[bucketNum].max < (*l)[i] {
            buckets[bucketNum].max = (*l)[i]
        }
        if buckets[bucketNum].min > (*l)[i] {
            buckets[bucketNum].min = (*l)[i]
        }
    }

    // fmt.Println(min, max, buckets)

    var maxDiff int
    for i := 0; i < len(buckets)-1; {
        j := 1
        for ; i+j < len(buckets) && !buckets[i+j].notEmpty; j++ {
        }
        if i+j < len(buckets) {
            if maxDiff < buckets[i+j].min-buckets[i].max {
                maxDiff = buckets[i+j].min - buckets[i].max
            }
        }
        i = i + j
    }
    return maxDiff
}

func main() {
    l := []int{
        1, 1, 1, 1, 1, 5, 5, 5, 5, 5,
        // 3, 6, 19, 1,
    }
    fmt.Println(getMaxDiff(&l))
}
