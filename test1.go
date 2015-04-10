/*
这个问题考察的是如何将一个自然数拆成若干个小自然数的和，比如

3 = 3
3 = 2 + 1
3 = 1 + 1 + 1

或者
5 = 5
5 = 4 + 1
5 = 3 + 2
5 = 3 + 1 + 1
5 = 2 + 2 + 1
5 = 2 + 1 + 1 + 1
5 = 1 + 1 + 1 + 1 + 1

很明显这里只考虑不同的小自然数的组合，顺序是无关的，所以5=2+3和5=3+2被认为是一样的

写个小程序（语言不限），找出来12345这个数有多少种拆分方法

发信人: grepus (随风而去), 信区: NoobCoder
标  题: Re: 新马娱乐题目
发信站: 水木社区 (Sun Apr  5 22:31:19 2015), 站内

无聊分析了一下。
记对于整数 n 的所有分解的个数为 total(n), 其中最大值为 m 的分解个数为 f(n,m)。则有下列关系：
total(n) = f(n,n) + f(n,n-1) + f(n,n-2) +...+f(n,1)
对于 f(n,m) 有下列递推关系：
f(n,m) = f(n-m,m) + f(n-m,m-1) + f(n-m,m-2) +...+f(n-m,1)。

C代码如下：
#include <stdio.h>
int count_decompose(int target, int ceil){
    int result=0;
    if (ceil == 1)
        return 1;
    if(target < ceil)
        return 0;
    else if(target == ceil)
        return 1;
    else{
        for(int i=ceil; i>0; i--)
             result += count_decompose(target-ceil, i);
        return result;
    }
}
                             
int total(int target){
    int sum=0;
    for(int i=1; i<=target; i++)
        sum += count_decompose(target,i);
    return sum;
}

int main(){
    for(int i=1; i<100; i++)
        printf("%d: %d\n", i, total(i));
    return 0;
}

运行结果如下：
1: 1
2: 2
3: 3
4: 5
5: 7
6: 11
7: 15
8: 22
9: 30
10: 42
11: 56
12: 77
13: 101
14: 135
*/



package main

import (
        "fmt"
)

// f(n, m) = f(n-m, m) + f(n-m, m-1) + f(n-m, m-2) + ... + f(n-m, 1)
// f(n, m) = f(n-m, n-m) + f(n-m, n-m-1) + f(n-m, n-m-2) + ... + f(n-m, 1)
func count(n, m int) (int) {
        if n == 1 {
                return 1
        } else if n == m {
                return 1
        } else {
                result := 0
                min := 0
                if n-m > m {
                        min = m
                } else {
                        min = n-m
                }
                for i := 0; i < min; i++ {
                        result += count(n-m, i+1)
                }
                return result
        }
}


// total(n) = f(n, n) + f(n, n-1) + f(n, n-2) + ... + f(n, 1)
func total1(n int) (int) {
        result := 0
        for i := 0; i < n; i++ {
                result += count(n, i+1)
        }
        return result
}

func main() {
        for i := 1; i < 10; i++ {
                fmt.Println(i, total1(i))
        }
}
