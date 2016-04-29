/**********************************************************************************
*
* You are climbing a stair case. It takes n steps to reach to the top.
*
* Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
*
*
**********************************************************************************/

package main

func climb1(n int) {
    var v []int
    v[0] = 0
    v[1] = 1
    v[2] = 2
    for i := 3; i < n+1; i++ {
        v[i] = v[i-1] + v[i-2]
    }
    println(v[n])
}

func climb2(n int) {
    x := 1
    y := 2
    var z int
    for i := 3; i < n+1; i++ {
        z = x + y
        x = y
        y = z
    }
    println(z)
}

func main() {
    climb(n)
}
