/**********************************************************************************
*
* Given a m x n grid filled with non-negative numbers, find a path from top left to
* bottom right which minimizes the sum of all numbers along its path.
*
* Note: You can only move either down or right at any point in time.
*
**********************************************************************************/

package main

import (
    "fmt"
)

func getMinPathSum2(grid *[][]int) int {
    dp := make([][]int, len(*grid))
    for i := 0; i < len(*grid); i++ {
        dp[i] = make([]int, len((*grid)[i]))
    }
    dp[0][0] = (*grid)[0][0]
    for i := 1; i < len(*grid); i++ {
        dp[i][0] = (*grid)[i][0] + (dp)[i-1][0]
    }
    for j := 1; j < len((*grid)[0]); j++ {
        dp[0][j] = (*grid)[0][j] + (dp)[0][j-1]
    }
    // fmt.Println(dp)

    for i := 1; i < len((*grid)); i++ {
        for j := 1; j < len((*grid)[i]); j++ {
            dp[i][j] = (*grid)[i][j]
            if (dp)[i-1][j] <= (dp)[i][j-1] {
                dp[i][j] += (dp)[i-1][j]
            } else {
                dp[i][j] += (dp)[i][j-1]
            }
        }
    }
    // fmt.Println(dp)
    return dp[len(*grid)-1][len((*grid)[0])-1]
}

func getMinPathSum(grid *[][]int, i, j int) int {
    if i >= len(*grid) || j >= len((*grid)[i]) {
        return 0
    } else if i == len(*grid)-1 && j == len((*grid)[i])-1 {
        return (*grid)[i][j]
    } else if i == len(*grid)-1 {
        var left int
        for k := j + 1; k < len((*grid)[i]); k++ {
            left = left + (*grid)[i][k]
        }
        return left
    } else if j == len((*grid)[i])-1 {
        var left int
        for k := i + 1; k < len(*grid); k++ {
            left = left + (*grid)[k][j]
        }
        return left
    }

    var down, right, minPathSum int
    if i < len(*grid) {
        down = (*grid)[i+1][j] + getMinPathSum(grid, i+1, j)
    }
    if j < len((*grid)[i]) {
        right = (*grid)[i][j+1] + getMinPathSum(grid, i, j+1)
    }

    minPathSum = down
    if minPathSum > right {
        minPathSum = right
    }
    return minPathSum + (*grid)[i][j]
}

func main() {
    grid := [][]int{
        {7, 2}, {6, 6}, {8, 6}, {8, 7}, {5, 0}, {6, 0},
    }

    fmt.Println(getMinPathSum(&grid, 0, 0))
    fmt.Println(getMinPathSum2(&grid))
}
