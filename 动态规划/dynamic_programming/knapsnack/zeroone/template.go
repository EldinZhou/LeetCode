package main

import (
	"fmt"
	"math"
)

//01背包问题
//Golang中没有内置的int类型的max函数
func maxValue(N int, V int, w []int, val []int) int{
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, V+1)
	}
	//初始化dp数组
	for j := 0; j < V+1 ; j++{
		if j >= w[0] {
			dp[0][j] = w[0]
		}
	}
	//状态转移方程
	for i := 1; i < N; i++ {
		for j := 0; j < V+1 ; j++{
			if j >= w[i] {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]),float64(dp[i-1][j-w[i]] + val[i])))
			} else{
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[N-1][V]
}

func maxValueSimple(N int, V int, w []int, val []int) int {
	//利用滚动数组的思想实现一维dp解决此题
	dp := make([]int, V+1)

	for i := 0; i < N; i++ {
		for j := V; j >= w[i]; j-- {
			dp[j] = int(math.Max(float64(dp[j]), float64(dp[j-w[i]] + val[i])))
		}
	}
	return dp[V]
}

func main(){
	N, V := 3, 5
	w, val := []int{4,2,3}, []int{4,2,3}
	fmt.Println(maxValue(N,V,w,val))
	fmt.Println(maxValueSimple(N,V,w,val))
}
